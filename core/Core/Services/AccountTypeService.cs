using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class AccountTypeService(IAccountTypeRepository accountTypeRepository, IMapper mapper) : Core.AccountTypeService.AccountTypeServiceBase
{
    public override async Task<AccountTypeModel> Add(AddAccountTypeRequest request, ServerCallContext context)
    {
       var accountType = mapper.Map<AccountType>(request);
       await accountTypeRepository.AddAsync(accountType);
       return mapper.Map<AccountTypeModel>(accountType);
    }

    public override async Task<GetAllAccountTypesResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var accountTypes = await accountTypeRepository.GetAllAsync(request.PageN, request.PageSize);
        return new GetAllAccountTypesResponse
        {
            AccountTypes = { mapper.Map<IEnumerable<AccountTypeModel>>(accountTypes) }
        };  
    }
    
    public override async Task<AccountTypeModel> GetById(GetAccountTypeByIdRequest request, ServerCallContext context)
    {
        var accountType = await accountTypeRepository.GetByIdAsync(request.TypeId);
        if (accountType == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Account type not found"));
        }
        return mapper.Map<AccountTypeModel>(accountType);
    }

    public override async Task<Empty> Update(UpdateAccountTypeRequest request, ServerCallContext context)
    {
        var accountType = await accountTypeRepository.GetByIdAsync(request.TypeId);
        if (accountType == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Account type not found"));
        }
        mapper.Map(request, accountType);
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