from decimal import Decimal
from typing import List

import grpc
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.loan_pb2 import *
from api.generated.loan_pb2_grpc import LoanServiceStub
from domain.loan.loan import Loan
from domain.loan.loan_repo import LoanRepositoryAbc


from mappers.loan_mapper import LoanMapper


class LoanRepository(LoanRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = LoanServiceStub(channel=self.chanel)

    async def get_by_ids(self, ids: List[int]) -> List[Loan]:
        request = GetLoanByIdsRequest(loan_ids=ids)
        result = await self._execute(self.stub.GetByIds(request))
        if result:
            return LoanMapper.to_domain_list(result.loans)
        return []


    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_count_by_status(self, status: str) -> int:
        result = await self._execute(self.stub.GetCountByStatus(GetLoanCountByStatusRequest(status=status)))
        return result.count

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Loan]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return LoanMapper.to_domain_list(result.loans)
        return []

    async def get_by_id(self, loan_id: int) -> Loan | None:
        request = GetLoanByIdRequest(loan_id=loan_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return LoanMapper.to_domain(result)
        return None

    async def get_by_customer_id(self, customer_id: int) -> List[Loan]:
        request = GetLoansByCustomerRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetByCustomer(request))
        if result:
            return LoanMapper.to_domain_list(result.loans)
        return []
