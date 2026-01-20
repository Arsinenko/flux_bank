from abc import ABC, abstractmethod
from typing import List

from domain.atm.atm import Atm


class AtmRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[Atm]:
        pass
    @abstractmethod
    async def get_by_id(self, atm_id: int) -> Atm:
        pass

    @abstractmethod
    async def get_by_status(self, status: str) ->List[Atm]:
        pass
    @abstractmethod
    async def get_by_location_substr(self, sub_str: str) -> List[Atm]:
        pass

    @abstractmethod
    async def get_by_branch(self, branch_id: int) -> List[Atm]:
        pass

    @abstractmethod
    async def get_count(self) -> int:
        pass

    @abstractmethod
    async def get_count_by_status(self, status: str) -> int:
        pass