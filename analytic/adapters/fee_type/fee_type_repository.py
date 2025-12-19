from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.fee_type_pb2 import *
from api.generated.fee_type_pb2_grpc import FeeTypeServiceStub
from domain.fee_type.fee_type import FeeType
from domain.fee_type.fee_type_repo import FeeTypeRepositoryAbc


class FeeTypeRepository(FeeTypeRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = FeeTypeServiceStub(channel=self.chanel)

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
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, fee_id: int) -> FeeType | None:
        request = GetFeeTypeByIdRequest(fee_id=fee_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None
