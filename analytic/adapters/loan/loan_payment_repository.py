from decimal import Decimal
from typing import List

import grpc
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.loan_payment_pb2 import *
from api.generated.loan_payment_pb2_grpc import LoanPaymentServiceStub
from domain.loan.loan_payment import LoanPayment
from domain.loan.loan_payment_repo import LoanPaymentRepositoryAbc


class LoanPaymentRepository(LoanPaymentRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = LoanPaymentServiceStub(channel=self.channel)

    @staticmethod
    def to_model(domain: LoanPayment) -> LoanPaymentModel:
        model = LoanPaymentModel(
            payment_id=domain.payment_id,
            loan_id=domain.loan_id,
            amount=str(domain.amount) if domain.amount is not None else None,
            is_paid=domain.is_paid
        )
        if domain.payment_date:
            model.payment_date.FromDatetime(domain.payment_date)
        return model

    @staticmethod
    def to_domain(model: LoanPaymentModel) -> LoanPayment:
        return LoanPayment(
            payment_id=model.payment_id,
            loan_id=model.loan_id if model.HasField("loan_id") else None,
            amount=Decimal(model.amount) if model.HasField("amount") else None,
            payment_date=model.payment_date.ToDatetime() if model.HasField("payment_date") else None,
            is_paid=model.is_paid if model.HasField("is_paid") else None
        )

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count


    @staticmethod
    def response_to_list(response: GetAllLoanPaymentsResponse) -> List[LoanPayment]:
        return [LoanPaymentRepository.to_domain(model) for model in response.loan_payments]

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[LoanPayment]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, payment_id: int) -> LoanPayment | None:
        request = GetLoanPaymentByIdRequest(payment_id=payment_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None

    async def get_by_loan_id(self, loan_id: int) -> List[LoanPayment]:
        request = GetLoanPaymentsByLoanRequest(loan_id=loan_id)
        result = await self._execute(self.stub.GetByLoan(request))
        if result:
            return self.response_to_list(result)
        return []
