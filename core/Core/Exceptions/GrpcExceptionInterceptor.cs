using System.ComponentModel.DataAnnotations;
using Grpc.Core;
using Grpc.Core.Interceptors;
using Microsoft.EntityFrameworkCore;

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
        catch (DbUpdateConcurrencyException ex)
        {
            throw new RpcException(new Status(StatusCode.Aborted, "Concurrency conflict"));
        }
        catch (DbUpdateException ex)
        {
            throw new RpcException(new Status(StatusCode.Internal, "Database error"));
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
