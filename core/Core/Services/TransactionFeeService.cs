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
    public override async Task<GetAllTransactionFeesResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var transactionFees = await transactionFeeRepository.GetAllAsync(request.PageN, request.PageSize);

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
}
