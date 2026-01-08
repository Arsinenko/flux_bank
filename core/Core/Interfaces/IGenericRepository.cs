using System.Linq.Expressions;

namespace Core.Interfaces;

public interface IGenericRepository<TEntity, TId> where TEntity : class
{
    public Task<IEnumerable<TEntity>> GetAllAsync(int? pageN, int? pageSize);
    public Task<TEntity?> GetByIdAsync(TId id);
    public Task<int> GetCountAsync();
    public Task<int> GetCountByDateRangeAsync(DateTime fromDate, DateTime ToDate);
    public Task<IEnumerable<TEntity?>> GetByIdsAsync(IEnumerable<TId> ids);
    public Task<IEnumerable<TEntity?>> GetByDateRange(DateTime from, DateTime to, int? pageN, int? pageSize);
    public Task AddAsync(TEntity entity);
    public Task UpdateAsync(TEntity entity);
    public Task DeleteAsync(TId id);
    public Task<IEnumerable<TEntity>> FindAsync(Expression<Func<TEntity, bool>> predicate);
    
    public Task AddRangeAsync(IEnumerable<TEntity> entities);
    public Task UpdateRangeAsync(IEnumerable<TEntity> entities);
    public Task DeleteRangeAsync(IEnumerable<TEntity> entities);
}
