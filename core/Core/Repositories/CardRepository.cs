using Core.Context;
using Core.Interfaces;
using Core.Models;
using Microsoft.EntityFrameworkCore;

namespace Core.Repositories;

public class CardRepository : GenericRepository<Card, int>, ICardRepository
{
    public CardRepository(MyDbContext context) : base(context)
    {
    }

    public async Task<int> GetCountByStatusAsync(string status)
    {
        return await DbSet.Where(c => c.Status == status).CountAsync();
    }
}
