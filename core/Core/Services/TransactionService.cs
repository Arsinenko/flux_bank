using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using System.Linq;
using Core.Exceptions;

namespace Core.Services;

public class TransactionService(ITransactionRepository transactionRepository, IMapper mapper, ICacheService cacheService, IStatsService statsService)
    : Core.TransactionService.TransactionServiceBase
{
    public override async Task<GetAllTransactionsResponse> GetAll(GetAllRequest request, ServerCallContext context)
    {
        var transactions = await transactionRepository.GetAllAsync(request.PageN, request.PageSize, request.OrderBy, request.IsDesc ?? false);

        return new GetAllTransactionsResponse
        {
            Transactions = { mapper.Map<IEnumerable<TransactionModel>>(transactions) }
        };
    }

    public override async Task<TransactionModel> Add(AddTransactionRequest request, ServerCallContext context)
    {
        var transaction = mapper.Map<Transaction>(request);

        await transactionRepository.AddAsync(transaction);
        cacheService.Remove("BankStats");

        return mapper.Map<TransactionModel>(transaction);
    }

    public override async Task<TransactionModel> GetById(GetTransactionByIdRequest request, ServerCallContext context)
    {
        var transaction = await transactionRepository.GetByIdAsync(request.TransactionId);

        if (transaction == null)
            throw new NotFoundException("Transaction not found");

        return mapper.Map<TransactionModel>(transaction);
    }

    public override async Task<Empty> Update(UpdateTransactionRequest request, ServerCallContext context)
    {
        var transaction = await transactionRepository.GetByIdAsync(request.TransactionId);

        if (transaction == null)
            throw new NotFoundException("Transaction not found");

        mapper.Map(request, transaction);

        await transactionRepository.UpdateAsync(transaction);
        cacheService.Remove("BankStats");

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteTransactionRequest request, ServerCallContext context)
    {
        await transactionRepository.DeleteAsync(request.TransactionId);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteTransactionBulkRequest request, ServerCallContext context)
    {
        var ids = request.Transactions.Select(t => t.TransactionId).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No transactions to delete");
        }
        var transactions = await transactionRepository.GetByIdsAsync(ids);
        await transactionRepository.DeleteRangeAsync(transactions.Where(t => t is not null)!);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateTransactionBulkRequest request, ServerCallContext context)
    {
        var transactions = request.Transactions.Select(mapper.Map<Transaction>).ToList();
        if (!transactions.Any())
        {
            throw new ValidationException("No transactions to update");
        }
        await transactionRepository.UpdateRangeAsync(transactions);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddTransactionBulkRequest request, ServerCallContext context)
    {
        var transactions = request.Transactions.Select(mapper.Map<Transaction>).ToList();
        await transactionRepository.AddRangeAsync(transactions);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<GetAllTransactionsResponse> GetByDateRange(GetByDateRangeRequest request, ServerCallContext context)
    {
        var transactions = await transactionRepository.GetByDateRange(request.FromDate.ToDateTime(), request.ToDate.ToDateTime(), request.PageN, request.PageSize);
        return new GetAllTransactionsResponse()
        {
            Transactions = { mapper.Map<IEnumerable<TransactionModel>>(transactions) }
        };
    }

    public override async Task<GetAllTransactionsResponse> GetAccountExpenses(GetAccountExpensesRequest request, ServerCallContext context)
    {
        var transactions = await transactionRepository.GetExpensesAsync(request.SourceAccount, request.DateRange?.FromDate.ToDateTime(), request.DateRange?.ToDate.ToDateTime(), request.DateRange?.PageN, request.DateRange?.PageSize);
        return new GetAllTransactionsResponse()
        {
            Transactions = { mapper.Map<IEnumerable<TransactionModel>>(transactions) }
        };
    }

    public override async Task<GetAllTransactionsResponse> GetAccountRevenue(GetAccountRevenueRequest request, ServerCallContext context)
    {
        var transactions = await transactionRepository.GetRevenueAsync(request.TargetAccount, request.DateRange?.FromDate.ToDateTime(), request.DateRange?.ToDate.ToDateTime(), request.DateRange?.PageN, request.DateRange?.PageSize);
        return new GetAllTransactionsResponse()
        {
            Transactions = { mapper.Map<IEnumerable<TransactionModel>>(transactions) }
        };
    }

    public override async Task<GetAllTransactionsResponse> GetByIds(GetTransactionByIdsRequest request, ServerCallContext context)
    {
        var transactions = await transactionRepository.GetByIdsAsync(request.TransactionIds);
        return new GetAllTransactionsResponse()
        {
            Transactions = { mapper.Map<TransactionModel>(transactions) }
        };
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var count = await transactionRepository.GetCountAsync();
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<CountResponse> GetCountAccountExpenses(GetAccountExpensesRequest request, ServerCallContext context)
    {
        var from = request.DateRange.FromDate.ToDateTime();
        var to = request.DateRange.ToDate.ToDateTime();
        var count = await transactionRepository.GetCountAsync(t => t.SourceAccount == request.SourceAccount && t.CreatedAt >= from && t.CreatedAt <= to);
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<CountResponse> GetCountAccountRevenue(GetAccountRevenueRequest request, ServerCallContext context)
    {
        var from = request.DateRange.FromDate.ToDateTime();
        var to = request.DateRange.ToDate.ToDateTime();
        var count = await transactionRepository.GetCountAsync(t => t.TargetAccount == request.TargetAccount && t.CreatedAt >= from && t.CreatedAt <= to);
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<CountResponse> GetCountByDateRange(GetByDateRangeRequest request, ServerCallContext context)
    {
        var count = await transactionRepository.GetCountByDateRangeAsync(request.FromDate.ToDateTime(),
            request.ToDate.ToDateTime());
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<TotalAmountResponse> GetTotalAmount(Empty request, ServerCallContext context)
    {
        var stats = await statsService.GetStatsAsync();
        return new TotalAmountResponse()
        {
            TotalAmount = stats.TotalTransactionSum.ToString()
        };
    }
}
