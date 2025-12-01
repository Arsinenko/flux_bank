using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class AccountService(IAccountRepository accountRepository) : Core.AccountService.AccountServiceBase
{
    public override async Task<GetAllAccountsResponse> GetAll(Empty request, ServerCallContext context)
    {
        var accounts = await accountRepository.GetAllAsync();
        var accountsRep = new RepeatedField<AccountModel>();
        foreach (var account in accounts)
        {
            accountsRep.Add(new AccountModel
            {
                AccountId = account.AccountId,
                CustomerId = account.CustomerId,
                TypeId = account.TypeId,
                Iban = account.Iban,
                Balance = account.Balance?.ToString(),
                CreatedAt = account.CreatedAt.HasValue ? Timestamp.FromDateTime(account.CreatedAt.Value) : null,
                IsActive = account.IsActive
            });
        }
        return new GetAllAccountsResponse { Accounts = { accountsRep }};
    }

    public override async Task<AccountModel> Add(AddAccountRequest request, ServerCallContext context)
    {
        var account = new Account()
        {
            CustomerId = request.CustomerId,
            TypeId = request.TypeId,
            Iban = request.Iban,
            Balance = request.Balance != null ? decimal.Parse(request.Balance) : (decimal?)null,
            IsActive = request.IsActive
        };
        await accountRepository.AddAsync(account);
        return new AccountModel
        {
            AccountId = account.AccountId,
            CustomerId = account.CustomerId,
            TypeId = account.TypeId,
            Iban = account.Iban,
            Balance = account.Balance?.ToString(),
            CreatedAt = account.CreatedAt.HasValue ? Timestamp.FromDateTime(account.CreatedAt.Value) : null,
            IsActive = account.IsActive
        };
    }

    public override async Task<AccountModel> GetById(GetAccountByIdRequest request, ServerCallContext context)
    {
        var account = await accountRepository.GetByIdAsync(request.AccountId);
        if (account == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Account not found"));
        }
        return new AccountModel
        {
            AccountId = account.AccountId,
            CustomerId = account.CustomerId,
            TypeId = account.TypeId,
            Iban = account.Iban,
            Balance = account.Balance?.ToString(),
            CreatedAt = account.CreatedAt.HasValue ? Timestamp.FromDateTime(account.CreatedAt.Value) : null,
            IsActive = account.IsActive
        };
    }

    public override async Task<Empty> Update(UpdateAccountRequest request, ServerCallContext context)
    {
        var account = await accountRepository.GetByIdAsync(request.AccountId);
        if (account == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Account not found"));
        }
        account.CustomerId = request.CustomerId;
        account.TypeId = request.TypeId;
        account.Iban = request.Iban;
        account.Balance = request.Balance != null ? decimal.Parse(request.Balance) : (decimal?)null;
        account.IsActive = request.IsActive;
        await accountRepository.UpdateAsync(account);
        return new Empty(); 
    }

    public override async Task<Empty> Delete(DeleteAccountRequest request, ServerCallContext context)
    {
        await accountRepository.DeleteAsync(request.AccountId);
        return new Empty();
    }
}