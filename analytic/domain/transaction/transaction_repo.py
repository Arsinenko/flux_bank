from abc import ABC, abstractmethod
from datetime import datetime
from typing import List

from domain.transaction.transaction import Transaction


class TransactionRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[Transaction]:
        pass

    @abstractmethod
    async def get_by_id(self, transaction_id: int) -> Transaction | None:
        pass

    @abstractmethod
    async def get_by_date_range(self, start_date: datetime, end_date: datetime) -> List[Transaction]:
        pass
