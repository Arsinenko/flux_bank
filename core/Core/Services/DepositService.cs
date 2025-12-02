using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class DepositService(IDepositRepository depositRepository, IMapper mapper)
    : Core.DepositService.DepositServiceBase
{
    public override async Task<GetAllDepositsResponse> GetAll(Empty request, ServerCallContext context)
    {
        var deposits = await depositRepository.GetAllAsync();

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
            throw new RpcException(new Status(StatusCode.NotFound, "Deposit not found"));

        return mapper.Map<DepositModel>(deposit);
    }

    public override async Task<Empty> Update(UpdateDepositRequest request, ServerCallContext context)
    {
        var deposit = await depositRepository.GetByIdAsync(request.DepositId);

        if (deposit == null)
            throw new RpcException(new Status(StatusCode.NotFound, "Deposit not found"));

        mapper.Map(request, deposit);

        await depositRepository.UpdateAsync(deposit);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteDepositRequest request, ServerCallContext context)
    {
        await depositRepository.DeleteAsync(request.DepositId);
        return new Empty();
    }
}
