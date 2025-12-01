using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class AccountTypeService(IAccountTypeRepository accountTypeRepository) : Core.AccountTypeService.AccountTypeServiceBase
{
    public override async Task<AccountTypeModel> Add(AddAccountTypeRequest request, ServerCallContext context)
    {
        var accountType = new AccountType
        {
            Name = request.Name,
            Description = request.Description
        };
        await accountTypeRepository.AddAsync(accountType);
        return new AccountTypeModel()
        {
            TypeId = accountType.TypeId,
            Name = accountType.Name,
            Description = accountType.Description
        };
    }

    public override async Task<GetAllAccountTypesResponse> GetAll(Empty request, ServerCallContext context)
    {
        var accountTypes = await accountTypeRepository.GetAllAsync();
        var accountTypesRep = new RepeatedField<AccountTypeModel>();
        foreach (var accountType in accountTypes)
        {
            accountTypesRep.Add(new AccountTypeModel
            {
                TypeId = accountType.TypeId,
                Name = accountType.Name,
                Description = accountType.Description
            });
        }
        return new GetAllAccountTypesResponse { AccountTypes = { accountTypesRep } };
    }
    
    public override async Task<AccountTypeModel> GetById(GetAccountTypeByIdRequest request, ServerCallContext context)
    {
        var accountType = await accountTypeRepository.GetByIdAsync(request.TypeId);
        if (accountType == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Account type not found"));
        }
        return new AccountTypeModel
        {
            TypeId = accountType.TypeId,
            Name = accountType.Name,
            Description = accountType.Description
        };
    }

    public override async Task<Empty> Update(UpdateAccountTypeRequest request, ServerCallContext context)
    {
        var accountType = await accountTypeRepository.GetByIdAsync(request.TypeId);
        if (accountType == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Account type not found"));
        }
        accountType.Name = request.Name;
        accountType.Description = request.Description;
        await accountTypeRepository.UpdateAsync(accountType);
        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteAccountTypeRequest request, ServerCallContext context)
    {
        var accountType = await accountTypeRepository.GetByIdAsync(request.TypeId);
        if (accountType == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Account type not found"));
        }
        await accountTypeRepository.DeleteAsync(request.TypeId);
        return new Empty();
    }
}