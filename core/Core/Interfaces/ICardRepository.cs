using Core.Models;

namespace Core.Interfaces;

public interface ICardRepository : IGenericRepository<Card, int>
{
    Task<int> GetCountByStatusAsync(string status);
}
