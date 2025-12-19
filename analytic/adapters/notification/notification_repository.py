from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest, GetByDateRangeRequest
from api.generated.notification_pb2 import NotificationModel, GetNotificationByIdRequest, GetNotificationsByCustomerRequest, GetAllNotificationsResponse
from api.generated.notification_pb2_grpc import NotificationServiceStub
from domain.notification.notification import Notification
from domain.notification.notification_repo import NotificationRepositoryAbc


class NotificationRepository(NotificationRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = NotificationServiceStub(self.chanel)

    @staticmethod
    def to_domain(model: NotificationModel):
        return Notification(
            notification_id=model.notification_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            message=model.message if model.HasField("message") else None,
            created_at=model.created_at.ToDatetime() if model.HasField("created_at") else None,
            is_read=model.is_read if model.HasField("is_read") else None
        )

    @staticmethod
    def response_to_list(response: GetAllNotificationsResponse) -> List[Notification]:
        return [NotificationRepository.to_domain(model) for model in response.notifications]

    async def get_all(self, page_n: int, page_size: int) -> List[Notification]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, notification_id: int) -> Notification | None:
        request = GetNotificationByIdRequest(notification_id=notification_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None

    async def get_by_date_range(self, start_date: str, end_date: str) -> List[Notification]:
        request = GetByDateRangeRequest(fromDate=start_date, toDate=end_date)
        result = await self._execute(self.stub.GetByDateRange(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_customer(self, customer_id: int) -> List[Notification]:
        request = GetNotificationsByCustomerRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetByCustomer(request))
        if result:
            return self.response_to_list(result)
        return []
