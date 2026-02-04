from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from api.generated.custom_types_pb2 import GetAllRequest
from api.generated.customer_address_pb2 import *
from api.generated.customer_address_pb2_grpc import CustomerAddressServiceStub
from domain.customer.customer_address import CustomerAddress
from domain.customer.customer_address_repo import CustomerAddressRepositoryAbc


class CustomerAddressRepository(CustomerAddressRepositoryAbc, BaseGrpcRepository):
    def __init__(self, channel):
        super().__init__(channel)
        self.stub = CustomerAddressServiceStub(channel=self.channel)

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
    def to_model(domain: CustomerAddress) -> CustomerAddressModel:
        return CustomerAddressModel(
            address_id=domain.address_id,
            customer_id=domain.customer_id,
            country=domain.country,
            city=domain.city,
            street=domain.street,
            zip_code=domain.zip_code,
            is_primary=domain.is_primary
        )

    @staticmethod
    def response_to_list(response: GetAllCustomerAddressesResponse) -> List[CustomerAddress]:
        return [CustomerAddressRepository.to_domain(model) for model in response.customer_addresses]

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[CustomerAddress]:
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

    async def get_by_id(self, address_id: int) -> CustomerAddress | None:
        request = GetCustomerAddressByIdRequest(address_id=address_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None
