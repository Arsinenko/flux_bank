from abc import ABC, abstractmethod
from datetime import datetime
from typing import List

from domain.customer.customer import Customer


class CustomerRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Customer]:
        pass

    @abstractmethod
    async def get_by_id(self, customer_id: int) -> Customer | None:
        pass

    @abstractmethod
    async def get_by_ids(self, ids: List[int]) -> List[Customer]:
        pass

    @abstractmethod
    async def get_count(self) -> int:
        pass

    @abstractmethod
    async def get_by_substring(self, substring: str) -> List[Customer]:
        pass

    @abstractmethod
    async def get_by_date_range(self, from_date, to_date, page_n: int, page_size: int) -> List[Customer]:
        pass

    @abstractmethod
    async def get_count_by_substring(self, substring: str) -> int:
        pass

    @abstractmethod
    async def get_count_by_date_range(self, from_date: datetime, to_date: datetime) -> int:
        pass
    @abstractmethod
    async def get_inactive(self, threshold: datetime, page_n: int, page_size: int) -> List[Customer]:
        pass
