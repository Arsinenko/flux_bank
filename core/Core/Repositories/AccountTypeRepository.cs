using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class AccountTypeRepository : GenericRepository<AccountType, int>, IAccountTypeRepository
{
    public AccountTypeRepository(MyDbContext context) : base(context)
    {
    }
}
