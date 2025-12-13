from abc import ABC, abstractmethod
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
    async def get_by_customer_id(self, customer_id: int) -> List[Account]:
        pass

    @abstractmethod
    async def get_by_date_range(self, from_date, to_date, page_n: int, page_size: int) -> Account:
        pass
