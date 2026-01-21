from abc import ABC, abstractmethod
from typing import List

from domain.transaction.transaction_category import TransactionCategory


class TransactionCategoryRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[TransactionCategory]:
        pass

    @abstractmethod
    async def get_by_id(self, category_id: int) -> TransactionCategory | None:
        pass

    @abstractmethod
    async def get_count(self) -> int:
        pass