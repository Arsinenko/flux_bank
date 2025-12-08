using Core.Context;
using Core.Interfaces;
using Core.Models;

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
}
