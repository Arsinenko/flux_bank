from abc import ABC, abstractmethod
from typing import List

from domain.branch.branch import Branch


class BranchRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[Branch]:
        pass

    @abstractmethod
    async def get_by_id(self, branch_id: int) ->Branch | None:
        pass
    
    @abstractmethod
    async def get_by_ids(self, ids: List[int]) -> List[Branch]:
        pass

    @abstractmethod
    async def get_count(self) -> int:
        pass
    