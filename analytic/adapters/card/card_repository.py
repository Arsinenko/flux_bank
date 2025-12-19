from typing import List

import grpc

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.card_pb2 import *
from api.generated.card_pb2_grpc import CardServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from domain.card.card import Card
from domain.card.card_repo import CardRepositoryAbc


class CardRepository(CardRepositoryAbc, BaseGrpcRepository):
    def __init__(self, target: str):
        super().__init__(target)
        self.stub = CardServiceStub(channel=self.chanel)

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
    def response_to_list(response: GetAllCardsResponse) -> List[Card]:
        return [CardRepository.to_domain(model) for model in response.cards]

    async def get_all(self, page_n: int, page_size: int) -> List[Card]:
        request = GetAllRequest(pageN=page_n, pageSize=page_size)
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return self.response_to_list(result)
        return []

    async def get_by_id(self, branch_id: int) -> Card | None:
        request = GetCardByIdRequest(card_id=branch_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return self.to_domain(result)
        return None


    async def get_by_account_id(self, account_id: int) -> List[Card]:
        request = GetCardsByAccountRequest(account_id=account_id)
        result = await self._execute(self.stub.GetByAccountId(request))
        if result:
            return self.response_to_list(result)
        return []
