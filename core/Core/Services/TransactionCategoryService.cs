using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class TransactionCategoryService(ITransactionCategoryRepository transactionCategoryRepository, IMapper mapper)
    : Core.TransactionCategoryService.TransactionCategoryServiceBase
{
    public override async Task<GetAllTransactionCategoriesResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var transactionCategories = await transactionCategoryRepository.GetAllAsync(request.PageN, request.PageSize);

        return new GetAllTransactionCategoriesResponse
        {
            TransactionCategories = { mapper.Map<IEnumerable<TransactionCategoryModel>>(transactionCategories) }
        };
    }

    public override async Task<TransactionCategoryModel> Add(AddTransactionCategoryRequest request, ServerCallContext context)
    {
        var transactionCategory = mapper.Map<TransactionCategory>(request);

        await transactionCategoryRepository.AddAsync(transactionCategory);

        return mapper.Map<TransactionCategoryModel>(transactionCategory);
    }

    public override async Task<TransactionCategoryModel> GetById(GetTransactionCategoryByIdRequest request, ServerCallContext context)
    {
        var transactionCategory = await transactionCategoryRepository.GetByIdAsync(request.CategoryId);

        if (transactionCategory == null)
            throw new NotFoundException("TransactionCategory not found");

        return mapper.Map<TransactionCategoryModel>(transactionCategory);
    }

    public override async Task<Empty> Update(UpdateTransactionCategoryRequest request, ServerCallContext context)
    {
        var transactionCategory = await transactionCategoryRepository.GetByIdAsync(request.CategoryId);

        if (transactionCategory == null)
            throw new NotFoundException("TransactionCategory not found");

        mapper.Map(request, transactionCategory);

        await transactionCategoryRepository.UpdateAsync(transactionCategory);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteTransactionCategoryRequest request, ServerCallContext context)
    {
        await transactionCategoryRepository.DeleteAsync(request.CategoryId);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteTransactionCategoryBulkRequest request, ServerCallContext context)
    {
        var ids = request.TransactionCategories.Select(t => t.CategoryId).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No transaction categories to delete");
        }
        var transactionCategories = await transactionCategoryRepository.GetByIdsAsync(ids);
        var foundTransactionCategories = transactionCategories.Where(tc => tc != null).ToList();
        if (foundTransactionCategories.Count != ids.Count)
        {
            throw new ValidationException("Some transaction categories not found");
        }
        await transactionCategoryRepository.DeleteRangeAsync(foundTransactionCategories!);
        return new Empty();
    }
    
    public override async Task<Empty> UpdateBulk(UpdateTransactionCategoryBulkRequest request, ServerCallContext context)
    {
        var transactionCategories = request.TransactionCategories.Select(mapper.Map<TransactionCategory>).ToList();
        if (!transactionCategories.Any())
        {
            throw new ValidationException("No transaction categories to update");
        }
        await transactionCategoryRepository.UpdateRangeAsync(transactionCategories);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddTransactionCategoryBulkRequest request, ServerCallContext context)
    {
        var transactionCategories = request.TransactionCategories.Select(mapper.Map<TransactionCategory>).ToList();
        await transactionCategoryRepository.AddRangeAsync(transactionCategories);
        return new Empty();
    }

    public override async Task<GetAllTransactionCategoriesResponse> GetByIds(GetTransactionCategoryByIdsRequest request, ServerCallContext context)
    {
        var transactionCategories = await transactionCategoryRepository.GetByIdsAsync(request.CategoryIds);
        return new GetAllTransactionCategoriesResponse()
        {
            TransactionCategories = { mapper.Map<TransactionCategoryModel>(transactionCategories) }
        };
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var count = await transactionCategoryRepository.GetCountAsync();
        return new CountResponse()
        {
            Count = count
        };
    }
}
