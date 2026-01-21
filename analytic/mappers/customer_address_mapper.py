from typing import List
from api.generated.customer_address_pb2 import CustomerAddressModel
from domain.customer.customer_address import CustomerAddress


class CustomerAddressMapper:
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
    def to_domain_list(models: List[CustomerAddressModel]) -> List[CustomerAddress]:
        return [CustomerAddressMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[CustomerAddress]) -> List[CustomerAddressModel]:
        return [CustomerAddressMapper.to_model(domain) for domain in domains]
