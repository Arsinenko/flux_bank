from typing import List

import grpc

from api.generated.custom_types_pb2 import GetAllRequest, GetByDateRangeRequest
from api.generated.notification_pb2 import NotificationModel, GetNotificationByIdRequest, GetNotificationsByCustomerRequest
from api.generated.notification_pb2_grpc import NotificationServiceStub
from domain.notification.notification import Notification
from domain.notification.notification_repo import NotificationRepositoryAbc


class NotificationRepository(NotificationRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = NotificationServiceStub(self.chanel)
    
    async def close(self):
        await self.chanel.close()
    
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
    def response_to_list(self, response) -> List[Notification]:
        return [self.to_domain(model) for model in response.notifications]

    async def get_all(self, page_n: int, page_size: int) -> List[Notification]:
        try:
            result = await self.stub.GetAll(GetAllRequest(pageN=page_n, pageSize=page_size))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, notification_id: int) -> Notification | None:
        try:
            result = await self.stub.GetById(GetNotificationByIdRequest(notification_id=notification_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None


    async def get_by_date_range(self, start_date: str, end_date: str) -> List[Notification]:
        try:
            request = GetByDateRangeRequest(fromDate=start_date, toDate=end_date)
            result = await self.stub.GetByDateRange(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByDateRange: {err}")
            return []



    async def get_by_customer(self, customer_id: int) -> List[Notification]:
        try:
            result = await self.stub.GetByCustomer(GetNotificationsByCustomerRequest(customer_id=customer_id))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByCustomer: {err}")
            return []