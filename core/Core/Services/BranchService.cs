using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Core.Exceptions;

namespace Core.Services;

public class BranchService(IBranchRepository branchRepository, IMapper mapper) : Core.BranchService.BranchServiceBase
{
    public override async Task<BranchModel> Add(AddBranchRequest request, ServerCallContext context)
    {
        var branch = mapper.Map<Branch>(request);
        await branchRepository.AddAsync(branch);
        return mapper.Map<BranchModel>(branch);
    }

    public override async Task<GetAllBranchesResponse> GetAll(GetAllRequest request, ServerCallContext context)
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
            throw new NotFoundException("Branch not found");
        }
        return mapper.Map<BranchModel>(branch);
    }

    public override async Task<Empty> Update(UpdateBranchRequest request, ServerCallContext context)
    {
        var branch = await branchRepository.GetByIdAsync(request.BranchId);
        if (branch == null)
        {
            throw new NotFoundException("Branch not found");
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
            throw new NotFoundException("Branch not found");
        }
        await branchRepository.DeleteAsync(branch.BranchId);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteBranchBulkRequest request, ServerCallContext context)
    {
        var ids = request.Branches.Select(b => b.BranchId).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No branches to delete");
        }
        var branches = await branchRepository.GetByIdsAsync(ids);
        var foundBranches = branches.Where(b => b != null).ToList();
        if (foundBranches.Count != ids.Count)
        {
            throw new ValidationException("Some branches not found");
        }
        await branchRepository.DeleteRangeAsync(foundBranches!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateBranchBulkRequest request, ServerCallContext context)
    {
        var branches = request.Branches.Select(mapper.Map<Branch>).ToList();
        if (!branches.Any())
        {
            throw new ValidationException("No branches to update");
        }
        await branchRepository.UpdateRangeAsync(branches);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddBranchBulkRequest request, ServerCallContext context)
    {
        var branches = request.Branches.Select(mapper.Map<Branch>).ToList();
        await branchRepository.AddRangeAsync(branches);
        return new Empty();
    }

    public override async Task<GetAllBranchesResponse> GetByIds(GetBranchByIdsRequest request, ServerCallContext context)
    {
        var branches = await branchRepository.GetByIdsAsync(request.BranchIds);
        return new GetAllBranchesResponse()
        {
            Branches = { mapper.Map<IEnumerable<BranchModel>>(branches) }
        };
    }
}