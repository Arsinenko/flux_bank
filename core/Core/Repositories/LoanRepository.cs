using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class LoanRepository : GenericRepository<Loan, int>, ILoanRepository
{
    public LoanRepository(MyDbContext context) : base(context)
    {
    }
}
