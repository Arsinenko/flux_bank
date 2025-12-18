using Grpc.Core;
using Grpc.Core.Interceptors;
using Microsoft.EntityFrameworkCore;
using Npgsql;
namespace Core.Exceptions;


public class GrpcExceptionInterceptor : Interceptor
{
    private readonly ILogger<GrpcExceptionInterceptor> _logger;

    public GrpcExceptionInterceptor(ILogger<GrpcExceptionInterceptor> logger)
    {
        _logger = logger;
    }

    public override async Task<TResponse> UnaryServerHandler<TRequest, TResponse>(
        TRequest request,
        ServerCallContext context,
        UnaryServerMethod<TRequest, TResponse> continuation)
    {
        try
        {
            return await continuation(request, context);
        }
        catch (NotFoundException ex)
        {
            throw new RpcException(new Status(StatusCode.NotFound, ex.Message));
        }
        catch (ValidationException ex)
        {
            throw new RpcException(new Status(StatusCode.InvalidArgument, ex.Message));
        }
        catch (DbUpdateConcurrencyException)
        {
            throw new RpcException(new Status(StatusCode.Aborted, "Concurrency conflict"));
        }
        catch (DbUpdateException ex)
        {
            // Пытаемся достать конкретную ошибку Postgres
            if (ex.InnerException is PostgresException pgEx)
            {
                // 23503 = Foreign Key Violation (ссылка на несуществующую запись)
                if (pgEx.SqlState == "23503")
                {
                    // Логируем как Warning, а не Error, так как это ошибка клиента
                    _logger.LogWarning(ex, "Foreign key constraint violation: {ConstraintName}", pgEx.ConstraintName);
                    
                    // Возвращаем InvalidArgument или NotFound. 
                    // InvalidArgument лучше подходит для "вы передали ID, которого нет".
                    // В сообщении можно указать ConstraintName, чтобы понять, какое поле неверное.
                    throw new RpcException(new Status(StatusCode.InvalidArgument, $"Invalid reference: one of the provided IDs does not exist. DB Constraint: {pgEx.ConstraintName}"));
                }
                
                // 23505 = Unique Violation (дубликат уникального поля, например email или телефон)
                if (pgEx.SqlState == "23505")
                {
                    _logger.LogWarning(ex, "Unique constraint violation: {ConstraintName}", pgEx.ConstraintName);
                    throw new RpcException(new Status(StatusCode.AlreadyExists, $"Record already exists. DB Constraint: {pgEx.ConstraintName}"));
                }
            }

            // Если ошибка другая - логируем как Error и возвращаем Internal
            _logger.LogError(ex, "Database update failed");
            throw new RpcException(new Status(StatusCode.Internal, "Database error"));
        }
        catch (AutoMapper.AutoMapperMappingException ex)
        {
            _logger.LogWarning(ex, "Mapping error");
            throw new RpcException(new Status(StatusCode.InvalidArgument, "Invalid input format or data structure"));
        }
        catch (Exception ex)
        {
            _logger.LogError(ex, "Unhandled error");
            throw new RpcException(new Status(StatusCode.Internal, "Internal server error"));
        }
    }
}

public abstract class AppException : Exception
{
    protected AppException(string message) : base(message) {}
}

public class NotFoundException : AppException
{
    public NotFoundException(string message) : base(message) {}
}

public class ValidationException : AppException
{
    public ValidationException(string message) : base(message) {}
}
