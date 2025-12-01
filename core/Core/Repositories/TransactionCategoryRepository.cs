using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class TransactionCategoryRepository : GenericRepository<TransactionCategory, int>, ITransactionCategoryRepository
{
    public TransactionCategoryRepository(MyDbContext context) : base(context)
    {
    }
}
