import asyncio
from datetime import datetime
from typing import List

import pandas as pd
from dateutil.relativedelta import relativedelta
from google.protobuf.timestamp_pb2 import Timestamp

from api.generated.custom_types_pb2 import GetAllRequest, CountResponse, GetByDateRangeRequest
from api.generated.customer_analytic_pb2 import (
    GetCustomerDetailsRequest,
    CustomerDetailsResponse,
    GetCustomersByTransactionQuantityRangeRequest,
    GetCustomerLifeTimeRequest,
    GetCustomerLifeTimeResponse,
    GetCustomersLifeTimeResponse,
)
from api.generated.customer_analytic_pb2_grpc import CustomerAnalyticServiceServicer
from api.generated.customer_pb2 import GetAllCustomersResponse, GetCustomerByIdRequest, GetCustomerByIdsRequest, \
    GetInactiveCustomersRequest, GetBySubstringRequest
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
from mappers.account_mapper import AccountMapper
from mappers.card_mapper import CardMapper
from mappers.customer_mapper import CustomerMapper
from mappers.deposit_mapper import DepositMapper
from mappers.loan_mapper import LoanMapper


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
        threshold = datetime.now() - relativedelta(months=2)
        
        inactive_customers = await self.customer_repo.get_inactive(
            threshold=threshold,
            page_n=request.pageN,
            page_size=request.pageSize
        )
        
        return GetAllCustomersResponse(
            customers=CustomerMapper.to_model_list(inactive_customers)
        )

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.customer_repo.get_all(
            page_n=request.pageN,
            page_size=request.pageSize,
            order_by=request.order_by.value,
            is_desc=request.is_desc.value
        )
        return GetAllCustomersResponse(customers=CustomerMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetCustomerByIdRequest, context):
        result = await self.customer_repo.get_by_id(request.customer_id)
        if result is None:
            return None
        return CustomerMapper.to_model(result)

    async def ProcessGetByIds(self, request: GetCustomerByIdsRequest, context):
        result = await self.customer_repo.get_by_ids(request.customer_ids)
        return GetAllCustomersResponse(customers=CustomerMapper.to_model_list(result))

    async def ProcessGetBySubstring(self, request: GetBySubstringRequest, context):
        result = await self.customer_repo.get_by_substring(request.sub_str)
        return GetAllCustomersResponse(customers=CustomerMapper.to_model_list(result))

    async def ProcessGetByDateRange(self, request: GetByDateRangeRequest, context):
        result = await self.customer_repo.get_by_date_range(
            from_date=request.fromDate,
            to_date=request.toDate,
            page_n=1,
            page_size=1000
        )
        return GetAllCustomersResponse(customers=CustomerMapper.to_model_list(result))

    async def ProcessGetCount(self, request, context):
        c = await self.customer_repo.get_count()
        return CountResponse(count=c)

    async def ProcessGetCountBySubstring(self, request: GetBySubstringRequest, context):
        c = await self.customer_repo.get_count_by_substring(request.sub_str)
        return CountResponse(count=c)

    async def ProcessGetCountByDateRange(self, request: GetByDateRangeRequest, context):
        c = await self.customer_repo.get_count_by_date_range(request.fromDate, request.toDate)
        return CountResponse(count=c)

    async def ProcessGetInactive(self, request: GetInactiveCustomersRequest, context):
        result = await self.customer_repo.get_inactive(
            threshold=request.threshold_time,
            page_n=request.page_n,
            page_size=request.page_size
        )
        return GetAllCustomersResponse(customers=CustomerMapper.to_model_list(result))

    async def GetCustomersByTransactionQuantityRange(self, request: GetCustomersByTransactionQuantityRangeRequest, context):
        # NOTE: Using pandas to aggregate transaction counts per customer
        accounts = await self.account_repo.get_all(page_n=1, page_size=10000)
        transactions = await self.transaction_repo.get_all(page_n=1, page_size=100000)
        
        df_acc = pd.DataFrame([vars(a) for a in accounts])
        df_trans = pd.DataFrame([vars(t) for t in transactions])
        
        # Merge transactions with accounts to get customer_id
        df_merged = pd.merge(df_trans, df_acc, left_on='source_account', right_on='account_id', how='left')
        
        # Group by customer_id and count
        counts = df_merged.groupby('customer_id').size().reset_index(name='qty')
        
        # Filter by range
        filtered_ids = counts[(counts['qty'] >= request.min_quantity) & (counts['qty'] <= request.max_quantity)]['customer_id'].tolist()
        
        customers = await self.customer_repo.get_by_ids(filtered_ids)
        return GetAllCustomersResponse(customers=CustomerMapper.to_model_list(customers))

    async def GetCustomerLifeTime(self, request: GetCustomerLifeTimeRequest, context):
        customer = await self.customer_repo.get_by_id(request.customer_id)
        if not customer:
            return GetCustomerLifeTimeResponse()
            
        ts = Timestamp()
        if customer.created_at:
            ts.FromDatetime(customer.created_at)
            
        return GetCustomerLifeTimeResponse(customer_id=customer.customer_id, life_time=ts)

    async def GetCustomersLifeTime(self, request, context):
        customers = await self.customer_repo.get_all(page_n=1, page_size=1000)
        responses = []
        for c in customers:
            ts = Timestamp()
            if c.created_at:
                ts.FromDatetime(c.created_at)
            responses.append(GetCustomerLifeTimeResponse(customer_id=c.customer_id, life_time=ts))
        return GetCustomersLifeTimeResponse(customer_life_time=responses)

    async def GetCountByTransactionQuantityRange(self, request: GetCustomersByTransactionQuantityRangeRequest, context):
        # NOTE: Using pandas to aggregate transaction counts per customer
        accounts = await self.account_repo.get_all(page_n=1, page_size=10000)
        transactions = await self.transaction_repo.get_all(page_n=1, page_size=100000)
        
        df_acc = pd.DataFrame([vars(a) for a in accounts])
        df_trans = pd.DataFrame([vars(t) for t in transactions])
        
        df_merged = pd.merge(df_trans, df_acc, left_on='source_account', right_on='account_id', how='left')
        counts = df_merged.groupby('customer_id').size().reset_index(name='qty')
        
        count = len(counts[(counts['qty'] >= request.min_quantity) & (counts['qty'] <= request.max_quantity)])
        return CountResponse(count=count)

    async def GetCountInactiveCustomers(self, request, context):
        threshold = datetime.now() - relativedelta(months=2)
        # Note: we need a get_count_inactive in repo or use get_inactive with large size
        # Assuming we can just get all and count for now or if repo has it
        inactive = await self.customer_repo.get_inactive(threshold=threshold, page_n=1, page_size=10000)
        return CountResponse(count=len(inactive))
