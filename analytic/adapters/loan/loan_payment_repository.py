from decimal import Decimal
from typing import List

import grpc

from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.loan_payment_pb2 import *
from api.generated.loan_payment_pb2_grpc import LoanPaymentServiceStub
from domain.loan.loan_payment import LoanPayment
from domain.loan.loan_payment_repo import LoanPaymentRepositoryAbc


class LoanPaymentRepository(LoanPaymentRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = LoanPaymentServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model: LoanPaymentModel) -> LoanPayment:
        return LoanPayment(
            payment_id=model.payment_id,
            loan_id=model.loan_id if model.HasField("loan_id") else None,
            amount=Decimal(model.amount) if model.HasField("amount") else None,
            payment_date=model.payment_date.ToDatetime() if model.HasField("payment_date") else None,
            is_paid=model.is_paid if model.HasField("is_paid") else None
        )

    @staticmethod
    def response_to_list(response: GetAllLoanPaymentsResponse) -> List[LoanPayment]:
        return [LoanPaymentRepository.to_domain(model) for model in response.loan_payments]

    async def get_all(self, page_n: int, page_size: int) -> List[LoanPayment]:
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, payment_id: int) -> LoanPayment | None:
        try:
            result = await self.stub.GetById(GetLoanPaymentByIdRequest(payment_id=payment_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None

    async def get_by_loan_id(self, loan_id: int) -> List[LoanPayment]:
        try:
            result = await self.stub.GetByLoan(GetLoanPaymentsByLoanRequest(loan_id=loan_id))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByLoan: {err}")
            return []
