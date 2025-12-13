from typing import List

import grpc.aio

from api.generated.atm_pb2 import *
from api.generated.atm_pb2_grpc import AtmServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.atm.atm import Atm
from domain.atm.atm_repo import AtmRepositoryAbc


class AtmRepository(AtmRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = AtmServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

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

    async def get_all(self, page_n: int, page_size: int) -> List[Atm]:
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, atm_id: int) -> Atm | None:
        try:
            result = await self.stub.GetById(GetAtmByIdRequest(atm_id=atm_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None

    async def get_by_status(self, status: str) -> List[Atm]:
        try:
            result = await self.stub.GetByStatus(GetAtmsByStatusRequest(status=status))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByStatus: {err}")
            return []


    async def get_by_location_substr(self, sub_str: str) -> List[Atm]:
        try:
            result = await self.stub.GetByLocationSubStr(GetAtmsByLocationSubStrRequest(sub_str=sub_str))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByLocationSubStr: {err}")
            return []


    async def get_by_branch(self, branch_id: int) -> List[Atm]:
        try:
            result = await self.stub.GetByBranch(GetAtmsByBranchRequest(branch_id=branch_id))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByBranch: {err}")
            return []
