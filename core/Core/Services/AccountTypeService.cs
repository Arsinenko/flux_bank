using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
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
            throw new NotFoundException("Account type not found");
        }
        return mapper.Map<AccountTypeModel>(accountType);
    }

    public override async Task<Empty> Update(UpdateAccountTypeRequest request, ServerCallContext context)
    {
        var accountType = await accountTypeRepository.GetByIdAsync(request.TypeId);
        if (accountType == null)
        {
            throw new NotFoundException("Account type not found");
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
            throw new NotFoundException("Account type not found");
        }
        await accountTypeRepository.DeleteAsync(request.TypeId);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteAccountTypeBulkRequest request, ServerCallContext context)
    {
        var ids = request.AccountTypes.Select(a => a.TypeId).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No account types to delete");
        }
        var accountTypes = await accountTypeRepository.GetByIdsAsync(ids);
        var foundAccountTypes = accountTypes.Where(at => at != null).ToList();
        if (foundAccountTypes.Count != ids.Count)
        {
            throw new ValidationException("One or more account types not found");
        }
        await accountTypeRepository.DeleteRangeAsync(foundAccountTypes!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateAccountTypeBulkRequest request, ServerCallContext context)
    {
        var accountTypes = request.AccountTypes.Select(mapper.Map<AccountType>).ToList();
        if (!accountTypes.Any())
        {
            throw new ValidationException("No account types found");
        }
        await accountTypeRepository.UpdateRangeAsync(accountTypes);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddAccountTypeBulkRequest request, ServerCallContext context)
    {
        var accountTypes = request.AccountTypes.Select(a => mapper.Map<AccountType>(a));
        await accountTypeRepository.AddRangeAsync(accountTypes);
        return new Empty();
    }

    public override async Task<GetAllAccountTypesResponse> GetByIds(GetAccountTypeByIdsRequest request,
        ServerCallContext context)
    {
        var accountTypes =
            await accountTypeRepository.GetByIdsAsync(request.TypeIds);
        return new GetAllAccountTypesResponse()
        {
            AccountTypes = { mapper.Map<IEnumerable<AccountTypeModel>>(accountTypes) }
        };
    }
}