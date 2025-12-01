using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class NotificationRepository : GenericRepository<Notification, int>, INotificationRepository
{
    public NotificationRepository(MyDbContext context) : base(context)
    {
    }
}
