using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class BranchService(IBranchRepository branchRepository) : Core.BranchService.BranchServiceBase
{
    public override async Task<BranchModel> Add(AddBranchRequest request, ServerCallContext context)
    {
        var branch = new Branch()
        {
            Name = request.Name,
            Address = request.Address,
            City = request.City,
            Phone = request.Phone
        };
        await branchRepository.AddAsync(branch);
        return new BranchModel
        {
            BranchId = branch.BranchId,
            Name = branch.Name,
            Address = branch.Address
        };
    }

    public override async Task<GetAllBranchesResponse> GetAll(Empty request, ServerCallContext context)
    {
        var branches = await branchRepository.GetAllAsync();
        var branchesRep = new RepeatedField<BranchModel>();
        foreach (var branch in branches)
        {
            branchesRep.Add(new BranchModel
            {
                BranchId = branch.BranchId,
                Name = branch.Name,
                Address = branch.Address,
                City = branch.City,
                Phone = branch.Phone
            });
        }
        return new GetAllBranchesResponse { Branches = { branchesRep } };
    }
    public override async Task<BranchModel> GetById(GetBranchByIdRequest request, ServerCallContext context)
    {
        var branch = await branchRepository.GetByIdAsync(request.BranchId);
        if (branch == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Branch not found"));
        }
        return new BranchModel
        {
            BranchId = branch.BranchId,
            Name = branch.Name,
            Address = branch.Address,
            City = branch.City,
            Phone = branch.Phone
        };
    }

    public override async Task<Empty> Update(UpdateBranchRequest request, ServerCallContext context)
    {
        var branch = await branchRepository.GetByIdAsync(request.BranchId);
        if (branch == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Branch not found"));
        }
        branch.Name = request.Name;
        branch.Address = request.Address;
        branch.City = request.City;
        branch.Phone = request.Phone;
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