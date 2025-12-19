from decimal import Decimal
from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.deposit_pb2 import *
from api.generated.deposit_pb2_grpc import DepositServiceStub
from domain.deposit.deposit import Deposit
from domain.deposit.deposit_repo import DepositRepositoryAbc


class DepositRepository(DepositRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = DepositServiceStub(channel=self.chanel)

    @staticmethod
    def to_domain(model: DepositModel) -> Deposit:
        return Deposit(
            deposit_id=model.deposit_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            amount=Decimal(model.amount) if model.HasField("amount") else None,
            interest_rate=Decimal(model.interest_rate) if model.HasField("interest_rate") else None,
            start_date=model.start_date.ToDatetime() if model.HasField("start_date") else None,
            end_date=model.end_date.ToDatetime() if model.HasField("end_date") else None,
            status=model.status if model.HasField("status") else None
        )

    @staticmethod
    def response_to_list(response: GetAllDepositsResponse) -> List[Deposit]:
        return [DepositRepository.to_domain(model) for model in response.deposits]

    async def get_all(self, page_n: int, page_size: int) -> List[Deposit]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, deposit_id: int) -> Deposit | None:
        request = GetDepositByIdRequest(deposit_id=deposit_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None

    async def get_by_customer_id(self, customer_id: int) -> List[Deposit]:
        request = GetDepositsByCustomerRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetByCustomer(request))
        if result:
            return self.response_to_list(result)
        return []
