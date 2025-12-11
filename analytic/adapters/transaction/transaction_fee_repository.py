from decimal import Decimal
from typing import List

import grpc

from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.transaction_fee_pb2 import *
from api.generated.transaction_fee_pb2_grpc import TransactionFeeServiceStub
from domain.transaction.transaction_fee import TransactionFee
from domain.transaction.transaction_fee_repo import TransactionFeeRepositoryAbc


class TransactionFeeRepository(TransactionFeeRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = TransactionFeeServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model: TransactionFeeModel) -> TransactionFee:
        return TransactionFee(
            id=model.id,
            transaction_id=model.transaction_id if model.HasField("transaction_id") else None,
            fee_id=model.fee_id if model.HasField("fee_id") else None,
            amount=Decimal(model.amount) if model.HasField("amount") else None
        )

    @staticmethod
    def response_to_list(response: GetAllTransactionFeesResponse) -> List[TransactionFee]:
        return [TransactionFeeRepository.to_domain(model) for model in response.transaction_fees]

    async def get_all(self, page_n: int, page_size: int) -> List[TransactionFee]:
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, id: int) -> TransactionFee | None:
        try:
            result = await self.stub.GetById(GetTransactionFeeByIdRequest(id=id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None
