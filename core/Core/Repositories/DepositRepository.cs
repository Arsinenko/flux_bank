using Core.Context;
using Core.Interfaces;
using Core.Models;
using Microsoft.EntityFrameworkCore;

namespace Core.Repositories;

public class DepositRepository : GenericRepository<Deposit, int>, IDepositRepository
{
    public DepositRepository(MyDbContext context) : base(context)
    {
    }
}
