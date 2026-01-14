from decimal import Decimal
from typing import List
from api.generated.transaction_pb2 import TransactionModel
from domain.transaction.transaction import Transaction

class TransactionMapper:
    @staticmethod
    def to_domain(model: TransactionModel) -> Transaction:
        return Transaction(
            transaction_id=model.transaction_id,
            source_account=model.source_account if model.HasField("source_account") else None,
            target_account=model.target_account if model.HasField("target_account") else None,
            amount=Decimal(model.amount),
            currency=model.currency,
            created_at=model.created_at.ToDatetime() if model.HasField("created_at") else None,
            status=model.status if model.HasField("status") else None
        )

    @staticmethod
    def to_model(domain: Transaction) -> TransactionModel:
        model = TransactionModel(
            transaction_id=domain.transaction_id,
            source_account=domain.source_account,
            target_account=domain.target_account,
            amount=str(domain.amount),
            currency=domain.currency,
            status=domain.status
        )
        if domain.created_at:
            model.created_at.FromDatetime(domain.created_at)
        return model

    @staticmethod
    def to_domain_list(models: List[TransactionModel]) -> List[Transaction]:
        return [TransactionMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[Transaction]) -> List[TransactionModel]:
        return [TransactionMapper.to_model(domain) for domain in domains]
