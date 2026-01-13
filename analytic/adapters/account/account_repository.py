import decimal
from datetime import datetime
from typing import List

import grpc.aio
from decorator import EMPTY
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated import account_pb2
from api.generated.account_pb2 import *
from api.generated.account_pb2_grpc import AccountServiceStub
from api.generated.custom_types_pb2 import GetAllRequest, GetByDateRangeRequest, CountResponse
from domain.account.account import Account
from domain.account.account_repo import AccountRepositoryAbc


class AccountRepository(AccountRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = AccountServiceStub(self.chanel)


    @staticmethod
    def to_domain(self, model) -> Account:
        return Account(
            account_id=model.account_id,
            customer_id=model.customer_id,
            type_id=model.type_id,
            is_active=model.status,
            created_at=model.created_at.ToDatetime(),
            balance=model.balance,
            iban=model.iban
        )
    @staticmethod
    def response_to_list(self, response: account_pb2.GetAllAccountsResponse) ->List[Account]:
        return [self.to_domain(model) for model in response.accounts]


    async def get_all(self, page_n: int, page_size: int) -> List[Account]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        return self.response_to_list(result)

    async def get_by_id(self, account_id: int) -> Account | None:
        request = account_pb2.GetAccountByIdRequest(account_id=account_id)
        result: AccountModel = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None


    async def get_by_customer_id(self, customer_id: int) -> List[Account]:
        request = account_pb2.GetAccountByCustomerIdRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetByCustomerId(request))
        return self.response_to_list(result)

    async def get_by_date_range(self, from_date, to_date, page_n: int, page_size: int) -> List[Account]:
        request = GetByDateRangeRequest(
            fromDate=from_date,
            toDate=to_date,
            pageN=page_n,
            pageSize=page_size
        )
        result = await self._execute(self.stub.GetByDateRange(request))
        return self.response_to_list(result)

    async def get_count_by_customer_id(self, customer_id: int) -> int:
        request = GetAccountByCustomerIdRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetCountByCustomerId(request))
        return result.count

    async def get_by_ids(self, ids: List[int]) -> List[Account]:
        request = GetAccountByIdsRequest(account_ids=ids)
        result = await self._execute(self.stub.GetByIds(request))
        return self.response_to_list(result)

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_total_balance(self) -> decimal.Decimal:
        result: account_pb2.TotalBalanceResponse = await self._execute(self.stub.GetTotalBalance(Empty()))
        return decimal.Decimal(result.total_balance)

    async def get_count_by_date_range(self, from_date: datetime, to_date: datetime) -> int:
        request = GetByDateRangeRequest(
            fromDate=from_date,
            toDate=to_date
        )
        result: CountResponse = await self._execute(self.stub.GetCountByDateRange(request))
        return result.count

    async def get_count_by_status(self, status: bool) -> int:
        request = GetAccountsByStatusRequest(status=status)
        result = await self._execute(self.stub.GetCountByStatus(request))
        return result.count