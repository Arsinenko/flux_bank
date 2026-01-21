from api.generated.login_log_analytic_pb2_grpc import LoginLogAnalyticServiceServicer
from api.generated.login_log_pb2 import GetLoginLogByIdRequest, GetLoginLogByIdsRequest, GetAllLoginLogsResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.login_log.login_log_repo import LoginLogRepositoryAbc
from mappers.login_log_mapper import LoginLogMapper
from google.protobuf.empty_pb2 import Empty


class LoginLogAnalyticService(LoginLogAnalyticServiceServicer):
    def __init__(self, login_log_repo: LoginLogRepositoryAbc):
        self.login_log_repo = login_log_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.login_log_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllLoginLogsResponse(login_logs=LoginLogMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetLoginLogByIdRequest, context):
        result = await self.login_log_repo.get_by_id(request.log_id)
        return LoginLogMapper.to_model(result) if result else None

    # async def ProcessGetByIds(self, request: GetLoginLogByIdsRequest, context):
    #     result = await self.login_log_repo.get_by_ids(request.log_ids)
    #     return GetAllLoginLogsResponse(login_logs=LoginLogMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.login_log_repo.get_count()
        return CountResponse(count=result)
