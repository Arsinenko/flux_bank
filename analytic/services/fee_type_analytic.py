from api.generated.fee_type_analytic_pb2_grpc import FeeTypeAnalyticServiceServicer
from api.generated.fee_type_pb2 import GetFeeTypeByIdRequest, GetFeeTypeByIdsRequest, GetAllFeeTypesResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.fee_type.fee_type_repo import FeeTypeRepositoryAbc
from mappers.fee_type_mapper import FeeTypeMapper
from google.protobuf.empty_pb2 import Empty


class FeeTypeAnalyticService(FeeTypeAnalyticServiceServicer):
    def __init__(self, fee_type_repo: FeeTypeRepositoryAbc):
        self.fee_type_repo = fee_type_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.fee_type_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllFeeTypesResponse(fee_types=FeeTypeMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetFeeTypeByIdRequest, context):
        result = await self.fee_type_repo.get_by_id(request.fee_type_id)
        return FeeTypeMapper.to_model(result) if result else None

    # async def ProcessGetByIds(self, request: GetFeeTypeByIdsRequest, context):
    #     result = await self.fee_type_repo.get_by_ids(request.fee_type_ids)
    #     return GetAllFeeTypesResponse(fee_types=FeeTypeMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.fee_type_repo.get_count()
        return CountResponse(count=result)
