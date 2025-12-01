using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class CustomerRepository : GenericRepository<Customer, int>, ICustomerRepository
{
    public CustomerRepository(MyDbContext context) : base(context)
    {
    }
}
