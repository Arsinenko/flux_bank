from datetime import datetime

import grpc.aio
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.login_log_pb2 import *
from api.generated.login_log_pb2_grpc import LoginLogServiceStub
from domain.login_log.login_log_repo import LoginLogRepositoryAbc
from typing import List
from domain.login_log.login_log import LoginLog


from mappers.login_log_mapper import LoginLogMapper


class LoginLogRepository(LoginLogRepositoryAbc, BaseGrpcRepository):
    def __init__(self, channel):
        super().__init__(channel)
        self.stub = LoginLogServiceStub(self.channel)

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count


    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[LoginLog]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return LoginLogMapper.to_domain_list(result.login_logs)
        return []

    async def get_by_id(self, log_id: int) -> LoginLog | None:
        request = GetLoginLogByIdRequest(log_id=log_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return LoginLogMapper.to_domain(result)
        return None

    async def get_by_customer(self, customer_id: int) -> List[LoginLog]:
        request = GetLoginLogsByCustomerRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetByCustomer(request))
        if result:
            return LoginLogMapper.to_domain_list(result.login_logs)
        return []

    async def get_in_time_range(self, start_time: datetime, end_time: datetime) -> List[LoginLog]:
        request = GetLoginLogsInTimeRangeRequest(
            start_time=start_time.isoformat(),
            end_time=end_time.isoformat()
        )
        result = await self._execute(self.stub.GetInTimeRange(request))
        if result:
            return LoginLogMapper.to_domain_list(result.login_logs)
        return []
