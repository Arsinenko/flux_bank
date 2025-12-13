from datetime import datetime
from decimal import Decimal
from typing import List

import grpc

from api.generated.custom_types_pb2 import GetAllRequest, GetByDateRangeRequest
from api.generated.transaction_pb2 import *
from api.generated.transaction_pb2_grpc import TransactionServiceStub
from domain.transaction.transaction import Transaction
from domain.transaction.transaction_repo import TransactionRepositoryAbc


class TransactionRepository(TransactionRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = TransactionServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model: TransactionModel) -> Transaction:
        return Transaction(
            transaction_id=model.transaction_id,
            source_account=model.source_account if model.HasField("source_account") else None,
            target_account=model.target_account if model.HasField("target_account") else None,
            amount=Decimal(model.amount),
            currency=model.currency,
            created_at=model.created_at.ToDatetime() if model.HasField("created_at") else None,
            status=model.status if model.HasField("status") else None
        )

    @staticmethod
    def response_to_list(response: GetAllTransactionsResponse) -> List[Transaction]:
        return [TransactionRepository.to_domain(model) for model in response.transactions]

    async def get_all(self, page_n: int, page_size: int) -> List[Transaction]:
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, transaction_id: int) -> Transaction | None:
        try:
            result = await self.stub.GetById(GetTransactionByIdRequest(transaction_id=transaction_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None

    async def get_by_date_range(self, start_date: datetime, end_date: datetime) -> List[Transaction]:
        try:
            request = GetByDateRangeRequest(startDate=start_date, endDate=end_date)
            result = await self.stub.GetByDateRange(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByDateRange: {err}")
            return []
