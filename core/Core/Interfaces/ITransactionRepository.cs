using Core.Models;

namespace Core.Interfaces;

public interface ITransactionRepository : IGenericRepository<Transaction, int>
{
    Task<IEnumerable<Transaction>> GetRevenueAsync(int accountId, DateTime? from, DateTime? to, int? pageN,
        int? pageSize);
    Task<int> GetCountRevenuesAsync(int accountId, DateTime? from, DateTime? to);

    Task<IEnumerable<Transaction>> GetExpensesAsync(int accountId, DateTime? from, DateTime? to, int? pageN,
        int? pageSize);

    Task<int> GetCountExpensesAsync(int accountId, DateTime? from, DateTime? to);

    Task<decimal> GetTotalAmountAsync();
}
