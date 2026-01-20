from abc import ABC, abstractmethod
from datetime import datetime
from decimal import Decimal
from typing import List

from domain.transaction.transaction import Transaction


class TransactionRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Transaction]:
        pass

    @abstractmethod
    async def get_by_id(self, transaction_id: int) -> Transaction | None:
        pass

    @abstractmethod
    async def get_by_date_range(self, start_date: datetime, end_date: datetime) -> List[Transaction]:
        pass

    @abstractmethod
    async def get_by_ids(self, ids: List[int]) -> List[Transaction]:
        pass
    @abstractmethod
    async def get_revenue(self, account_id) -> List[Transaction]:
        pass

    @abstractmethod
    async def get_account_expenses(self, account_id) -> List[Transaction]:
        pass

    @abstractmethod
    async def get_count(self) -> int:
        pass

    @abstractmethod
    async def get_count_by_date_range(self, start_date: datetime, end_date: datetime) -> int:
        pass

    @abstractmethod
    async def get_count_revenue(self, account: int, start_date: datetime = None, end_date: datetime = None) -> int:
        pass

    @abstractmethod
    async def get_count_expenses(self, account: int, start_date: datetime = None, end_date: datetime = None) -> int:
        pass

    @abstractmethod
    async def get_total_amount(self) -> Decimal:
        pass