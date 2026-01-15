import asyncio
from typing import List

from api.generated.custom_types_pb2 import GetAllRequest
from mappers.account_mapper import AccountMapper
from mappers.card_mapper import CardMapper
from mappers.customer_mapper import CustomerMapper
from mappers.deposit_mapper import DepositMapper
from mappers.loan_mapper import LoanMapper
from api.generated.customer_analytic_pb2 import GetCustomerDetailsRequest, CustomerDetailsResponse, GetAllCustomersResponse
from api.generated.customer_analytic_pb2_grpc import CustomerAnalyticServiceServicer
from domain.account.account import Account
from domain.account.account_repo import AccountRepositoryAbc
from domain.card.card import Card
from domain.card.card_repo import CardRepositoryAbc
from domain.customer.customer import Customer
from domain.customer.customer_repo import CustomerRepositoryAbc
from domain.deposit.deposit import Deposit
from domain.deposit.deposit_repo import DepositRepositoryAbc
from domain.loan.loan import Loan
from domain.loan.loan_repo import LoanRepositoryAbc
from domain.login_log.login_log import LoginLog
from domain.login_log.login_log_repo import LoginLogRepositoryAbc
from domain.transaction.transaction_repo import TransactionRepositoryAbc


class CustomerAnalyticService(CustomerAnalyticServiceServicer):
    def __init__(self, account_repo: AccountRepositoryAbc,
                 customer_repo: CustomerRepositoryAbc,
                 card_repo: CardRepositoryAbc,
                 transaction_repo: TransactionRepositoryAbc,
                 loan_repo: LoanRepositoryAbc,
                 deposit_repo: DepositRepositoryAbc,
                 login_log_repo: LoginLogRepositoryAbc):
        self.login_log_repo = login_log_repo
        self.deposit_repo = deposit_repo
        self.loan_repo = loan_repo
        self.transaction_repo = transaction_repo
        self.card_repo = card_repo
        self.customer_repo = customer_repo
        self.account_repo = account_repo

    async def GetCustomerDetails(self, request: GetCustomerDetailsRequest, context):
        tasks = [
            self.customer_repo.get_by_id(request.customer_id),
            self.account_repo.get_by_customer_id(request.customer_id),
            self.loan_repo.get_by_customer_id(request.customer_id),
            self.deposit_repo.get_by_customer_id(request.customer_id),
            self.login_log_repo.get_by_customer(request.customer_id)
        ]

        result = await asyncio.gather(*tasks)

        customer: Customer
        accounts: List[Account]
        cards: List[Card] = []
        loans: List[Loan]
        deposits: List[Deposit]
        login_logs: List[LoginLog]

        (customer, accounts, loans, deposits, login_logs) = result

        accounts_ids = [a.account_id for a in accounts]
        cards_tasks = []
        for account_id in accounts_ids:
            cards_tasks.append(self.card_repo.get_by_account_id(account_id))

        card_results = await asyncio.gather(*cards_tasks)
        for card_result in card_results:
            for card in card_result:
                cards.append(card)


        return CustomerDetailsResponse(
            customer=CustomerMapper.to_model(customer),
            accounts=AccountMapper.to_model_list(accounts),
            cards=CardMapper.to_model_list(cards),
            loans=LoanMapper.to_model_list(loans),
            deposits=DepositMapper.to_model_list(deposits),
        )

    async def GetInactiveCustomers(self, request: GetAllRequest, context):
        # By default, we search for customers inactive for more than 2 months
        from datetime import datetime
        from dateutil.relativedelta import relativedelta
        
        threshold = datetime.now() - relativedelta(months=2)
        
        inactive_customers = await self.customer_repo.get_inactive(
            threshold=threshold,
            page_n=request.pageN,
            page_size=request.pageSize
        )
        
        return GetAllCustomersResponse(
            customers=CustomerMapper.to_model_list(inactive_customers)
        )
