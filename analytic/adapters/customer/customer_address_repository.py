from typing import List

import grpc

from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.customer_address_pb2 import *
from api.generated.customer_address_pb2_grpc import CustomerAddressServiceStub
from domain.customer.customer_address import CustomerAddress
from domain.customer.customer_address_repo import CustomerAddressRepositoryAbc


class CustomerAddressRepository(CustomerAddressRepositoryAbc):
    def __init__(self, target: str):
        self.chanel = grpc.aio.insecure_channel(target)
        self.stub = CustomerAddressServiceStub(channel=self.chanel)

    async def close(self):
        await self.chanel.close()

    @staticmethod
    def to_domain(model: CustomerAddressModel) -> CustomerAddress:
        return CustomerAddress(
            address_id=model.address_id,
            customer_id=model.customer_id if model.HasField("customer_id") else None,
            country=model.country if model.HasField("country") else None,
            city=model.city if model.HasField("city") else None,
            street=model.street if model.HasField("street") else None,
            zip_code=model.zip_code if model.HasField("zip_code") else None,
            is_primary=model.is_primary if model.HasField("is_primary") else None
        )

    @staticmethod
    def response_to_list(response: GetAllCustomerAddressesResponse) -> List[CustomerAddress]:
        return [CustomerAddressRepository.to_domain(model) for model in response.customer_addresses]

    async def get_all(self, page_n: int, page_size: int) -> List[CustomerAddress]:
        try:
            request = GetAllRequest(pageN=page_n, pageSize=page_size)
            result = await self.stub.GetAll(request)
            return self.response_to_list(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetAll: {err}")
            return []

    async def get_by_id(self, address_id: int) -> CustomerAddress | None:
        try:
            result = await self.stub.GetById(GetCustomerAddressByIdRequest(address_id=address_id))
            return self.to_domain(result)
        except grpc.aio.AioRpcError as err:
            print(f"Error calling GetById: {err}")
            return None
