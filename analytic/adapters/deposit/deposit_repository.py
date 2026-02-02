from decimal import Decimal
from typing import List

import grpc
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.deposit_pb2 import *
from api.generated.deposit_pb2_grpc import DepositServiceStub
from domain.deposit.deposit import Deposit
from domain.deposit.deposit_repo import DepositRepositoryAbc


from mappers.deposit_mapper import DepositMapper


class DepositRepository(DepositRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = DepositServiceStub(channel=self.channel)

    async def get_by_ids(self, ids: List[int]):
        request = GetDepositByIdsRequest(deposit_ids=ids)
        result = await self._execute(self.stub.GetByIds(request))
        if result:
            return DepositMapper.to_domain_list(result.deposits)
        return []


    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_count_by_status(self, status: str):
        result = await self._execute(self.stub.GetCountByStatus(GetDepositCountByStatusRequest(status=status)))
        return result.count


    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Deposit]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return DepositMapper.to_domain_list(result.deposits)
        return []

    async def get_by_id(self, deposit_id: int) -> Deposit | None:
        request = GetDepositByIdRequest(deposit_id=deposit_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return DepositMapper.to_domain(result)
        return None

    async def get_by_customer_id(self, customer_id: int) -> List[Deposit]:
        request = GetDepositsByCustomerRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetByCustomer(request))
        if result:
            return DepositMapper.to_domain_list(result.deposits)
        return []
