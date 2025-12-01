using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class LoginLogRepository : GenericRepository<LoginLog, int>, ILoginLogRepository
{
    public LoginLogRepository(MyDbContext context) : base(context)
    {
    }
}
