from typing import List

import grpc.aio
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.atm_pb2 import *
from api.generated.atm_pb2_grpc import AtmServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.atm.atm import Atm
from domain.atm.atm_repo import AtmRepositoryAbc


class AtmRepository(AtmRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = AtmServiceStub(channel=self.chanel)

    @staticmethod
    def to_domain(model: AtmModel) -> Atm:
        return Atm(
            atm_id=model.atm_id,
            location=model.location,
            status=model.status,
            branch_id=model.branch_id
        )

    @staticmethod
    def response_to_list(self, response: GetAllAtmsResponse) -> List[Atm]:
        return [self.to_domain(model) for model in response.atms]

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_count_by_status(self, status: str) -> int:
        result = await self._execute(self.stub.GetCountByStatus(GetAtmsByStatusRequest(status=status)))
        return result.count

    async def get_all(self, page_n: int, page_size: int) -> List[Atm]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, atm_id: int) -> Atm | None:
        request = GetAtmByIdRequest(atm_id=atm_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None

    async def get_by_status(self, status: str) -> List[Atm]:
        request = GetAtmsByStatusRequest(status=status)
        result = await self._execute(self.stub.GetByStatus(request))
        if result:
            return self.response_to_list(result)
        return []


    async def get_by_location_substr(self, sub_str: str) -> List[Atm]:
        request = GetAtmsByLocationSubStrRequest(sub_str=sub_str)
        result = await self._execute(self.stub.GetByLocationSubStr(request))
        if result:
            return self.response_to_list(result)
        return []


    async def get_by_branch(self, branch_id: int) -> List[Atm]:
        request = GetAtmsByBranchRequest(branch_id=branch_id)
        result = await self._execute(self.stub.GetByBranch(request))
        if result:
            return self.response_to_list(result)
        return []
