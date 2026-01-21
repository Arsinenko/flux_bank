from api.generated.transaction_category_analytic_pb2_grpc import TransactionCategoryAnalyticServiceServicer
from api.generated.transaction_category_pb2 import GetTransactionCategoryByIdRequest, GetTransactionCategoryByIdsRequest, GetAllTransactionCategoriesResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.transaction.transaction_category_repo import TransactionCategoryRepositoryAbc
from mappers.transaction_category_mapper import TransactionCategoryMapper
from google.protobuf.empty_pb2 import Empty


class TransactionCategoryAnalyticService(TransactionCategoryAnalyticServiceServicer):
    def __init__(self, transaction_category_repo: TransactionCategoryRepositoryAbc):
        self.transaction_category_repo = transaction_category_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.transaction_category_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllTransactionCategoriesResponse(transaction_categories=TransactionCategoryMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetTransactionCategoryByIdRequest, context):
        result = await self.transaction_category_repo.get_by_id(request.category_id)
        return TransactionCategoryMapper.to_model(result) if result else None

    # async def ProcessGetByIds(self, request: GetTransactionCategoryByIdsRequest, context):
    #     result = await self.transaction_category_repo.get_by_ids(request.category_ids)
    #     return GetAllTransactionCategoriesResponse(transaction_categories=TransactionCategoryMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.transaction_category_repo.get_count()
        return CountResponse(count=result)
