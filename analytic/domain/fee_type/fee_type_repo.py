from abc import ABC, abstractmethod
from typing import List

from domain.fee_type.fee_type import FeeType


class FeeTypeRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[FeeType]:
        pass

    @abstractmethod
    async def get_by_id(self, fee_id: int) -> FeeType | None:
        pass
