from typing import List
from api.generated.card_pb2 import CardModel
from domain.card.card import Card

class CardMapper:
    @staticmethod
    def to_domain(model: CardModel) -> Card:
        return Card(
            card_id=model.card_id,
            account_id=model.account_id,
            card_number=model.card_number,
            cvv=model.cvv,
            expiry_date=model.expiry_date.ToDatetime() if model.HasField("expiry_date") else None,
            status=model.status
        )

    @staticmethod
    def to_model(domain: Card) -> CardModel:
        model = CardModel(
            card_id=domain.card_id,
            account_id=domain.account_id,
            card_number=domain.card_number,
            cvv=domain.cvv,
            status=domain.status
        )
        if domain.expiry_date:
            model.expiry_date.FromDatetime(domain.expiry_date)
        return model

    @staticmethod
    def to_domain_list(models: List[CardModel]) -> List[Card]:
        return [CardMapper.to_domain(model) for model in models]

    @staticmethod
    def to_model_list(domains: List[Card]) -> List[CardModel]:
        return [CardMapper.to_model(domain) for domain in domains]
