from datetime import datetime
from typing import List

import grpc
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.customer_pb2 import *
from api.generated.customer_pb2_grpc import CustomerServiceStub
from api.generated.custom_types_pb2 import GetAllRequest, GetByDateRangeRequest
from domain.customer.customer import Customer
from domain.customer.customer_repo import CustomerRepositoryAbc


from mappers.customer_mapper import CustomerMapper


class CustomerRepository(CustomerRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = CustomerServiceStub(channel=self.chanel)

    async def get_by_ids(self, ids: List[int]) -> List[Customer]:
        request = GetCustomerByIdsRequest(customer_ids=ids)
        result = await self._execute(self.stub.GetByIds(request))
        if result:
            return CustomerMapper.to_domain_list(result.customers)
        return []

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_by_substring(self, substring: str) -> List[Customer]:
        request = GetBySubstringRequest(sub_str=substring)
        result = await self._execute(self.stub.GetBySubstring(request))
        if result:
            return CustomerMapper.to_domain_list(result.customers)
        return []

    async def get_by_date_range(self, from_date, to_date, page_n: int, page_size: int) -> List[Customer]:
        request = GetByDateRangeRequest(fromDate=from_date, toDate=to_date, pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetByDateRange(request))
        if result:
            return CustomerMapper.to_domain_list(result.customers)
        return []

    async def get_count_by_substring(self, substring: str) -> int:
        result = await self._execute(self.stub.GetCountBySubstring(GetBySubstringRequest(sub_str=substring)))
        return result.count

    async def get_count_by_date_range(self, from_date: datetime, to_date: datetime) -> int:
        request = GetByDateRangeRequest(fromDate=from_date, toDate=to_date)
        result = await self._execute(self.stub.GetCountByDateRange(request))
        return result.count

    async def get_all(self, page_n: int, page_size: int) -> List[Customer]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return CustomerMapper.to_domain_list(result.customers)
        return []

    async def get_by_id(self, customer_id: int) -> Customer | None:
        request = GetCustomerByIdRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return CustomerMapper.to_domain(result)
        return None
