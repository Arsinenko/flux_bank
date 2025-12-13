using Core.Context;
using Core.Interfaces;
using Core.Models;
using Microsoft.EntityFrameworkCore;

namespace Core.Repositories;

public class TransactionRepository : GenericRepository<Transaction, int>, ITransactionRepository
{
    public TransactionRepository(MyDbContext context) : base(context)
    {
    }

    public async Task<IEnumerable<Transaction>> GetRevenueAsync(int accountId, DateTime? from, DateTime? to, int? pageN, int? pageSize)
    {
        IQueryable<Transaction> query = DbSet.Where(t => t.TargetAccount == accountId);
        query = TransactionsPaged(from, to, pageN, pageSize, query);
        return await query.ToListAsync();
    }


    public async Task<IEnumerable<Transaction>> GetExpensesAsync(int accountId, DateTime? from, DateTime? to, int? pageN, int? pageSize)
    {
        IQueryable<Transaction> query = DbSet.Where(t => t.SourceAccount == accountId);
        query = TransactionsPaged(from, to, pageN, pageSize, query);
        return await query.ToListAsync();
    }

    private IQueryable<Transaction> TransactionsPaged(DateTime? from, DateTime? to, int? pageN, int? pageSize, IQueryable<Transaction> query)
    {
        if (from.HasValue && to.HasValue)
        {
            query = query.Where(t => t.CreatedAt >= from.Value && t.CreatedAt <= to.Value);
        }

        if (pageN.HasValue && pageSize.HasValue)
        {
            if (pageN <= 0 || pageSize <= 0)
            {
                throw new ArgumentException("pageN and pageSize must be greater than 0");
            }
            var keyName = GetEntityKey();
            query = query.OrderBy(e => EF.Property<int>(e, keyName)).Skip((pageN.Value - 1) * pageSize.Value).Take(pageSize.Value);
        }
        return query;
    }
}
