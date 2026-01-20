

from api.generated.account_pb2 import TotalBalanceResponse
from api.generated.account_analytic_pb2_grpc import AccountAnalyticServiceServicer
from domain.account.account_repo import AccountRepositoryAbc


class AccountAnalyticService(AccountAnalyticServiceServicer):
    def __init__(self, account_repo: AccountRepositoryAbc):
        self.account_repo = account_repo

    async def ProcessGetAvgBalance(self, request, context):
        balance = await self.account_repo.get_avg_balance(request.account_type_id, request.is_active)
        return TotalBalanceResponse(total_balance=str(balance))

    async def ProcessGetTotalBalanceByAccountType(self, request, context):
        balance = await self.account_repo.get_total_balance_by_type(request.account_type_id)
        return TotalBalanceResponse(total_balance=str(balance))
