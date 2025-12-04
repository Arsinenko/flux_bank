using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class BranchService(IBranchRepository branchRepository, IMapper mapper) : Core.BranchService.BranchServiceBase
{
    public override async Task<BranchModel> Add(AddBranchRequest request, ServerCallContext context)
    {
        var branch = mapper.Map<Branch>(request);
        await branchRepository.AddAsync(branch);
        return mapper.Map<BranchModel>(branch); 
    }

    public override async Task<GetAllBranchesResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var branches = await branchRepository.GetAllAsync(request.PageN, request.PageSize);
        return new GetAllBranchesResponse()
        {
            Branches = { mapper.Map<IEnumerable<BranchModel>>(branches) }
        };
    }
    public override async Task<BranchModel> GetById(GetBranchByIdRequest request, ServerCallContext context)
    {
        var branch = await branchRepository.GetByIdAsync(request.BranchId);
        if (branch == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Branch not found"));
        }
        return mapper.Map<BranchModel>(branch);
    }

    public override async Task<Empty> Update(UpdateBranchRequest request, ServerCallContext context)
    {
        var branch = await branchRepository.GetByIdAsync(request.BranchId);
        if (branch == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Branch not found"));
        }
        mapper.Map(request, branch);
        await branchRepository.UpdateAsync(branch);
        return new Empty();
    }
    public override async Task<Empty> Delete(DeleteBranchRequest request, ServerCallContext context)
    {
        var branch = await branchRepository.GetByIdAsync(request.BranchId);
        if (branch == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Branch not found"));
        }
        await branchRepository.DeleteAsync(branch.BranchId);
        return new Empty();
    }
}