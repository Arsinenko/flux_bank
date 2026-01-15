from typing import List
from api.generated.customer_pb2 import CustomerModel
from domain.customer.customer import Customer
from google.protobuf.timestamp_pb2 import Timestamp

class CustomerMapper:
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
    def to_model(domain: Customer) -> CustomerModel:
        model = CustomerModel(
            customer_id=domain.customer_id,
            first_name=domain.first_name,
            last_name=domain.last_name,
            email=domain.email,
            phone=domain.phone
        )
        if domain.birth_date:
            model.birth_date.FromDatetime(domain.birth_date)
        if domain.created_at:
            model.created_at.FromDatetime(domain.created_at)
        return model

    @staticmethod
    def to_domain_list(models: List[CustomerModel]) -> List[Customer]:
        return [CustomerMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[Customer]) -> List[CustomerModel]:
        return [CustomerMapper.to_model(domain) for domain in domains]

    @staticmethod
    def to_timestamp(dt) -> Timestamp:
        ts = Timestamp()
        ts.FromDatetime(dt)
        return ts
