from api.generated.notification_analytic_pb2_grpc import NotificationAnalyticServiceServicer
from api.generated.notification_pb2 import GetNotificationByIdRequest, GetNotificationByIdsRequest, GetAllNotificationsResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.notification.notification_repo import NotificationRepositoryAbc
from mappers.notification_mapper import NotificationMapper
from google.protobuf.empty_pb2 import Empty


class NotificationAnalyticService(NotificationAnalyticServiceServicer):
    def __init__(self, notification_repo: NotificationRepositoryAbc):
        self.notification_repo = notification_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.notification_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllNotificationsResponse(notifications=NotificationMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetNotificationByIdRequest, context):
        result = await self.notification_repo.get_by_id(request.notification_id)
        return NotificationMapper.to_model(result) if result else None

    async def ProcessGetByIds(self, request: GetNotificationByIdsRequest, context):
        result = await self.notification_repo.get_by_ids(request.notification_ids)
        return GetAllNotificationsResponse(notifications=NotificationMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.notification_repo.get_count()
        return CountResponse(count=result)
