import decimal
from decimal import Decimal
from typing import List

import grpc
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.transaction_fee_pb2 import *
from api.generated.transaction_fee_pb2_grpc import TransactionFeeServiceStub
from domain.transaction.transaction_fee import TransactionFee
from domain.transaction.transaction_fee_repo import TransactionFeeRepositoryAbc
from mappers.transaction_fee_mapper import TransactionFeeMapper


class TransactionFeeRepository(TransactionFeeRepositoryAbc, BaseGrpcRepository):
    def __init__(self, channel):
        super().__init__(channel)
        self.stub = TransactionFeeServiceStub(channel=self.channel)

    async def get_by_ids(self, ids: List[int]) -> List[TransactionFee]:
        request = GetTransactionFeeByIdsRequest(ids=ids)
        result = await self._execute(self.stub.GetByIds(request))
        if result:
            return TransactionFeeMapper.to_domain_list(result)
        return []



    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        if result:
            return result.count
        return 0
    async def get_total_fee(self) -> decimal.Decimal:
        result = await self._execute(self.stub.GetTotalFee(Empty()))
        if result:
            return Decimal(result.total_fee)
        return Decimal(0)




    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[TransactionFee]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return TransactionFeeMapper.to_domain_list(result)
        return []

    async def get_by_id(self, id: int) -> TransactionFee | None:
        request = GetTransactionFeeByIdRequest(id=id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return TransactionFeeMapper.to_domain(result)
        return None
