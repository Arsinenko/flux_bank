from decimal import Decimal
from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.transaction_fee_pb2 import *
from api.generated.transaction_fee_pb2_grpc import TransactionFeeServiceStub
from domain.transaction.transaction_fee import TransactionFee
from domain.transaction.transaction_fee_repo import TransactionFeeRepositoryAbc


class TransactionFeeRepository(TransactionFeeRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = TransactionFeeServiceStub(channel=self.chanel)

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
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, id: int) -> TransactionFee | None:
        request = GetTransactionFeeByIdRequest(id=id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None
