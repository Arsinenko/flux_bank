from datetime import datetime
from decimal import Decimal
from typing import List

import grpc
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.custom_types_pb2 import GetAllRequest, GetByDateRangeRequest
from api.generated.transaction_pb2 import *
from api.generated.transaction_pb2_grpc import TransactionServiceStub
from domain.transaction.transaction import Transaction
from domain.transaction.transaction_repo import TransactionRepositoryAbc


from mappers.transaction_mapper import TransactionMapper


class TransactionRepository(TransactionRepositoryAbc, BaseGrpcRepository):
    def __init__(self, channel):
        super().__init__(channel)
        self.stub = TransactionServiceStub(channel=self.channel)

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Transaction]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return TransactionMapper.to_domain_list(result.transactions)
        return []

    async def get_by_ids(self, ids: List[int]) -> List[Transaction]:
        request = GetTransactionByIdsRequest(transaction_ids=ids)
        result = await self._execute(self.stub.GetByIds(request))
        if result:
            return TransactionMapper.to_domain_list(result.transactions)
        return []


    async def get_revenue(self, account_id) -> List[Transaction]:
        request = GetAccountRevenueRequest(target_account=account_id)
        result = await self._execute(self.stub.GetAccountRevenue(request))
        if result:
            return TransactionMapper.to_domain_list(result.transactions)
        return []


    async def get_account_expenses(self, account_id) -> List[Transaction]:
        request = GetAccountExpensesRequest(source_account=account_id)
        result = await self._execute(self.stub.GetAccountExpenses(request))
        if result:
            return TransactionMapper.to_domain_list(result.transactions)
        return []


    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count


    async def get_count_by_date_range(self, start_date: datetime, end_date: datetime) -> int:
        request = GetByDateRangeRequest(fromDate=start_date, toDate=end_date)
        result = await self._execute(self.stub.GetCountByDateRange(request))
        return result.count


    async def get_count_revenue(self, account: int, start_date: datetime = None, end_date: datetime = None) -> int:
        request = GetAccountRevenueRequest(target_account=account)
        if start_date and end_date:
            request.date_range.CopyFrom(
                GetByDateRangeRequest(
                    fromDate=start_date,
                    toDate=end_date
                )
            )
        result = await self._execute(self.stub.GetCountAccountRevenue(request))
        return result.count


    async def get_count_expenses(self, account: int, start_date: datetime = None, end_date: datetime = None) -> int:
        request = GetAccountExpensesRequest(source_account=account)
        if start_date and end_date:
            request.date_range.CopyFrom(
                GetByDateRangeRequest(
                    fromDate=start_date,
                    toDate=end_date
                )
            )
        result = await self._execute(self.stub.GetCountAccountExpenses(request))
        return result.count

    async def get_total_amount(self) -> Decimal:
        result = await self._execute(self.stub.GetTotalAmount(Empty()))
        return Decimal(result.total_amount)

    async def get_by_id(self, transaction_id: int) -> Transaction | None:
        request = GetTransactionByIdRequest(transaction_id=transaction_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return TransactionMapper.to_domain(result)
        return None

    async def get_by_date_range(self, start_date: datetime, end_date: datetime) -> List[Transaction]:
        request = GetByDateRangeRequest(fromDate=start_date, toDate=end_date)
        result = await self._execute(self.stub.GetByDateRange(request))
        if result:
            return TransactionMapper.to_domain_list(result.transactions)
        return []
