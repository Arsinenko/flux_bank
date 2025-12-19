from typing import List

import grpc.aio

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.branch_pb2 import *
from api.generated.branch_pb2_grpc import BranchServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.branch.branch import Branch
from domain.branch.branch_repo import BranchRepositoryAbc


class BranchRepository(BranchRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = BranchServiceStub(channel=self.chanel)

    @staticmethod
    def to_domain(model: BranchModel) -> Branch:
        return Branch(
            branch_id=model.branch_id,
            name=model.name,
            city=model.city,
            address=model.address,
            phone=model.phone
        )

    @staticmethod
    def response_to_list(response: GetAllBranchesResponse) -> List[Branch]:
        return [BranchRepository.to_domain(model) for model in response.branches]

    async def get_all(self, page_n: int, page_size: int) -> List[Branch]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, branch_id: int) -> Branch | None:
        request = GetBranchByIdRequest(branch_id=branch_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None
