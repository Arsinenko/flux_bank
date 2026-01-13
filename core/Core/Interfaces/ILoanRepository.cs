using Core.Models;

namespace Core.Interfaces;

public interface ILoanRepository : IGenericRepository<Loan, int>
{
    Task<int> GetCountByStatus(string status);
}
