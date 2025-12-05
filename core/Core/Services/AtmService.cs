using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class AtmService(IAtmRepository atmRepository, IMapper mapper) : Core.AtmService.AtmServiceBase
{
    public override async Task<AtmModel> Add(AddAtmRequest request, ServerCallContext context)
    {
        var atm = mapper.Map<Atm>(request);
        await atmRepository.AddAsync(atm);
        return mapper.Map<AtmModel>(atm);
    }

    public override async Task<GetAllAtmsResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var atms = await atmRepository.GetAllAsync(request.PageN, request.PageSize);
        return new GetAllAtmsResponse
        {
            Atms = { mapper.Map<IEnumerable<AtmModel>>(atms) }
        };
    }

    public override async Task<Empty> Update(UpdateAtmRequest request, ServerCallContext context)
    {
        var atm = await atmRepository.GetByIdAsync(request.AtmId);
        if (atm == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "ATM not found"));
        }
        mapper.Map(request, atm);
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
        return mapper.Map<AtmModel>(atm);
    }

    public override async Task<Empty> AddBulk(AddAtmBulkRequest request, ServerCallContext context)
    {
        try
        {
            var atm = request.Atms.Select(mapper.Map<Atm>).ToList();
            await atmRepository.AddRangeAsync(atm);
            return new Empty();
        }
        catch (Exception e)
        {
            throw new RpcException(new Status(StatusCode.Internal, e.Message));
        }
    }

    public override async Task<Empty> DeleteBulk(DeleteAtmBulkRequest request, ServerCallContext context)
    {
        var ids = request.Atms.Select(a => a.AtmId).ToList();
        if (ids.Count == 0)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No ATMs to delete"));
        }
        var atms = await atmRepository.GetByIdsAsync(ids);
        var foundAtms = atms.Where(a => a != null).ToList();
        if (foundAtms.Count != ids.Count)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Some ATMs not found"));
        }
        await atmRepository.DeleteRangeAsync(foundAtms!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateAtmBulkRequest request, ServerCallContext context)
    {
        var atms = request.Atms.Select(mapper.Map<Atm>).ToList();
        if (!atms.Any())
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No ATMs to update"));
        }
        await atmRepository.UpdateRangeAsync(atms);
        return new Empty(); 
    }
}