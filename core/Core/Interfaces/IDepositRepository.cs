using Core.Models;

namespace Core.Interfaces;

public interface IDepositRepository : IGenericRepository<Deposit, int>
{
    Task<int> GetCountByStatus(string status);
}
