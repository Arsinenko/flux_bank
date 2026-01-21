from api.generated.deposit_analytic_pb2_grpc import DepositAnalyticServiceServicer
from api.generated.deposit_pb2 import GetDepositByIdRequest, GetDepositByIdsRequest, GetAllDepositsResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.deposit.deposit_repo import DepositRepositoryAbc
from mappers.deposit_mapper import DepositMapper
from google.protobuf.empty_pb2 import Empty


class DepositAnalyticService(DepositAnalyticServiceServicer):
    def __init__(self, deposit_repo: DepositRepositoryAbc):
        self.deposit_repo = deposit_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.deposit_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllDepositsResponse(deposits=DepositMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetDepositByIdRequest, context):
        result = await self.deposit_repo.get_by_id(request.deposit_id)
        return DepositMapper.to_model(result) if result else None

    async def ProcessGetByIds(self, request: GetDepositByIdsRequest, context):
        result = await self.deposit_repo.get_by_ids(request.deposit_ids)
        return GetAllDepositsResponse(deposits=DepositMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.deposit_repo.get_count()
        return CountResponse(count=result)

    async def ProcessGetCountByStatus(self, request, context):
        result = await self.deposit_repo.get_count_by_status(request.status)
        return CountResponse(count=result)
