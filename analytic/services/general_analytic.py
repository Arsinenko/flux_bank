import asyncio

from api.generated.general_analytic_pb2 import GeneralCountsResponse
from api.generated.general_analytic_pb2_grpc import GeneralAnalyticServiceServicer
from domain.account.account_repo import AccountRepositoryAbc
from domain.card.card_repo import CardRepositoryAbc
from domain.customer.customer_repo import CustomerRepositoryAbc
from domain.deposit.deposit_repo import DepositRepositoryAbc
from domain.loan.loan_repo import LoanRepositoryAbc
from domain.transaction.transaction_repo import TransactionRepositoryAbc


class GeneralAnalyticService(GeneralAnalyticServiceServicer):
    def __init__(self, account_repo: AccountRepositoryAbc, transaction_repo: TransactionRepositoryAbc,
                 card_repo: CardRepositoryAbc, loan_repo: LoanRepositoryAbc, deposit_repo: DepositRepositoryAbc,
                 customer_repo: CustomerRepositoryAbc):
        self.account_repo = account_repo
        self.transaction_repo = transaction_repo
        self.card_repo = card_repo
        self.loan_repo = loan_repo
        self.deposit_repo = deposit_repo
        self.customer_repo = customer_repo

    async def GetGeneralCounts(self, request, context):
        tasks = [
            self.customer_repo.get_count(),
            self.account_repo.get_count_by_status(status=True),
            self.card_repo.get_count_by_status(status="active"),
            self.loan_repo.get_count_by_status(status="active"),
            self.deposit_repo.get_count_by_status(status="active"),
            self.transaction_repo.get_count(),
            self.transaction_repo.get_total_amount()
        ]
        results = await asyncio.gather(*tasks)

        (total_customers,
         total_active_accounts,
         total_active_cards,
         total_active_loans,
         total_active_deposits,
         total_transactions_count,
         total_transactions_amount) = results

        return GeneralCountsResponse(
            total_customers=total_customers,
            total_active_accounts=total_active_accounts,
            total_active_cards=total_active_cards,
            total_active_loans=total_active_loans,
            total_active_deposits=total_active_deposits,
            total_transactions_count=total_transactions_count,
            total_transactions_amount=str(total_transactions_amount)
        )


