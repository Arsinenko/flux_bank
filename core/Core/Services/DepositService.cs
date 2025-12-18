using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class DepositService(IDepositRepository depositRepository, IMapper mapper)
    : Core.DepositService.DepositServiceBase
{
    public override async Task<GetAllDepositsResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var deposits = await depositRepository.GetAllAsync(request.PageN, request.PageSize);

        return new GetAllDepositsResponse
        {
            Deposits = { mapper.Map<IEnumerable<DepositModel>>(deposits) }
        };
    }

    public override async Task<DepositModel> Add(AddDepositRequest request, ServerCallContext context)
    {
        var deposit = mapper.Map<Deposit>(request);

        await depositRepository.AddAsync(deposit);

        return mapper.Map<DepositModel>(deposit);
    }

    public override async Task<DepositModel> GetById(GetDepositByIdRequest request, ServerCallContext context)
    {
        var deposit = await depositRepository.GetByIdAsync(request.DepositId);

        if (deposit == null)
            throw new NotFoundException("Deposit not found");

        return mapper.Map<DepositModel>(deposit);
    }

    public override async Task<Empty> Update(UpdateDepositRequest request, ServerCallContext context)
    {
        var deposit = await depositRepository.GetByIdAsync(request.DepositId);

        if (deposit == null)
            throw new NotFoundException("Deposit not found");

        mapper.Map(request, deposit);

        await depositRepository.UpdateAsync(deposit);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteDepositRequest request, ServerCallContext context)
    {
        await depositRepository.DeleteAsync(request.DepositId);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteDepositBulkRequest request, ServerCallContext context)
    {
        var ids = request.Deposits.Select(d => d.DepositId).ToList();
        if (ids.Count == 0)
        {
            throw new NotFoundException("No deposits to delete");
        }
        var deposits = await depositRepository.GetByIdsAsync(ids);
        var foundDeposits = deposits.Where(d => d != null).ToList();
        if (foundDeposits.Count != ids.Count)
        {
            throw new NotFoundException("Some deposits not found");
        }
        await depositRepository.DeleteRangeAsync(foundDeposits!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateDepositBulkRequest request, ServerCallContext context)
    {
        var deposits = request.Deposits.Select(mapper.Map<Deposit>).ToList();
        if (!deposits.Any())
        {
            throw new NotFoundException("No deposits to update");
        }
        await depositRepository.UpdateRangeAsync(deposits);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddDepositBulkRequest request, ServerCallContext context)
    {
        var deposits = request.Deposits.Select(mapper.Map<Deposit>).ToList();
        await depositRepository.AddRangeAsync(deposits);
        return new Empty();
    }

    public override async Task<GetAllDepositsResponse> GetByCustomer(GetDepositsByCustomerRequest request, ServerCallContext context)
    {
        var deps = await depositRepository.FindAsync(d => d.CustomerId == request.CustomerId);
        return new GetAllDepositsResponse()
        {
            Deposits = { mapper.Map<DepositModel>(deps) }
        };
    }

    public override async Task<GetAllDepositsResponse> GetByIds(GetDepositByIdsRequest request, ServerCallContext context)
    {
        var deps = await depositRepository.GetByIdsAsync(request.DepositIds);
        return new GetAllDepositsResponse()
        {
            Deposits = { mapper.Map<IEnumerable<DepositModel>>(deps) }
        };
    }
}
