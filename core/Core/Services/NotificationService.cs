using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class NotificationService(INotificationRepository notificationRepository, IMapper mapper)
    : Core.NotificationService.NotificationServiceBase
{
    public override async Task<GetAllNotificationsResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var notifications = await notificationRepository.GetAllAsync(request.PageN, request.PageSize);

        return new GetAllNotificationsResponse
        {
            Notifications = { mapper.Map<IEnumerable<NotificationModel>>(notifications) }
        };
    }

    public override async Task<NotificationModel> Add(AddNotificationRequest request, ServerCallContext context)
    {
        var notification = mapper.Map<Notification>(request);

        await notificationRepository.AddAsync(notification);

        return mapper.Map<NotificationModel>(notification);
    }

    public override async Task<NotificationModel> GetById(GetNotificationByIdRequest request, ServerCallContext context)
    {
        var notification = await notificationRepository.GetByIdAsync(request.NotificationId);

        if (notification == null)
            throw new RpcException(new Status(StatusCode.NotFound, "Notification not found"));

        return mapper.Map<NotificationModel>(notification);
    }

    public override async Task<Empty> Update(UpdateNotificationRequest request, ServerCallContext context)
    {
        var notification = await notificationRepository.GetByIdAsync(request.NotificationId);

        if (notification == null)
            throw new RpcException(new Status(StatusCode.NotFound, "Notification not found"));

        mapper.Map(request, notification);

        await notificationRepository.UpdateAsync(notification);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteNotificationRequest request, ServerCallContext context)
    {
        await notificationRepository.DeleteAsync(request.NotificationId);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteNotificationBulkRequest request, ServerCallContext context)
    {
        var ids = request.Notifications.Select(n => n.NotificationId).ToList();
        if (ids.Count == 0)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No notifications to delete"));
        }
        var notifications = await notificationRepository.GetByIdsAsync(ids);
        var foundNotifications = notifications.Where(n => n != null).ToList();
        if (foundNotifications.Count != ids.Count)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Some notifications not found"));
        }
        await notificationRepository.DeleteRangeAsync(foundNotifications!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateNotificationBulkRequest request, ServerCallContext context)
    {
        var notifications = request.Notifications.Select(mapper.Map<Notification>).ToList();
        if (!notifications.Any())
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No notifications to update"));
        }

        await notificationRepository.UpdateRangeAsync(notifications);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddNotificationBulkRequest request, ServerCallContext context)
    {
        var notifications = request.Notifications.Select(mapper.Map<Notification>).ToList();
        try
        {
            await notificationRepository.AddRangeAsync(notifications);
            return new Empty();
        }
        catch (Exception e)
        {
            throw new RpcException(new Status(StatusCode.Internal, e.Message));
        }
    }
}
