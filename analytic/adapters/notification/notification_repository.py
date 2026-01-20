from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.custom_types_pb2 import GetAllRequest, GetByDateRangeRequest
from api.generated.notification_pb2 import NotificationModel, GetNotificationByIdRequest, GetNotificationsByCustomerRequest, GetAllNotificationsResponse
from api.generated.notification_pb2_grpc import NotificationServiceStub
from domain.notification.notification import Notification
from domain.notification.notification_repo import NotificationRepositoryAbc


from mappers.notification_mapper import NotificationMapper


class NotificationRepository(NotificationRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = NotificationServiceStub(self.chanel)

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Notification]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return NotificationMapper.to_domain_list(result.notifications)
        return []

    async def get_by_id(self, notification_id: int) -> Notification | None:
        request = GetNotificationByIdRequest(notification_id=notification_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return NotificationMapper.to_domain(result)
        return None

    async def get_by_date_range(self, start_date: str, end_date: str) -> List[Notification]:
        request = GetByDateRangeRequest(fromDate=start_date, toDate=end_date)
        result = await self._execute(self.stub.GetByDateRange(request))
        if result:
            return NotificationMapper.to_domain_list(result.notifications)
        return []

    async def get_by_customer(self, customer_id: int) -> List[Notification]:
        request = GetNotificationsByCustomerRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetByCustomer(request))
        if result:
            return NotificationMapper.to_domain_list(result.notifications)
        return []
