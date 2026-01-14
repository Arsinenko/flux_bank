from typing import List
from api.generated.fee_type_pb2 import FeeTypeModel
from domain.fee_type.fee_type import FeeType

class FeeTypeMapper:
    @staticmethod
    def to_domain(model: FeeTypeModel) -> FeeType:
        return FeeType(
            fee_id=model.fee_id,
            name=model.name if model.HasField("name") else None,
            description=model.description if model.HasField("description") else None
        )

    @staticmethod
    def to_model(domain: FeeType) -> FeeTypeModel:
        return FeeTypeModel(
            fee_id=domain.fee_id,
            name=domain.name,
            description=domain.description
        )

    @staticmethod
    def to_domain_list(models: List[FeeTypeModel]) -> List[FeeType]:
        return [FeeTypeMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[FeeType]) -> List[FeeTypeModel]:
        return [FeeTypeMapper.to_model(domain) for domain in domains]
