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


class CustomerRepository(CustomerRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = CustomerServiceStub(channel=self.chanel)

    @staticmethod
    def to_domain(model: CustomerModel) -> Customer:
        return Customer(
            customer_id=model.customer_id,
            first_name=model.first_name,
            last_name=model.last_name,
            email=model.email,
            phone=model.phone if model.HasField("phone") else None,
            birth_date=model.birth_date.ToDatetime() if model.HasField("birth_date") else None,
            created_at=model.created_at.ToDatetime() if model.HasField("created_at") else None
        )

    @staticmethod
    def response_to_list(response: GetAllCustomersResponse) -> List[Customer]:
        return [CustomerRepository.to_domain(model) for model in response.customers]

    async def get_by_ids(self, ids: List[int]) -> List[Customer]:
        request = GetCustomerByIdsRequest(customer_ids=ids)
        result = await self._execute(self.stub.GetByIds(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_by_substring(self, substring: str) -> List[Customer]:
        request = GetBySubstringRequest(sub_str=substring)
        result = await self._execute(self.stub.GetBySubstring(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_date_range(self, from_date, to_date, page_n: int, page_size: int) -> List[Customer]:
        request = GetByDateRangeRequest(fromDate=from_date, toDate=to_date, pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetByDateRange(request))
        if result:
            return self.response_to_list(result)
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
            return self.response_to_list(result)
        return []

    async def get_by_id(self, customer_id: int) -> Customer | None:
        request = GetCustomerByIdRequest(customer_id=customer_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None
