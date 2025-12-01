using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class TransactionRepository : GenericRepository<Transaction, int>, ITransactionRepository
{
    public TransactionRepository(MyDbContext context) : base(context)
    {
    }
}
