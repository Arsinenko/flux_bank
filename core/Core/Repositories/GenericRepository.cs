using System.Linq.Expressions;
using Core.Context;
using Core.Interfaces;
using Microsoft.EntityFrameworkCore;

namespace Core.Repositories;

public class GenericRepository<TEntity, TId> 
    : IGenericRepository<TEntity, TId>
    where TEntity : class
{
    protected readonly MyDbContext _context;
    protected readonly DbSet<TEntity> _dbSet;

    public GenericRepository(MyDbContext context)
    {
        _context = context;
        _dbSet = context.Set<TEntity>();
    }

    public virtual async Task<IEnumerable<TEntity>> GetAllAsync(int? pageN, int? pageSize)
    {
        IQueryable<TEntity> query = _dbSet;
        
        if (pageN.HasValue && pageSize.HasValue)
        {
            if (pageN <= 0 || pageSize <= 0) throw new ArgumentException("pageN and pageSize must be greater than 0");
            var keyName = GetEntityKey();
            query = query.OrderBy(e => EF.Property<TId>(e, keyName)).Skip((pageN.Value - 1) * pageSize.Value).Take(pageSize.Value);
        }
        
        return await query.ToListAsync();
    }

    public virtual async Task<TEntity?> GetByIdAsync(TId id)
    {
        return await _dbSet.FindAsync(id);
    }

    public async Task<IEnumerable<TEntity?>> GetByIdsAsync(IEnumerable<TId> ids)
    {
        var keyName = GetEntityKey();
        return await _dbSet.Where(e => ids.Contains(EF.Property<TId>(e, keyName))).ToListAsync();
    }


    public virtual async Task AddAsync(TEntity entity)
    {
        await _dbSet.AddAsync(entity);
        await _context.SaveChangesAsync();
    }

    public virtual async Task UpdateAsync(TEntity entity)
    {
        _dbSet.Attach(entity);
        _context.Entry(entity).State = EntityState.Modified;
        await _context.SaveChangesAsync();
    }

    public virtual async Task DeleteAsync(TId id)
    {
        var entity = await GetByIdAsync(id);
        if (entity != null)
        {
            _dbSet.Remove(entity);
            await _context.SaveChangesAsync();
        }
    }

    public virtual async Task<IEnumerable<TEntity>> FindAsync(Expression<Func<TEntity, bool>> predicate)
    {
        return await _dbSet.Where(predicate).ToListAsync();
    }

    public virtual async Task AddRangeAsync(IEnumerable<TEntity> entities)
    {
        await _dbSet.AddRangeAsync(entities);
        await _context.SaveChangesAsync();
    }

    public virtual async Task UpdateRangeAsync(IEnumerable<TEntity> entities)
    {
        _dbSet.UpdateRange(entities);
        await _context.SaveChangesAsync();
    }

    public virtual async Task DeleteRangeAsync(IEnumerable<TEntity> entities)
    {
        _dbSet.RemoveRange(entities);
        await _context.SaveChangesAsync();
    }

    public string GetEntityKey()
    {
        var entityType = _context.Model.FindEntityType(typeof(TEntity));
        var keyName = entityType.FindPrimaryKey().Properties.First().Name;
        return keyName;
    }
}
