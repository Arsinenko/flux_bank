using Core.Context;
using Core.Interfaces;
using Core.Models;
using Microsoft.EntityFrameworkCore;

namespace Core.Repositories;

public class AccountRepository : GenericRepository<Account, int>, IAccountRepository
{
    public AccountRepository(MyDbContext context) : base(context)
    {
    }

    public Task<IEnumerable<Account>> GetByCustomerIdAsync(int customerId)
    {
        return FindAsync(account => account.CustomerId == customerId);
    }

    public async Task<int> GetCountByStatusAsync(bool status)
    {
        return await DbSet.Where(a => a.IsActive == status).CountAsync();
    }

    public async Task<decimal?> GetTotalBalanceAsync()
    {
        return await DbSet.Select(a => a.Balance).SumAsync();
    }
}
