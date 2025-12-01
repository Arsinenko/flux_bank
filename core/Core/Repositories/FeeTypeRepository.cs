using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class FeeTypeRepository : GenericRepository<FeeType, int>, IFeeTypeRepository
{
    public FeeTypeRepository(MyDbContext context) : base(context)
    {
    }
}
