from api.generated.account_analytic_pb2 import GetTopAccountByBalanceRequest
from api.generated.account_pb2 import TotalBalanceResponse, GetAllAccountsResponse, GetAccountByCustomerIdRequest, \
    GetAccountByIdsRequest
from api.generated.account_analytic_pb2_grpc import AccountAnalyticServiceServicer
from api.generated.custom_types_pb2 import GetAllRequest, GetByDateRangeRequest, CountResponse
from domain.account.account_repo import AccountRepositoryAbc
from mappers.account_mapper import AccountMapper


class AccountAnalyticService(AccountAnalyticServiceServicer):
    def __init__(self, account_repo: AccountRepositoryAbc):
        self.account_repo = account_repo

    async def ProcessGetAvgBalance(self, request, context):
        balance = await self.account_repo.get_avg_balance(request.account_type_id, request.is_active)
        return TotalBalanceResponse(total_balance=str(balance))

    async def ProcessGetTotalBalanceByAccountType(self, request, context):
        balance = await self.account_repo.get_total_balance_by_type(request.account_type_id)
        return TotalBalanceResponse(total_balance=str(balance))

    async def GetTopAccountByBalance(self, request: GetTopAccountByBalanceRequest, context):
        accounts = await self.account_repo.get_all(page_size=request.top_n, page_n=1, order_by="balance", is_desc=True)
        return GetAllAccountsResponse(accounts=AccountMapper.to_model_list(accounts))

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.account_repo.get_all(page_size=request.pageSize, page_n=request.pageN, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllAccountsResponse(accounts=AccountMapper.to_model_list(result))

    async def ProcessGetById(self, request, context):
        result = await self.account_repo.get_by_id(request.account_id)
        return AccountMapper.to_model(result) if result else None

    async def ProcessGetByCustomerId(self, request: GetAccountByCustomerIdRequest, context):
        result = await self.account_repo.get_by_customer_id(request.customer_id)
        return GetAllAccountsResponse(accounts=AccountMapper.to_model_list(result))

    async def ProcessGetByDateRange(self, request: GetByDateRangeRequest, context):
        result = await self.account_repo.get_by_date_range(request.fromDate, request.toDate)
        return GetAllAccountsResponse(accounts=AccountMapper.to_model_list(result))

    async def ProcessGetByIds(self, request: GetAccountByIdsRequest, context):
        result = await self.account_repo.get_by_ids(request.account_ids)
        return GetAllAccountsResponse(accounts=AccountMapper.to_model_list(result))

    async def ProcessGetCount(self, request, context):
        c = await self.account_repo.get_count()
        return CountResponse(count=c)

    async def ProcessGetCountByDateRange(self, request, context):
        c = await self.account_repo.get_count_by_date_range(request.fromDate, request.toDate)
        return CountResponse(count=c)

    async def ProcessGetCountByCustomerId(self, request, context):
        c = await self.account_repo.get_count_by_customer_id(request.customer_id)
        return CountResponse(count=c)

    async def ProcessGetCountByStatus(self, request, context):
        c = await self.account_repo.get_count_by_status(request.status)
        return CountResponse(count=c)