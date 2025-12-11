from decimal import Decimal
from typing import List

import grpc

from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.loan_pb2 import *
from api.generated.loan_pb2_grpc import LoanServiceStub
from domain.loan.loan import Loan
from domain.loan.loan_repo import LoanRepositoryAbc


class LoanRepository(LoanRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = LoanServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model: LoanModel) -> Loan:
        return Loan(
            loan_id=model.loan_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            principal=Decimal(model.principal) if model.HasField("principal") else None,
            interest_rate=Decimal(model.interest_rate) if model.HasField("interest_rate") else None,
            start_date=model.start_date.ToDatetime() if model.HasField("start_date") else None,
            end_date=model.end_date.ToDatetime() if model.HasField("end_date") else None,
            status=model.status if model.HasField("status") else None
        )

    @staticmethod
    def response_to_list(response: GetAllLoansResponse) -> List[Loan]:
        return [LoanRepository.to_domain(model) for model in response.loans]

    async def get_all(self, page_n: int, page_size: int) -> List[Loan]:
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, loan_id: int) -> Loan | None:
        try:
            result = await self.stub.GetById(GetLoanByIdRequest(loan_id=loan_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None

    async def get_by_customer_id(self, customer_id: int) -> List[Loan]:
        try:
            result = await self.stub.GetByCustomer(GetLoansByCustomerRequest(customer_id=customer_id))
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetByCustomer: {err}")
            return []
