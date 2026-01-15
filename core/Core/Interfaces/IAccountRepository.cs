using Core.Models;

namespace Core.Interfaces;

public interface IAccountRepository : IGenericRepository<Account, int>
{
    Task<IEnumerable<Account>> GetByCustomerIdAsync(int customerId);
}
