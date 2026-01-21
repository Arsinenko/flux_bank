from abc import abstractmethod, ABC
from decimal import Decimal
from typing import List

from api.generated.transaction_fee_pb2 import TransactionFeeModel
from domain.transaction.transaction_fee import TransactionFee


class TransactionFeeMapper:
    @staticmethod
    def to_domain(model: TransactionFeeModel) -> TransactionFee:
        return TransactionFee(
            id=model.id,
            transaction_id=model.transaction_id if model.HasField("transaction_id") else None,
            fee_id=model.fee_id if model.HasField("fee_id") else None,
            amount=Decimal(model.amount) if model.HasField("amount") else None
        )

    @staticmethod
    def to_model(domain: TransactionFee) -> TransactionFeeModel:
        return TransactionFeeModel(
            id=domain.id,
            transaction_id=domain.transaction_id,
            fee_id=domain.fee_id,
            amount=str(domain.amount) if domain.amount is not None else None
        )

    @abstractmethod
    def to_domain_list(self, models: List[TransactionFeeModel]):
        return [TransactionFeeMapper.to_domain(model) for model in models]

    @abstractmethod
    def to_model_list(self, domains: List[TransactionFee]):
        return [TransactionFeeMapper.to_model(domain) for domain in domains]