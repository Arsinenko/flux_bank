from abc import ABC, abstractmethod
from typing import List

from domain.user_credential.user_credential import UserCredential


class UserCredentialRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[UserCredential]:
        pass

    @abstractmethod
    async def get_by_id(self, customer_id: int) -> UserCredential | None:
        pass

    @abstractmethod
    async def get_by_ids(self, customer_ids: List[int]) -> List[UserCredential]:
        pass

    @abstractmethod
    async def get_by_username(self, username: str) -> UserCredential | None:
        pass

    @abstractmethod
    async def get_count(self) -> int:
        pass
