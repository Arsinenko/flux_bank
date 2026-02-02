from typing import List

import grpc.aio
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.account_type_pb2 import *
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.account_type_pb2 import *
from api.generated.account_type_pb2_grpc import AccountTypeServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.account.account_type import AccountType
from domain.account.account_type_repo import AccountTypeRepositoryAbc


class AccountTypeRepository(AccountTypeRepositoryAbc, BaseGrpcRepository):
    async def get_by_ids(self, ids: List[int]) -> List[AccountType]:
        result = await self._execute(self.stub.GetByIds(GetAccountTypeByIdsRequest(type_ids=ids)))
        return [self.to_domain(model) for model in result.account_types]

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count


    def __init__(self, target: str):
        super().__init__(target)
        self.stub = AccountTypeServiceStub(channel=self.channel)

    @staticmethod
    def to_domain(model: AccountTypeModel) -> AccountType:
        return AccountType(
            type_id=model.type_id,
            name=model.name,
            description=model.description
        )

    @staticmethod
    def to_model(domain: AccountType) -> AccountTypeModel:
        return AccountTypeModel(
            type_id=domain.type_id,
            name=domain.name,
            description=domain.description
        )
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[AccountType]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        return [self.to_domain(model) for model in result.account_types]


    async def get_by_id(self, type_id: int) -> AccountType | None:
        request = GetAccountTypeByIdRequest(type_id=type_id)
        result: AccountTypeModel = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None

