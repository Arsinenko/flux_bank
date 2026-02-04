from typing import List

import grpc
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.fee_type_pb2 import *
from api.generated.fee_type_pb2_grpc import FeeTypeServiceStub
from domain.fee_type.fee_type import FeeType
from domain.fee_type.fee_type_repo import FeeTypeRepositoryAbc


from mappers.fee_type_mapper import FeeTypeMapper


class FeeTypeRepository(FeeTypeRepositoryAbc, BaseGrpcRepository):
    def __init__(self, channel):
        super().__init__(channel)
        self.stub = FeeTypeServiceStub(channel=self.channel)

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[FeeType]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return FeeTypeMapper.to_domain_list(result.fee_types)
        return []

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count


    async def get_by_id(self, fee_id: int) -> FeeType | None:
        request = GetFeeTypeByIdRequest(fee_id=fee_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return FeeTypeMapper.to_domain(result)
        return None
