from typing import List

import grpc

from api.generated.customer_pb2 import *
from api.generated.customer_pb2_grpc import CustomerServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.customer.customer import Customer
from domain.customer.customer_repo import CustomerRepositoryAbc


class CustomerRepository(CustomerRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = CustomerServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

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

    async def get_all(self, page_n: int, page_size: int) -> List[Customer]:
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, customer_id: int) -> Customer | None:
        try:
            result = await self.stub.GetById(GetCustomerByIdRequest(customer_id=customer_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None
