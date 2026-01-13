from decimal import Decimal
from typing import List

import grpc
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.loan_pb2 import *
from api.generated.loan_pb2_grpc import LoanServiceStub
from domain.loan.loan import Loan
from domain.loan.loan_repo import LoanRepositoryAbc


class LoanRepository(LoanRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = LoanServiceStub(channel=self.chanel)

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

    async def get_by_ids(self, ids: List[int]) -> List[Loan]:
        request = GetLoanByIdsRequest(loan_ids=ids)
        result = await self._execute(self.stub.GetByIds(request))
        if result:
            return self.response_to_list(result)
        return []


    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_count_by_status(self, status: str) -> int:
        result = await self._execute(self.stub.GetCountByStatus(GetLoanCountByStatusRequest(status=status)))
        return result.count

    async def get_all(self, page_n: int, page_size: int) -> List[Loan]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, loan_id: int) -> Loan | None:
        request = GetLoanByIdRequest(loan_id=loan_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None

    async def get_by_customer_id(self, customer_id: int) -> List[Loan]:
        request = GetLoansByCustomerRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetByCustomer(request))
        if result:
            return self.response_to_list(result)
        return []
