from abc import ABC, abstractmethod
from typing import List

from domain.account.account_type import AccountType


class AccountTypeRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[AccountType]:
        pass
    @abstractmethod
    async def get_by_id(self, type_id: int) -> AccountType:
        pass