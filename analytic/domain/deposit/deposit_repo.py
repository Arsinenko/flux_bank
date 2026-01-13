from abc import ABC, abstractmethod
from typing import List

from domain.deposit.deposit import Deposit


class DepositRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[Deposit]:
        pass

    @abstractmethod
    async def get_by_id(self, deposit_id: int) -> Deposit | None:
        pass

    @abstractmethod
    async def get_by_customer_id(self, customer_id: int) -> List[Deposit]:
        pass

    @abstractmethod
    async def get_by_ids(self, ids: List[int]):
        pass

    @abstractmethod
    async def get_count(self) -> int:
        pass

    @abstractmethod
    async def get_count_by_status(self, status: str):
        pass