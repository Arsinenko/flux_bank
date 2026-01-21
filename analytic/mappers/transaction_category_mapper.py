from typing import List
from api.generated.transaction_category_pb2 import TransactionCategoryModel
from domain.transaction.transaction_category import TransactionCategory


class TransactionCategoryMapper:
    @staticmethod
    def to_domain(model: TransactionCategoryModel) -> TransactionCategory:
        return TransactionCategory(
            category_id=model.category_id,
            name=model.name
        )

    @staticmethod
    def to_model(domain: TransactionCategory) -> TransactionCategoryModel:
        return TransactionCategoryModel(
            category_id=domain.category_id,
            name=domain.name
        )

    @staticmethod
    def to_domain_list(models: List[TransactionCategoryModel]) -> List[TransactionCategory]:
        return [TransactionCategoryMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[TransactionCategory]) -> List[TransactionCategoryModel]:
        return [TransactionCategoryMapper.to_model(domain) for domain in domains]
