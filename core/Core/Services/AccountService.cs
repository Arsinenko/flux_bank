using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Microsoft.EntityFrameworkCore;

namespace Core.Services;

public class AccountService(IAccountRepository accountRepository, IMapper mapper, ICacheService cacheService, IStatsService statsService)
    : Core.AccountService.AccountServiceBase
{
    public override async Task<GetAllAccountsResponse> GetAll(GetAllRequest request, ServerCallContext context)
    {

        var accounts = await accountRepository.GetAllAsync(request.PageN, request.PageSize);

        return new GetAllAccountsResponse
        {
            Accounts = { mapper.Map<IEnumerable<AccountModel>>(accounts) }
        };
    }

    public override async Task<AccountModel> Add(AddAccountRequest request, ServerCallContext context)
    {
        ValidateBalance(request.Balance);
        var account = mapper.Map<Account>(request);

        await accountRepository.AddAsync(account);
        cacheService.Remove("BankStats");

        return mapper.Map<AccountModel>(account);
    }

    public override async Task<AccountModel> GetById(GetAccountByIdRequest request, ServerCallContext context)
    {
        var account = await accountRepository.GetByIdAsync(request.AccountId);

        if (account == null)
            throw new NotFoundException("Account not found");

        return mapper.Map<AccountModel>(account);
    }

    public override async Task<Empty> Update(UpdateAccountRequest request, ServerCallContext context)
    {
        ValidateBalance(request.Balance);
        var account = await accountRepository.GetByIdAsync(request.AccountId);

        if (account == null)
            throw new NotFoundException("Account not found");

        mapper.Map(request, account);

        await accountRepository.UpdateAsync(account);
        cacheService.Remove("BankStats");

        return new Empty();
    }

    private static void ValidateBalance(string? balance)
    {
        if (!string.IsNullOrWhiteSpace(balance) && !decimal.TryParse(balance, System.Globalization.CultureInfo.InvariantCulture, out _))
        {
            throw new ValidationException($"Invalid balance format: {balance}");
        }
    }

    public override async Task<Empty> Delete(DeleteAccountRequest request, ServerCallContext context)
    {
        await accountRepository.DeleteAsync(request.AccountId);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteAccountBulkRequest request, ServerCallContext context)
    {
        var accounts = (await accountRepository.GetByIdsAsync(request.Accounts
            .Select(a => a.AccountId))).ToList();
        if (accounts.Count != request.Accounts.Count || accounts.Count == 0)
        {
            throw new ValidationException("One or more accounts not found");
        }
        await accountRepository.DeleteRangeAsync(accounts!);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<GetAllAccountsResponse> GetByDateRange(GetByDateRangeRequest request, ServerCallContext context)
    {
        var accounts = await accountRepository.GetByDateRange(request.FromDate.ToDateTime(), request.ToDate.ToDateTime(), request.PageN, request.PageSize);
        return new GetAllAccountsResponse()
        {
            Accounts = { mapper.Map<IEnumerable<AccountModel>>(accounts) }
        };
    }

    public override async Task<Empty> UpdateBulk(UpdateAccountBulkRequest request, ServerCallContext context)
    {
        var accounts = request.Accounts.Select(mapper.Map<Account>).ToList();
        if (!accounts.Any())
        {
            throw new ValidationException("No accounts found");
        }
        await accountRepository.UpdateRangeAsync(accounts);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<GetAllAccountsResponse> GetByCustomerId(GetAccountByCustomerIdRequest request, ServerCallContext context)
    {
        var accounts = await accountRepository.GetByCustomerIdAsync(request.CustomerId);
        return new GetAllAccountsResponse()
        {
            Accounts = { mapper.Map<IEnumerable<AccountModel>>(accounts) }
        };
    }

    public override async Task<GetAllAccountsResponse> GetByIds(GetAccountByIdsRequest request, ServerCallContext context)
    {
        var accounts =
            await accountRepository.FindAsync(e => request.AccountIds.Contains(EF.Property<int>(e, "AccountId")));
        return new GetAllAccountsResponse()
        {
            Accounts = { mapper.Map<IEnumerable<AccountModel>>(accounts) }
        };
    }

    public override async Task<Empty> AddBulk(AddAccountBulkRequest request, ServerCallContext context)
    {
        var accounts = request.Accounts.Select(mapper.Map<Account>).ToList();
        await accountRepository.AddRangeAsync(accounts);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var count = await accountRepository.GetCountAsync();
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<CountResponse> GetCountByCustomerId(GetAccountByCustomerIdRequest request, ServerCallContext context)
    {
        var accounts = await accountRepository.FindAsync(a => a.CustomerId == request.CustomerId);
        return new CountResponse()
        {
            Count = accounts.Count()
        };
    }

    public override async Task<CountResponse> GetCountByDateRange(GetByDateRangeRequest request, ServerCallContext context)
    {
        var count = await accountRepository.GetCountByDateRangeAsync(request.FromDate.ToDateTime(), request.ToDate.ToDateTime());
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<CountResponse> GetCountByStatus(GetAccountsByStatusRequest request, ServerCallContext context)
    {
        if (request.Status)
        {
            var stats = await statsService.GetStatsAsync();
            return new CountResponse()
            {
                Count = stats.ActiveAccountCount
            };
        }
        var count = await accountRepository.GetCountByStatusAsync(request.Status);
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<TotalBalanceResponse> GetTotalBalance(Empty request, ServerCallContext context)
    {
        var stats = await statsService.GetStatsAsync();
        return new TotalBalanceResponse()
        {
            TotalBalance = stats.TotalBalance.ToString()
        };
    }
}