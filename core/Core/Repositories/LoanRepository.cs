using Core.Context;
using Core.Interfaces;
using Core.Models;
using Microsoft.EntityFrameworkCore;

namespace Core.Repositories;

public class LoanRepository : GenericRepository<Loan, int>, ILoanRepository
{
    public LoanRepository(MyDbContext context) : base(context)
    {
    }

}
