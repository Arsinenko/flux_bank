using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class AccountRepository : GenericRepository<Account, int>, IAccountRepository
{
    public AccountRepository(MyDbContext context) : base(context)
    {
    }
}
