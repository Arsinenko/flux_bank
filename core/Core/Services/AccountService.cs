using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class AccountService(IAccountRepository accountRepository, IMapper mapper)
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
        var account = mapper.Map<Account>(request);

        await accountRepository.AddAsync(account);

        return mapper.Map<AccountModel>(account);
    }

    public override async Task<AccountModel> GetById(GetAccountByIdRequest request, ServerCallContext context)
    {
        var account = await accountRepository.GetByIdAsync(request.AccountId);

        if (account == null)
            throw new RpcException(new Status(StatusCode.NotFound, "Account not found"));

        return mapper.Map<AccountModel>(account);
    }

    public override async Task<Empty> Update(UpdateAccountRequest request, ServerCallContext context)
    {
        var account = await accountRepository.GetByIdAsync(request.AccountId);

        if (account == null)
            throw new RpcException(new Status(StatusCode.NotFound, "Account not found"));

        mapper.Map(request, account);

        await accountRepository.UpdateAsync(account);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteAccountRequest request, ServerCallContext context)
    {
        await accountRepository.DeleteAsync(request.AccountId);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteAccountBulkRequest request, ServerCallContext context)
    {
        var accounts = (await accountRepository.GetByIdsAsync(request.Accounts
            .Select(a => a.AccountId))).ToList();
        if (accounts.Count != request.Accounts.Count)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "One or more accounts not found"));
        }
        await accountRepository.DeleteRangeAsync(accounts!);
        return new Empty();
    }

    public override async Task<GetAllAccountsResponse> GetByDateRange(GetByDateRangeRequest request, ServerCallContext context)
    {
        var accounts = await accountRepository.GetByDateRange(request.From.ToDateTime(), request.To.ToDateTime(), request.PageN, request.PageSize);
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
            throw new RpcException(new Status(StatusCode.NotFound, "No accounts to update"));
        }
        await accountRepository.UpdateRangeAsync(accounts);
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
}