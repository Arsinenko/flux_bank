using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class FeeTypeService(IFeeTypeRepository feeTypeRepository, IMapper mapper)
    : Core.FeeTypeService.FeeTypeServiceBase
{
    public override async Task<GetAllFeeTypesResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var feeTypes = await feeTypeRepository.GetAllAsync(request.PageN, request.PageSize);

        return new GetAllFeeTypesResponse
        {
            FeeTypes = { mapper.Map<IEnumerable<FeeTypeModel>>(feeTypes) }
        };
    }

    public override async Task<FeeTypeModel> Add(AddFeeTypeRequest request, ServerCallContext context)
    {
        var feeType = mapper.Map<FeeType>(request);

        await feeTypeRepository.AddAsync(feeType);

        return mapper.Map<FeeTypeModel>(feeType);
    }

    public override async Task<FeeTypeModel> GetById(GetFeeTypeByIdRequest request, ServerCallContext context)
    {
        var feeType = await feeTypeRepository.GetByIdAsync(request.FeeId);

        if (feeType == null)
            throw new RpcException(new Status(StatusCode.NotFound, "FeeType not found"));

        return mapper.Map<FeeTypeModel>(feeType);
    }

    public override async Task<Empty> Update(UpdateFeeTypeRequest request, ServerCallContext context)
    {
        var feeType = await feeTypeRepository.GetByIdAsync(request.FeeId);

        if (feeType == null)
            throw new RpcException(new Status(StatusCode.NotFound, "FeeType not found"));

        mapper.Map(request, feeType);

        await feeTypeRepository.UpdateAsync(feeType);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteFeeTypeRequest request, ServerCallContext context)
    {
        await feeTypeRepository.DeleteAsync(request.FeeId);
        return new Empty();
    }
    
    public override async Task<Empty> DeleteBulk(DeleteFeeTypeBulkRequest request, ServerCallContext context)
    {
        var ids = request.FeeTypes.Select(f => f.FeeId).ToList();
        if (ids.Count == 0)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No fee types to delete"));
        }
        var feeTypes = await feeTypeRepository.GetByIdsAsync(ids);
        await feeTypeRepository.DeleteRangeAsync(feeTypes);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateFeeTypeBulkRequest request, ServerCallContext context)
    {
        var feeTypes = request.FeeTypes.Select(mapper.Map<FeeType>).ToList();
        if (!feeTypes.Any())
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No fee types to update"));
        }
        await feeTypeRepository.UpdateRangeAsync(feeTypes);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddFeeTypeBulkRequest request, ServerCallContext context)
    {
        var feeTypes = request.FeeTypes.Select(mapper.Map<FeeType>).ToList();
        try
        {
            await feeTypeRepository.AddRangeAsync(feeTypes);
            return new Empty();
        }
        catch (Exception e)
        {
            throw new RpcException(new Status(StatusCode.Internal, e.Message));
        }
    }
}
