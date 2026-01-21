from api.generated.atm_analytic_pb2_grpc import AtmAnalyticServiceServicer
from api.generated.atm_pb2 import GetAtmByIdRequest, GetAtmsByLocationSubStrRequest
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.atm.atm_repo import AtmRepositoryAbc
from mappers.atm_mapper import AtmMapper


class AtmAnalyticService(AtmAnalyticServiceServicer):
    def __init__(self, atm_repo: AtmRepositoryAbc):
        self.atm_repo = atm_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.atm_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return AtmMapper.to_model_list(result)

    async def ProcessGetCount(self, request, context):
        result = await self.atm_repo.get_count()
        return CountResponse(count=result)
    async def ProcessGetById(self, request: GetAtmByIdRequest, context):
        result = await self.atm_repo.get_by_id(request.atm_id)
        return AtmMapper.to_model(result) if result else None

    async def ProcessGetCountByStatus(self, request, context):
        result = await self.atm_repo.get_count_by_status(request.status)
        return CountResponse(count=result)

    async def ProcessGetByBranch(self, request, context):
        result = await self.atm_repo.get_by_branch(request.branch_id)
        return AtmMapper.to_model_list(result)

    async def ProcessGetByStatus(self, request, context):
        result = await self.atm_repo.get_by_status(request.status)
        return AtmMapper.to_model_list(result)

    async def ProcessGetByLocationSubStr(self, request: GetAtmsByLocationSubStrRequest, context):
        result = await self.atm_repo.get_by_location_substr(request.sub_str)
        return AtmMapper.to_model_list(result)