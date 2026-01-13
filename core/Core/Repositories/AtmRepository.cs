using Core.Context;
using Core.Interfaces;
using Core.Models;
using Microsoft.EntityFrameworkCore;

namespace Core.Repositories;

public class AtmRepository : GenericRepository<Atm, int>, IAtmRepository
{
    public AtmRepository(MyDbContext context) : base(context)
    {
    }

    public async Task<int> GetCountByStatusAsync(string status)
    {
        return await DbSet.Where(a => a.Status == status).CountAsync();
    }
}
