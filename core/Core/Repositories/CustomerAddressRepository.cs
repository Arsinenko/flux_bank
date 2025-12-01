using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class CustomerAddressRepository : GenericRepository<CustomerAddress, int>, ICustomerAddressRepository
{
    public CustomerAddressRepository(MyDbContext context) : base(context)
    {
    }
}
