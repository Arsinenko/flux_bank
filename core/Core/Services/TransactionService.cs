using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class TransactionService(ITransactionRepository transactionRepository, IMapper mapper)
    : Core.TransactionService.TransactionServiceBase
{
    public override async Task<GetAllTransactionsResponse> GetAll(Empty request, ServerCallContext context)
    {
        var transactions = await transactionRepository.GetAllAsync();

        return new GetAllTransactionsResponse
        {
            Transactions = { mapper.Map<IEnumerable<TransactionModel>>(transactions) }
        };
    }

    public override async Task<TransactionModel> Add(AddTransactionRequest request, ServerCallContext context)
    {
        var transaction = mapper.Map<Transaction>(request);

        await transactionRepository.AddAsync(transaction);

        return mapper.Map<TransactionModel>(transaction);
    }

    public override async Task<TransactionModel> GetById(GetTransactionByIdRequest request, ServerCallContext context)
    {
        var transaction = await transactionRepository.GetByIdAsync(request.TransactionId);

        if (transaction == null)
            throw new RpcException(new Status(StatusCode.NotFound, "Transaction not found"));

        return mapper.Map<TransactionModel>(transaction);
    }

    public override async Task<Empty> Update(UpdateTransactionRequest request, ServerCallContext context)
    {
        var transaction = await transactionRepository.GetByIdAsync(request.TransactionId);

        if (transaction == null)
            throw new RpcException(new Status(StatusCode.NotFound, "Transaction not found"));

        mapper.Map(request, transaction);

        await transactionRepository.UpdateAsync(transaction);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteTransactionRequest request, ServerCallContext context)
    {
        await transactionRepository.DeleteAsync(request.TransactionId);
        return new Empty();
    }
}
