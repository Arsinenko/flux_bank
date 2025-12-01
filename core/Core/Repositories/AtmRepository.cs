using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class AtmRepository : GenericRepository<Atm, int>, IAtmRepository
{
    public AtmRepository(MyDbContext context) : base(context)
    {
    }
}
