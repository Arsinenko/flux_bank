from datetime import datetime

import grpc.aio

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.login_log_pb2 import *
from api.generated.login_log_pb2_grpc import LoginLogServiceStub
from domain.login_log.login_log_repo import LoginLogRepositoryAbc
from typing import List
from domain.login_log.login_log import LoginLog


class LoginLogRepository(LoginLogRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = LoginLogServiceStub(self.chanel)

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
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, log_id: int) -> LoginLog | None:
        request = GetLoginLogByIdRequest(log_id=log_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None

    async def get_by_customer(self, customer_id: int) -> List[LoginLog]:
        request = GetLoginLogsByCustomerRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetByCustomer(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_in_time_range(self, start_time: datetime, end_time: datetime) -> List[LoginLog]:
        request = GetLoginLogsInTimeRangeRequest(
            start_time=start_time.isoformat(),
            end_time=end_time.isoformat()
        )
        result = await self._execute(self.stub.GetInTimeRange(request))
        if result:
            return self.response_to_list(result)
        return []
