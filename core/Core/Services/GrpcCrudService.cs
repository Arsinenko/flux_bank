using AutoMapper;
using Core.Interfaces;
using Grpc.Core;

namespace Core.Services;

public abstract class GrpcCrudService<TEntity, TModel, TCreateRequest, TUpdateRequest, TKey>
    where TEntity : class  
    where TModel : class
{
    protected readonly IGenericRepository<TEntity, TKey> Repository;
    protected readonly IMapper Mapper;

    protected GrpcCrudService(IGenericRepository<TEntity, TKey> repository, IMapper mapper)
    {
        Repository = repository;
        Mapper = mapper;
    }

    protected async Task<TModel> AddAsync(TCreateRequest request)
    {
        var entity = Mapper.Map<TEntity>(request);
        await Repository.AddAsync(entity);
        return Mapper.Map<TModel>(entity);
    }

    protected async Task<TModel> GetByIdAsync(TKey id)
    {
        var entity = await Repository.GetByIdAsync(id);
        if (entity == null)
            throw new RpcException(new Status(StatusCode.NotFound, $"{typeof(TEntity).Name} not found"));

        return Mapper.Map<TModel>(entity);
    }

    protected async Task<IEnumerable<TModel>> GetAllAsync()
    {
        var entities = await Repository.GetAllAsync();
        return Mapper.Map<IEnumerable<TModel>>(entities);
    }

    protected async Task UpdateAsync(TKey id, TUpdateRequest request)
    {
        var entity = await Repository.GetByIdAsync(id);
        if (entity == null)
            throw new RpcException(new Status(StatusCode.NotFound, $"{typeof(TEntity).Name} not found"));

        Mapper.Map(request, entity);
        await Repository.UpdateAsync(entity);
    }

    protected async Task DeleteAsync(TKey id)
    {
        await Repository.DeleteAsync(id);
    }
}