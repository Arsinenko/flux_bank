from typing import List
from api.generated.account_pb2 import AccountModel
from domain.account.account import Account

class AccountMapper:
    @staticmethod
    def to_domain(model: AccountModel) -> Account:
        return Account(
            account_id=model.account_id,
            customer_id=model.customer_id,
            type_id=model.type_id,
            is_active=model.is_active,
            created_at=model.created_at.ToDatetime(),
            balance=model.balance,
            iban=model.iban
        )

    @staticmethod
    def to_model(domain: Account) -> AccountModel:
        model = AccountModel(
            account_id=domain.account_id,
            customer_id=domain.customer_id,
            type_id=domain.type_id,
            iban=domain.iban,
            balance=domain.balance,
            is_active=domain.is_active
        )
        if domain.created_at:
            model.created_at.FromDatetime(domain.created_at)
        return model

    @staticmethod
    def to_domain_list(models: List[AccountModel]) -> List[Account]:
        return [AccountMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[Account]) -> List[AccountModel]:
        return [AccountMapper.to_model(domain) for domain in domains]
