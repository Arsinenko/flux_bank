using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class TransactionFeeService(ITransactionFeeRepository transactionFeeRepository, IMapper mapper)
    : Core.TransactionFeeService.TransactionFeeServiceBase
{
    public override async Task<GetAllTransactionFeesResponse> GetAll(GetAllRequest request, ServerCallContext context)
    {
        var transactionFees = await transactionFeeRepository.GetAllAsync(request.PageN, request.PageSize, request.OrderBy, request.IsDesc ?? false);

        return new GetAllTransactionFeesResponse
        {
            TransactionFees = { mapper.Map<IEnumerable<TransactionFeeModel>>(transactionFees) }
        };
    }

    public override async Task<TransactionFeeModel> Add(AddTransactionFeeRequest request, ServerCallContext context)
    {
        var transactionFee = mapper.Map<TransactionFee>(request);

        await transactionFeeRepository.AddAsync(transactionFee);

        return mapper.Map<TransactionFeeModel>(transactionFee);
    }

    public override async Task<TransactionFeeModel> GetById(GetTransactionFeeByIdRequest request, ServerCallContext context)
    {
        var transactionFee = await transactionFeeRepository.GetByIdAsync(request.Id);

        if (transactionFee == null)
            throw new NotFoundException("TransactionFee not found");

        return mapper.Map<TransactionFeeModel>(transactionFee);
    }

    public override async Task<Empty> Update(UpdateTransactionFeeRequest request, ServerCallContext context)
    {
        var transactionFee = await transactionFeeRepository.GetByIdAsync(request.Id);

        if (transactionFee == null)
            throw new NotFoundException("TransactionFee not found");

        mapper.Map(request, transactionFee);

        await transactionFeeRepository.UpdateAsync(transactionFee);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteTransactionFeeRequest request, ServerCallContext context)
    {
        await transactionFeeRepository.DeleteAsync(request.Id);
        return new Empty();
    }

    public override async Task<GetAllTransactionFeesResponse> GetByIds(GetTransactionFeeByIdsRequest request, ServerCallContext context)
    {
        var transactionFees = await transactionFeeRepository.GetByIdsAsync(request.Ids);
        return new GetAllTransactionFeesResponse()
        {
            TransactionFees = { mapper.Map<TransactionFeeModel>(transactionFees) }
        };
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var count = await transactionFeeRepository.GetCountAsync();
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<Empty> AddBulk(AddTransactionFeeBulkRequest request, ServerCallContext context)
    {
        var transactionFees = request.TransactionFees.Select(mapper.Map<TransactionFee>).ToList();
        await transactionFeeRepository.AddRangeAsync(transactionFees);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteTransactionFeeBulkRequest request, ServerCallContext context)
    {
        var ids = request.TransactionFees.Select(t => t.Id).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No transaction fees to delete");
        }
        var transactionFees = await transactionFeeRepository.GetByIdsAsync(ids);
        var foundTransactionFees = transactionFees.Where(tf => tf != null).ToList();

        if (foundTransactionFees.Count != ids.Count)
        {
            throw new NotFoundException("Some transaction fees not found");
        }
        await transactionFeeRepository.DeleteRangeAsync(foundTransactionFees!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateTransactionFeeBulkRequest request, ServerCallContext context)
    {
        var transactionFees = request.TransactionFees.Select(mapper.Map<TransactionFee>).ToList();
        if (!transactionFees.Any())
        {
            throw new ValidationException("No transaction fees to update");
        }

        await transactionFeeRepository.UpdateRangeAsync(transactionFees);
        return new Empty();
    }
}
