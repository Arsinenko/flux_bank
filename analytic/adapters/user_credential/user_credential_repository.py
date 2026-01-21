from typing import List

from google.protobuf.empty_pb2 import Empty
from google.protobuf.wrappers_pb2 import StringValue, BoolValue

from api.generated.user_credential_pb2 import (
    GetUserCredentialByIdRequest,
    GetUserCredentialByIdsRequest,
    GetUserCredentialByUsernameRequest
)
from api.generated.user_credential_pb2_grpc import UserCredentialServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.user_credential.user_credential import UserCredential
from domain.user_credential.user_credential_repo import UserCredentialRepositoryAbc
from adapters.base_grpc_repository import BaseGrpcRepository
from mappers.user_credential_mapper import UserCredentialMapper


class UserCredentialRepository(UserCredentialRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = UserCredentialServiceStub(channel=self.chanel)

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[UserCredential]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return UserCredentialMapper.to_domain_list(result.user_credentials)
        return []

    async def get_by_id(self, customer_id: int) -> UserCredential | None:
        request = GetUserCredentialByIdRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return UserCredentialMapper.to_domain(result)
        return None

    async def get_by_ids(self, customer_ids: List[int]) -> List[UserCredential]:
        request = GetUserCredentialByIdsRequest(customer_ids=customer_ids)
        result = await self._execute(self.stub.GetByIds(request))
        if result:
            return UserCredentialMapper.to_domain_list(result.user_credentials)
        return []

    async def get_by_username(self, username: str) -> UserCredential | None:
        request = GetUserCredentialByUsernameRequest(username=username)
        result = await self._execute(self.stub.GetByUsername(request))
        if result:
            return UserCredentialMapper.to_domain(result)
        return None

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count
