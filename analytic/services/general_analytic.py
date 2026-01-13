import asyncio

from adapters.account.account_repository import AccountRepository
from adapters.card.card_repository import CardRepository
from adapters.customer.customer_repository import CustomerRepository
from adapters.deposit.deposit_repository import DepositRepository
from adapters.loan.loan_repository import LoanRepository
from adapters.transaction.transaction_repository import TransactionRepository
from api.generated.general_analytic_pb2 import GeneralCountsResponse
from api.generated.general_analytic_pb2_grpc import GeneralAnalyticServiceServicer
from domain.account.account_repo import AccountRepositoryAbc
from domain.card.card_repo import CardRepositoryAbc
from domain.customer.customer_repo import CustomerRepositoryAbc
from domain.deposit.deposit_repo import DepositRepositoryAbc
from domain.loan.loan_repo import LoanRepositoryAbc
from domain.transaction.transaction_repo import TransactionRepositoryAbc


class GeneralAnalyticService(GeneralAnalyticServiceServicer):
    def __init__(self, target: str):
        self.account_repo: AccountRepositoryAbc = AccountRepository(target)
        self.transaction_repo: TransactionRepositoryAbc = TransactionRepository(target)
        self.card_repo: CardRepositoryAbc = CardRepository(target)
        self.loan_repo: LoanRepositoryAbc = LoanRepository(target)
        self.deposit_repo: DepositRepositoryAbc = DepositRepository(target)
        self.customer_repo: CustomerRepositoryAbc = CustomerRepository(target)

    async def GetGeneralCounts(self, request, context):
        # Вместо распаковки сразу, можно сделать так:
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


