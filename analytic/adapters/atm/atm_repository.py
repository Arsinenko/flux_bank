from typing import List

import grpc.aio
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.atm_pb2 import *
from api.generated.atm_pb2_grpc import AtmServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.atm.atm import Atm
from domain.atm.atm_repo import AtmRepositoryAbc


from mappers.atm_mapper import AtmMapper


class AtmRepository(AtmRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = AtmServiceStub(channel=self.chanel)

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_count_by_status(self, status: str) -> int:
        result = await self._execute(self.stub.GetCountByStatus(GetAtmsByStatusRequest(status=status)))
        return result.count

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Atm]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return AtmMapper.to_domain_list(result.atms)
        return []

    async def get_by_id(self, atm_id: int) -> Atm | None:
        request = GetAtmByIdRequest(atm_id=atm_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return AtmMapper.to_domain(result)
        return None

    async def get_by_status(self, status: str) -> List[Atm]:
        request = GetAtmsByStatusRequest(status=status)
        result = await self._execute(self.stub.GetByStatus(request))
        if result:
            return AtmMapper.to_domain_list(result.atms)
        return []


    async def get_by_location_substr(self, sub_str: str) -> List[Atm]:
        request = GetAtmsByLocationSubStrRequest(sub_str=sub_str)
        result = await self._execute(self.stub.GetByLocationSubStr(request))
        if result:
            return AtmMapper.to_domain_list(result.atms)
        return []


    async def get_by_branch(self, branch_id: int) -> List[Atm]:
        request = GetAtmsByBranchRequest(branch_id=branch_id)
        result = await self._execute(self.stub.GetByBranch(request))
        if result:
            return AtmMapper.to_domain_list(result.atms)
        return []
