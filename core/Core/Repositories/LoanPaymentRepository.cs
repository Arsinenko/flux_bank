using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class LoanPaymentRepository : GenericRepository<LoanPayment, int>, ILoanPaymentRepository
{
    public LoanPaymentRepository(MyDbContext context) : base(context)
    {
    }
}
