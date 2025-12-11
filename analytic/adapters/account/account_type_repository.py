from typing import List

import grpc.aio

from api.generated.account_type_pb2 import AccountTypeModel
from api.generated.account_type_pb2_grpc import AccountTypeServiceStub
from domain.account.account_type import AccountType
from domain.account.account_type_repo import AccountTypeRepositoryAbc


class AccountTypeRepository(AccountTypeRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = AccountTypeServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model: AccountTypeModel) -> AccountType:
        return AccountType(
            type_id=model.type_id,
            name=model.name,
            description=model.description
        )
    async def get_all(self) -> List[AccountType]:
        try:
            result = await self.stub.GetAll(AccountTypeModel())
            return [self.to_domain(model) for model in result.account_types]
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, type_id: int) -> AccountType | None:
        try:
            result = await self.stub.GetById(AccountTypeModel(type_id=type_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None

