from api.generated.branch_analytic_pb2_grpc import BranchAnalyticServiceServicer
from api.generated.branch_pb2 import GetBranchByIdRequest, GetBranchByIdsRequest, GetAllBranchesResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.branch.branch_repo import BranchRepositoryAbc
from mappers.branch_mapper import BranchMapper
from google.protobuf.empty_pb2 import Empty


class BranchAnalyticService(BranchAnalyticServiceServicer):
    def __init__(self, branch_repo: BranchRepositoryAbc):
        self.branch_repo = branch_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.branch_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllBranchesResponse(branches=BranchMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetBranchByIdRequest, context):
        result = await self.branch_repo.get_by_id(request.branch_id)
        return BranchMapper.to_model(result) if result else None

    async def ProcessGetByIds(self, request: GetBranchByIdsRequest, context):
        result = await self.branch_repo.get_by_ids(request.branch_ids)
        return GetAllBranchesResponse(branches=BranchMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.branch_repo.get_count()
        return CountResponse(count=result)
