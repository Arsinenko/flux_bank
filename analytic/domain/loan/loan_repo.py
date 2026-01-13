from abc import ABC, abstractmethod
from typing import List

from domain.loan.loan import Loan


class LoanRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[Loan]:
        pass

    @abstractmethod
    async def get_by_id(self, loan_id: int) -> Loan | None:
        pass

    @abstractmethod
    async def get_by_customer_id(self, customer_id: int) -> List[Loan]:
        pass

    @abstractmethod
    async def get_by_ids(self, ids: List[int]) -> List[Loan]:
        pass

    @abstractmethod
    async def get_count(self) -> int:
        pass

    @abstractmethod
    async def get_count_by_status(self, status: str) -> int:
        pass
