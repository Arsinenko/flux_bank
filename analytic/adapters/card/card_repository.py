from typing import List

import grpc
from google.protobuf.empty_pb2 import Empty

from adapters.base_grpc_repository import BaseGrpcRepository
from api.generated.card_pb2 import *
from api.generated.card_pb2_grpc import CardServiceStub
from api.generated.custom_types_pb2 import GetAllRequest
from google.protobuf.wrappers_pb2 import StringValue, BoolValue
from domain.card.card import Card
from domain.card.card_repo import CardRepositoryAbc


from mappers.card_mapper import CardMapper


class CardRepository(CardRepositoryAbc, BaseGrpcRepository):
    def __init__(self, channel):
        super().__init__(channel)
        self.stub = CardServiceStub(channel=self.channel)

    async def get_count(self) -> int:
        result = await self._execute(self.stub.GetCount(Empty()))
        return result.count

    async def get_count_by_status(self, status: str):
        result = await self._execute(self.stub.GetCountByStatus(GetCardCountByStatus(status=status)))
        return result.count

    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Card]:
        request = GetAllRequest(
            pageN=page_n,
            pageSize=page_size,
            order_by=StringValue(value=order_by) if order_by else None,
            is_desc=BoolValue(value=is_desc)
        )
        result = await self._execute(self.stub.GetAll(request))
        if result:
            return CardMapper.to_domain_list(result.cards)
        return []

    async def get_by_id(self, branch_id: int) -> Card | None:
        request = GetCardByIdRequest(card_id=branch_id)
        result = await self._execute(self.stub.GetById(request))
        if result:
            return CardMapper.to_domain(result)
        return None


    async def get_by_account_id(self, account_id: int) -> List[Card]:
        request = GetCardsByAccountRequest(account_id=account_id)
        result = await self._execute(self.stub.GetByAccount(request))
        if result:
            return CardMapper.to_domain_list(result.cards)
        return []
