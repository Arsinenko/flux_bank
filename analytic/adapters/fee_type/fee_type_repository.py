from typing import List

import grpc

from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.fee_type_pb2 import *
from api.generated.fee_type_pb2_grpc import FeeTypeServiceStub
from domain.fee_type.fee_type import FeeType
from domain.fee_type.fee_type_repo import FeeTypeRepositoryAbc


class FeeTypeRepository(FeeTypeRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = FeeTypeServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model: FeeTypeModel) -> FeeType:
        return FeeType(
            fee_id=model.fee_id,
            name=model.name if model.HasField("name") else None,
            description=model.description if model.HasField("description") else None
        )

    @staticmethod
    def response_to_list(response: GetAllFeeTypesResponse) -> List[FeeType]:
        return [FeeTypeRepository.to_domain(model) for model in response.fee_types]

    async def get_all(self, page_n: int, page_size: int) -> List[FeeType]:
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, fee_id: int) -> FeeType | None:
        try:
            result = await self.stub.GetById(GetFeeTypeByIdRequest(fee_id=fee_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None
