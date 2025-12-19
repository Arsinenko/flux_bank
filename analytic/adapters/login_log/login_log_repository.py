from datetime import datetime

import grpc.aio

from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.login_log_pb2 import *
from api.generated.login_log_pb2_grpc import LoginLogServiceStub
from domain.login_log.login_log_repo import LoginLogRepositoryAbc
from typing import List
from domain.login_log.login_log import LoginLog


class LoginLogRepository(LoginLogRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = LoginLogServiceStub(self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model: LoginLogModel) -> LoginLog:
        return LoginLog(
            log_id=model.log_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            login_time=model.login_time.ToDatetime() if model.HasField("login_time") else None,
            ip_address=model.ip_address if model.HasField("ip_address") else None,
            device_info=model.device_info if model.HasField("device_info") else None
        )

    @staticmethod
    def response_to_list(response: GetAllLoginLogsResponse) -> List[LoginLog]:
        return [LoginLogRepository.to_domain(model) for model in response.login_logs]

    async def get_all(self, page_n: int, page_size: int) -> List[LoginLog]:
        try:
            result = await self.stub.GetAll(GetAllRequest(pageN=page_n, pageSize=page_size))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, log_id: int) -> LoginLog | None:
        try:
            result = await self.stub.GetById(GetLoginLogByIdRequest(log_id=log_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None

    async def get_by_customer(self, customer_id: int) -> List[LoginLog]:
        try:
            result = await self.stub.GetByCustomer(GetLoginLogsByCustomerRequest(customer_id=customer_id))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByCustomer: {err}")
            return []

    async def get_in_time_range(self, start_time: datetime, end_time: datetime) -> List[LoginLog]:
        try:
            request = GetLoginLogsInTimeRangeRequest(
                start_time=start_time.isoformat(),
                end_time=end_time.isoformat()
            )
            result = await self.stub.GetInTimeRange(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetInTimeRange: {err}")
            return []