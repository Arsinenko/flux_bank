from typing import List

from api.generated.transaction_analitic_pb2_grpc import TransactionAnalyticServiceServicer
from domain.transaction.transaction import Transaction
from domain.transaction.transaction_repo import TransactionRepositoryAbc
from api.generated.transaction_analitic_pb2 import *
import pandas as pd
class TransactionAnalyticService(TransactionAnalyticServiceServicer):
    def __init__(self, transaction_repository: TransactionRepositoryAbc):
        self.transaction_repository = transaction_repository

    async def GetSumOfTransactionsByDateRange(self, request: GetSumOfTransactionsByDateRangeRequest, context):
        transactions = await self.transaction_repository.get_by_date_range(request.start_date, request.end_date)
        df = pd.DataFrame(t.__dict__ for t in transactions)
        return GetSumOfTransactionsByDateRangeResponse(sum=str(df['amount'].sum()))

    async def GetAverageOfTransactionsByDateRange(self, request: GetAverageOfTransactionsByDateRangeRequest, context):
        transactions: List[Transaction] = await self.transaction_repository.get_by_date_range(request.start_date, request.end_date)
        df = pd.DataFrame(t.__dict__ for t in transactions)
        return GetAverageOfTransactionsByDateRangeResponse(average=str(df['amount'].mean()))

    async def GetCountOfTransactionsByDateRange(self, request: GetCountOfTransactionsByDateRangeRequest, context):
        transactions = await self.transaction_repository.get_by_date_range(request.start_date, request.end_date)
        return GetCountOfTransactionsByDateRangeResponse(count=len(transactions))

    async def GetMostFrequentTransactionsByDateRange(self, request, context):
        pass



