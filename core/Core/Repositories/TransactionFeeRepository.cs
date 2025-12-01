using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class TransactionFeeRepository : GenericRepository<TransactionFee, int>, ITransactionFeeRepository
{
    public TransactionFeeRepository(MyDbContext context) : base(context)
    {
    }
}
