using System.Linq.Expressions;
using Core.Context;
using Core.Exceptions;
using Core.Interfaces;
using Microsoft.EntityFrameworkCore;

namespace Core.Repositories;

public class GenericRepository<TEntity, TId> 
    : IGenericRepository<TEntity, TId>
    where TEntity : class
{
    protected readonly MyDbContext Context;
    protected readonly DbSet<TEntity> DbSet;

    public GenericRepository(MyDbContext context)
    {
        Context = context;
        DbSet = context.Set<TEntity>();
    }

    public virtual async Task<IEnumerable<TEntity>> GetAllAsync(int? pageN, int? pageSize)
    {
        IQueryable<TEntity> query = DbSet;
        
        if (pageN != 0 && pageSize != 0)
        {
            if (pageN <= 0 || pageSize <= 0) throw new ArgumentException("pageN and pageSize must be greater than 0");
            var keyName = GetEntityKey();
            query = query.OrderBy(e => EF.Property<TId>(e, keyName)).Skip((pageN.Value - 1) * pageSize.Value).Take(pageSize.Value);
        }
        
        return await query.ToListAsync();
    }

    public virtual async Task<TEntity?> GetByIdAsync(TId id)
    {
        return await DbSet.FindAsync(id);
    }

    public async Task<int> GetCountAsync()
    {
        return await DbSet.CountAsync();
    }

    public async Task<int> GetCountByDateRangeAsync(DateTime fromDate, DateTime ToDate)
    {
        return await DbSet.Where(e => 
            EF.Property<DateTime>(e, "CreatedAt") >= fromDate && EF.Property<DateTime>(e, "CreatedAt") <= ToDate).CountAsync();
    }

    public async Task<IEnumerable<TEntity?>> GetByIdsAsync(IEnumerable<TId> ids)
    {
        var keyName = GetEntityKey();
        return await DbSet.Where(e => ids.Contains(EF.Property<TId>(e, keyName))).ToListAsync();
    }

    public async Task<IEnumerable<TEntity?>> GetByDateRange(DateTime from, DateTime to, int? pageN, int? pageSize)
    {
        if (typeof(TEntity).GetProperty("CreatedAt") == null)
        {
            throw new InvalidOperationException($"Entity type '{typeof(TEntity).Name}' does not have a 'CreatedAt' property.");
        }

        IQueryable<TEntity> query = DbSet.Where(e => EF.Property<DateTime>(e, "CreatedAt") >= from && EF.Property<DateTime>(e, "CreatedAt") <= to);
        if (pageN.HasValue && pageSize.HasValue)
        {
            if (pageN <= 0 || pageSize <= 0) throw new ArgumentException("pageN and pageSize must be greater than 0");
            var keyName = GetEntityKey();
            query = query.OrderBy(e => EF.Property<TId>(e, keyName)).Skip((pageN.Value - 1) * pageSize.Value).Take(pageSize.Value);
        }
        return await query.ToListAsync();
    }


    public virtual async Task AddAsync(TEntity entity)
    {
        await DbSet.AddAsync(entity);
        await Context.SaveChangesAsync();
    }

    public virtual async Task UpdateAsync(TEntity entity)
    {
        DbSet.Attach(entity);
        Context.Entry(entity).State = EntityState.Modified;
        await Context.SaveChangesAsync();
    }

    public virtual async Task DeleteAsync(TId id)
    {
        var entity = await GetByIdAsync(id);
        if (entity == null)
            throw new NotFoundException($"Entity {entity} not found");
        DbSet.Remove(entity);
        await Context.SaveChangesAsync();
    }

    public virtual async Task<IEnumerable<TEntity>> FindAsync(Expression<Func<TEntity, bool>> predicate)
    {
        return await DbSet.Where(predicate).ToListAsync();
    }



    public virtual async Task AddRangeAsync(IEnumerable<TEntity> entities)
    {
        await DbSet.AddRangeAsync(entities);
        await Context.SaveChangesAsync();
    }

    public virtual async Task UpdateRangeAsync(IEnumerable<TEntity> entities)
    {
        DbSet.UpdateRange(entities);
        await Context.SaveChangesAsync();
    }

    public virtual async Task DeleteRangeAsync(IEnumerable<TEntity> entities)
    {
        DbSet.RemoveRange(entities);
        await Context.SaveChangesAsync();
    }

    public string GetEntityKey()
    {
        var entityType = Context.Model.FindEntityType(typeof(TEntity));
        if (entityType == null)
            throw new InvalidOperationException($"Entity type '{typeof(TEntity).Name}' not found in the model.");
    
        var primaryKey = entityType.FindPrimaryKey();
        if (primaryKey == null)
            throw new InvalidOperationException($"Primary key not found for entity type '{typeof(TEntity).Name}'.");

        var keyName = primaryKey.Properties.First().Name;
        return keyName;
    }
}
