using AutoMapper;
using Core.Exceptions;
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
            throw new NotFoundException("Notification not found");

        return mapper.Map<NotificationModel>(notification);
    }

    public override async Task<Empty> Update(UpdateNotificationRequest request, ServerCallContext context)
    {
        var notification = await notificationRepository.GetByIdAsync(request.NotificationId);

        if (notification == null)
            throw new NotFoundException("Notification not found");

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
            throw new NotFoundException("No notifications to delete");
        }
        var notifications = await notificationRepository.GetByIdsAsync(ids);
        var foundNotifications = notifications.Where(n => n != null).ToList();
        if (foundNotifications.Count != ids.Count)
        {
            throw new NotFoundException("Some notifications not found");
        }
        await notificationRepository.DeleteRangeAsync(foundNotifications!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateNotificationBulkRequest request, ServerCallContext context)
    {
        var notifications = request.Notifications.Select(mapper.Map<Notification>).ToList();
        if (!notifications.Any())
        {
            throw new NotFoundException("No notifications to update");
        }

        await notificationRepository.UpdateRangeAsync(notifications);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddNotificationBulkRequest request, ServerCallContext context)
    {
        var notifications = request.Notifications.Select(mapper.Map<Notification>).ToList();
        await notificationRepository.AddRangeAsync(notifications);
        return new Empty();
    }

    public override async Task<GetAllNotificationsResponse> GetByCustomer(GetNotificationsByCustomerRequest request, ServerCallContext context)
    {
        var notifications = await notificationRepository.FindAsync(n => n.CustomerId == request.CustomerId && n.IsRead == request.IsRead);
        return new GetAllNotificationsResponse()
        {
            Notifications = { mapper.Map<NotificationModel>(notifications) }
        };
    }

    public override async Task<GetAllNotificationsResponse> GetByDateRange(GetByDateRangeRequest request, ServerCallContext context)
    {
        var notifications = await notificationRepository.GetByDateRange(request.From.ToDateTime(),
            request.To.ToDateTime(), request.PageN, request.PageSize);
        return new GetAllNotificationsResponse()
        {
            Notifications = { mapper.Map<NotificationModel>(notifications) }
        };
    }

    public override async Task<GetAllNotificationsResponse> GetByIds(GetNotificationByIdsRequest request, ServerCallContext context)
    {
        var notifications = await notificationRepository.GetByIdsAsync(request.NotificationIds);
        return new GetAllNotificationsResponse()
        {
            Notifications = { mapper.Map<IEnumerable<NotificationModel>>(notifications) }
        };
    }
}
