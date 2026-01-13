import decimal
from abc import ABC, abstractmethod
from datetime import datetime
from typing import List

from domain.account.account import Account


class AccountRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[Account]:
        pass

    @abstractmethod
    async def get_by_id(self, account_id: int) -> Account:
        pass

    @abstractmethod
    async def get_by_ids(self, ids: List[int]) -> List[Account]:
        pass

    @abstractmethod
    async def get_by_customer_id(self, customer_id: int) -> List[Account]:
        pass

    @abstractmethod
    async def get_by_date_range(self, from_date, to_date, page_n: int, page_size: int) -> List[Account]:
        pass

    @abstractmethod
    async def get_count(self) -> int:
        pass

    @abstractmethod
    async def get_count_by_status(self, status: bool) -> int:
        pass

    @abstractmethod
    async def get_count_by_date_range(self, from_date: datetime, to_date: datetime) -> int:
        pass

    @abstractmethod
    async def get_count_by_customer_id(self, customer_id: int) -> int:
        pass

    @abstractmethod
    async def get_total_balance(self) -> decimal.Decimal:
        pass
