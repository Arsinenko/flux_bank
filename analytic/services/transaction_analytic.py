from google.protobuf.empty_pb2 import Empty

from api.generated.custom_types_pb2 import GetAllRequest, CountResponse, GetByDateRangeRequest
from api.generated.transaction_analitic_pb2_grpc import TransactionAnalyticServiceServicer
from api.generated.transaction_pb2 import GetTransactionByIdRequest, GetTransactionByIdsRequest, \
    GetAllTransactionsResponse, GetAccountExpensesRequest, GetAccountRevenueRequest, TotalAmountResponse
from domain.transaction.transaction_repo import TransactionRepositoryAbc
from mappers.transaction_mapper import TransactionMapper


class TransactionAnalyticService(TransactionAnalyticServiceServicer):
    def __init__(self, transaction_repo: TransactionRepositoryAbc):
        self.transaction_repo = transaction_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.transaction_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllTransactionsResponse(transactions=TransactionMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetTransactionByIdRequest, context):
        result = await self.transaction_repo.get_by_id(request.transaction_id)
        return TransactionMapper.to_model(result) if result else None

    async def ProcessGetByIds(self, request: GetTransactionByIdsRequest, context):
        result = await self.transaction_repo.get_by_ids(request.transaction_ids)
        return GetAllTransactionsResponse(transactions=TransactionMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.transaction_repo.get_count()
        return CountResponse(count=result)

    async def ProcessGetCountByDateRange(self, request: GetByDateRangeRequest, context):
        result = await self.transaction_repo.get_count_by_date_range(request.fromDate, request.toDate)
        return CountResponse(count=result)

    async def ProcessGetCountAccountExpenses(self, request: GetAccountExpensesRequest, context):
        result = await self.transaction_repo.get_count_expenses(account=request.source_account,
                                                                start_date=request.date_range.fromDate,
                                                                end_date=request.date_range.toDate)
        return CountResponse(count=result)

    async def ProcessGetAccountExpenses(self, request: GetAccountRevenueRequest, context):
        result = await self.transaction_repo.get_account_expenses(account_id=request.target_account)
        return GetAllTransactionsResponse(transactions=TransactionMapper.to_model_list(result))

    async def ProcessGetCountAccountRevenue(self, request: GetAccountRevenueRequest, context):
        result = await self.transaction_repo.get_count_revenue(account=request.target_account,
                                                               start_date=request.date_range.fromDate,
                                                               end_date=request.date_range.toDate)
        return CountResponse(count=result)

    async def ProcessGetTotalAmount(self, request, context):
        result = await self.transaction_repo.get_total_amount()
        return TotalAmountResponse(total_amount=str(result))

    async def ProcessGetByDateRange(self, request: GetByDateRangeRequest, context):
        result = await self.transaction_repo.get_by_date_range(start_date=request.fromDate, end_date=request.toDate)
        return GetAllTransactionsResponse(transactions=TransactionMapper.to_model_list(result))

    async def ProcessGetAccountRevenue(self, request: GetAccountRevenueRequest, context):
        result = await self.transaction_repo.get_revenue(account_id=request.target_account)
        return GetAllTransactionsResponse(transactions=TransactionMapper.to_model_list(result))
