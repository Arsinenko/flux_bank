using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class AtmService(IAtmRepository atmRepository) : Core.AtmService.AtmServiceBase
{
    public override async Task<AtmModel> Add(AddAtmRequest request, ServerCallContext context)
    {
        var atm = new Atm()
        {
            Location = request.Location,
            Status = request.Status,
            BranchId = request.BranchId
        };
        await atmRepository.AddAsync(atm);
        return new AtmModel
        {
            AtmId = atm.AtmId,
            Location = atm.Location,
            Status = atm.Status,
            BranchId = atm.BranchId.Value
        };
    }

    public override async Task<GetAllAtmsResponse> GetAll(Empty request, ServerCallContext context)
    {
        var atms = await atmRepository.GetAllAsync();
        var atmsRep = new RepeatedField<AtmModel>();
        foreach (var atm in atms)
        {
            atmsRep.Add(new AtmModel
            {
                AtmId = atm.AtmId,
                Location = atm.Location,
                Status = atm.Status,
                BranchId = atm.BranchId.Value
            });
        }
        return new GetAllAtmsResponse { Atms = { atmsRep } };
    }

    public override async Task<Empty> Update(UpdateAtmRequest request, ServerCallContext context)
    {
        var atm = await atmRepository.GetByIdAsync(request.AtmId);
        if (atm == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "ATM not found"));
        }
        atm.Location = request.Location;
        atm.Status = request.Status;
        atm.BranchId = request.BranchId;
        await atmRepository.UpdateAsync(atm);
        return new Empty();        
    }

    public override async Task<Empty> Delete(DeleteAtmRequest request, ServerCallContext context)
    {
        var atm = await atmRepository.GetByIdAsync(request.AtmId);
        if (atm == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "ATM not found"));
        }
        await atmRepository.DeleteAsync(request.AtmId);
        return new Empty();
    }
    
    public override async Task<AtmModel> GetById(GetAtmByIdRequest request, ServerCallContext context)
    {
        var atm = await atmRepository.GetByIdAsync(request.AtmId);
        if (atm == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "ATM not found"));
        }
        return new AtmModel
        {
            AtmId = atm.AtmId,
            Location = atm.Location,
            Status = atm.Status,
            BranchId = atm.BranchId.Value
        };
    }
}