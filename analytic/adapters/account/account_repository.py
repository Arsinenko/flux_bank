from datetime import datetime
from typing import List

import grpc.aio

from api.generated import account_pb2
from api.generated.account_pb2 import AccountModel
from api.generated.account_pb2_grpc import AccountServiceStub
from api.generated.custom_types_pb2 import GetAllRequest, GetByDateRangeRequest
from domain.account.account import Account
from domain.account.account_repo import AccountRepositoryAbc


class AccountRepository(AccountRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = AccountServiceStub(self.chanel)
    async def close(self):
        await self.chanel.close()

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
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []


    async def get_by_id(self, account_id: int) -> Account | None:
        try:
            request = account_pb2.GetAccountByIdRequest(account_id=account_id)
            result = await self.stub.GetById(request)
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None

    async def get_by_customer_id(self, customer_id: int) -> List[Account]:
        try:
            request = account_pb2.GetAccountByCustomerIdRequest(customer_id=customer_id)
            result: account_pb2.GetAllAccountsResponse = await self.stub.GetByCustomerId(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByCustomerId: {err}")
            return []

    async def get_by_date_range(self, from_date, to_date, page_n: int, page_size: int) -> List[Account]:
        try:
            request = GetByDateRangeRequest(fromDate=from_date, toDate=to_date)
            result = await self.stub.GetByDateRange(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as ex:
            print(f"Error calling GetByDateRange: {ex}")
            return []
