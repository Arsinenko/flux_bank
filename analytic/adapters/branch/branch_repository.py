from typing import List

from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.branch_pb2 import *
from api.generated.branch_pb2_grpc import BranchServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.branch.branch import Branch
from domain.branch.branch_repo import BranchRepositoryAbc


from mappers.branch_mapper import BranchMapper


class BranchRepository(BranchRepositoryAbc, BaseGrpcRepository):
    def __init__(self, channel):
        super().__init__(channel)
        self.stub = BranchServiceStub(channel=self.channel)

    async def get_by_ids(self, ids: List[int]) -> List[Branch]:
        request = GetBranchByIdsRequest(branch_ids=ids)
        result = await self._execute(self.stub.GetByIds(request))
        if result:
            return BranchMapper.to_domain_list(result.branches)
        return []

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Branch]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return BranchMapper.to_domain_list(result.branches)
        return []

    async def get_by_id(self, branch_id: int) -> Branch | None:
        request = GetBranchByIdRequest(branch_id=branch_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return BranchMapper.to_domain(result)
        return None
    
