using Core.Models;

namespace Core.Interfaces;

public interface IAtmRepository : IGenericRepository<Atm, int>
{
    Task<int> GetCountByStatusAsync(string status);
}
