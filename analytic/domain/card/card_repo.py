from abc import ABC, abstractmethod
from typing import List

from domain.branch.branch import Branch
from domain.card.card import Card


class CardRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[Card]:
        pass
    @abstractmethod
    async def get_by_id(self, branch_id: int) ->Card | None:
        pass
    @abstractmethod
    async def get_by_account_id(self, account_id: int) -> List[Card]:
        pass
    @abstractmethod
    async def get_count(self) -> int:
        pass

    @abstractmethod
    async def get_count_by_status(self, status: str):
        pass