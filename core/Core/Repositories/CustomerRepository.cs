using Core.Context;
using Core.Interfaces;
using Core.Models;
using Microsoft.EntityFrameworkCore;

namespace Core.Repositories;

public class CustomerRepository : GenericRepository<Customer, int>, ICustomerRepository
{
    public CustomerRepository(MyDbContext context) : base(context)
    {
    }

    public async Task<IEnumerable<Customer>> GetBySubstring(string subStr, int? pageN, int? pageSize, string order,
        bool desc = false)
    {
        var query = DbSet.Where(c =>
            c.FirstName.Contains(subStr) || c.LastName.Contains(subStr) || c.Email.Contains(subStr) ||
            c.Phone.Contains(subStr));

        query = order switch
        {
            "FirstName" => desc ? query.OrderByDescending(c => c.FirstName) : query.OrderBy(c => c.FirstName),
            "LastName" => desc ? query.OrderByDescending(c => c.LastName) : query.OrderBy(c => c.LastName),
            "Email" => desc ? query.OrderByDescending(c => c.Email) : query.OrderBy(c => c.Email),
            "Phone" => desc ? query.OrderByDescending(c => c.Phone) : query.OrderBy(c => c.Phone),
            "BirthDate" => desc ? query.OrderByDescending(c => c.BirthDate) : query.OrderBy(c => c.BirthDate),
            "CreatedAt" => desc ? query.OrderByDescending(c => c.CreatedAt) : query.OrderBy(c => c.CreatedAt),
            _ => query.OrderBy(c => c.CustomerId)
        };
        if (pageN.HasValue && pageSize.HasValue)
        {
            if (pageN <= 0 || pageSize <= 0) throw new ArgumentException("pageN and pageSize must be greater than 0");
            query = query.Skip((pageN.Value - 1) * pageSize.Value).Take(pageSize.Value);
        }

        return await query.ToListAsync();


    }

    public async Task<int> GetCountBySubstring(string subStr)
    {
        return await DbSet.Where(c => c.FirstName.Contains(subStr) || c.LastName.Contains(subStr) ||
                               c.Email.Contains(subStr) ||
                               c.Phone.Contains(subStr)).CountAsync();
    }
}
