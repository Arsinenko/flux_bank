using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class CardRepository : GenericRepository<Card, int>, ICardRepository
{
    public CardRepository(MyDbContext context) : base(context)
    {
    }
}
