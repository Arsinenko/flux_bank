using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class NotificationService(INotificationRepository notificationRepository, IMapper mapper)
    : Core.NotificationService.NotificationServiceBase
{
    public override async Task<GetAllNotificationsResponse> GetAll(Empty request, ServerCallContext context)
    {
        var notifications = await notificationRepository.GetAllAsync();

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
}
