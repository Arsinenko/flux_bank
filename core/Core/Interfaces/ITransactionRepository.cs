using Core.Models;

namespace Core.Interfaces;

public interface ITransactionRepository : IGenericRepository<Transaction, int>
{
    Task<IEnumerable<Transaction>> GetRevenueAsync(int accountId, DateTime? from, DateTime? to, int? pageN,
        int? pageSize);

    Task<IEnumerable<Transaction>> GetExpensesAsync(int accountId, DateTime? from, DateTime? to, int? pageN,
        int? pageSize);

    Task<bool> MakeTransactionAsync(Transaction transaction);
}
