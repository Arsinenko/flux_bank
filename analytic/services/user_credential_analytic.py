from api.generated.user_credential_analytic_pb2_grpc import UserCredentialAnalyticServiceServicer
from api.generated.user_credential_pb2 import GetUserCredentialByIdRequest, GetUserCredentialByIdsRequest, GetUserCredentialByUsernameRequest, GetAllUserCredentialsResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.user_credential.user_credential_repo import UserCredentialRepositoryAbc
from mappers.user_credential_mapper import UserCredentialMapper
from google.protobuf.empty_pb2 import Empty


class UserCredentialAnalyticService(UserCredentialAnalyticServiceServicer):
    def __init__(self, user_credential_repo: UserCredentialRepositoryAbc):
        self.user_credential_repo = user_credential_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.user_credential_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllUserCredentialsResponse(user_credentials=UserCredentialMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetUserCredentialByIdRequest, context):
        result = await self.user_credential_repo.get_by_id(request.customer_id)
        return UserCredentialMapper.to_model(result) if result else None

    async def ProcessGetByIds(self, request: GetUserCredentialByIdsRequest, context):
        result = await self.user_credential_repo.get_by_ids(request.customer_ids)
        return GetAllUserCredentialsResponse(user_credentials=UserCredentialMapper.to_model_list(result))

    async def ProcessGetByUsername(self, request: GetUserCredentialByUsernameRequest, context):
        result = await self.user_credential_repo.get_by_username(request.username)
        return UserCredentialMapper.to_model(result) if result else None

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.user_credential_repo.get_count()
        return CountResponse(count=result)
