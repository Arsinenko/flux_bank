using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class ExchangeRateRepository : GenericRepository<ExchangeRate, int>, IExchangeRateRepository
{
    public ExchangeRateRepository(MyDbContext context) : base(context)
    {
    }
}
