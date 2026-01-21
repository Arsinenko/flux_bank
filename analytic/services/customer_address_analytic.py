from api.generated.customer_address_analytic_pb2_grpc import CustomerAddressAnalyticServiceServicer
from api.generated.customer_address_pb2 import GetCustomerAddressByIdRequest, GetCustomerAddressByIdsRequest, GetAllCustomerAddressesResponse
from api.generated.custom_types_pb2 import GetAllRequest, CountResponse
from domain.customer.customer_address_repo import CustomerAddressRepositoryAbc
from mappers.customer_address_mapper import CustomerAddressMapper
from google.protobuf.empty_pb2 import Empty


class CustomerAddressAnalyticService(CustomerAddressAnalyticServiceServicer):
    def __init__(self, customer_address_repo: CustomerAddressRepositoryAbc):
        self.customer_address_repo = customer_address_repo

    async def ProcessGetAll(self, request: GetAllRequest, context):
        result = await self.customer_address_repo.get_all(page_n=request.pageN, page_size=request.pageSize, order_by=request.order_by, is_desc=request.is_desc)
        return GetAllCustomerAddressesResponse(customer_addresses=CustomerAddressMapper.to_model_list(result))

    async def ProcessGetById(self, request: GetCustomerAddressByIdRequest, context):
        result = await self.customer_address_repo.get_by_id(request.address_id)
        return CustomerAddressMapper.to_model(result) if result else None

    async def ProcessGetByIds(self, request: GetCustomerAddressByIdsRequest, context):
        result = await self.customer_address_repo.get_by_ids(request.address_ids)
        return GetAllCustomerAddressesResponse(customer_addresses=CustomerAddressMapper.to_model_list(result))

    async def ProcessGetCount(self, request: Empty, context):
        result = await self.customer_address_repo.get_count()
        return CountResponse(count=result)
