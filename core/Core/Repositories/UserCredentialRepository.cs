using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class UserCredentialRepository : GenericRepository<UserCredential, int>, IUserCredentialRepository
{
    public UserCredentialRepository(MyDbContext context) : base(context)
    {
    }
}
