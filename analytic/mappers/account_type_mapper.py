from typing import List

from api.generated.account_type_pb2 import AccountTypeModel
from domain.account.account_type import AccountType


class AccountTypeMapper:
    @staticmethod
    def to_domain(model: AccountTypeModel) -> AccountType:
        return AccountType(
            type_id=model.type_id,
            name=model.name,
            description=model.description
        )
    @staticmethod
    def to_model(domain: AccountType) -> AccountTypeModel:
        return AccountTypeModel(
            type_id=domain.type_id,
            name=domain.name,
            description=domain.description
        )

    @staticmethod
    def to_domain_list(models: List[AccountTypeModel]) -> List[AccountType]:
        return [AccountTypeMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[AccountType]) -> List[AccountTypeModel]:
        return [AccountTypeMapper.to_model(domain) for domain in domains]