using Core.Models;

namespace Core.Interfaces;

public interface ICustomerRepository : IGenericRepository<Customer, int>
{
    Task<IEnumerable<Customer>> GetBySubstring(string subStr, int? pageN, int? pageSize, string order, bool desc = false);
    Task<int> GetCountBySubstring(string subStr);
}
