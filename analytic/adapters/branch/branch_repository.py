from typing import List

import grpc.aio

from api.generated.branch_pb2 import *
from api.generated.branch_pb2_grpc import BranchServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.branch.branch import Branch
from domain.branch.branch_repo import BranchRepositoryAbc


class BranchRepository(BranchRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = BranchServiceStub(channel=self.chanel)
        
    async def close(self):
        await self.chanel.close()
    
    @staticmethod
    def to_domain(model: BranchModel) -> Branch:
        return Branch(
            branch_id=model.branch_id,
            name=model.name,
            city=model.city,
            address=model.address,
            phone=model.phone
        )

    async def get_all(self, page_n: int, page_size: int) -> List[Branch]:
        try:
            result = await self.stub.GetAll(GetAllRequest(pageN=page_n, pageSize=page_size))
            return [self.to_domain(model) for model in result.branches]
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, branch_id: int) -> Branch | None:
        try:
            result = await self.stub.GetById(GetBranchByIdRequest(branch_id=branch_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None