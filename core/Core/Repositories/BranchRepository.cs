using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class BranchRepository : GenericRepository<Branch, int>, IBranchRepository
{
    public BranchRepository(MyDbContext context) : base(context)
    {
    }
}
