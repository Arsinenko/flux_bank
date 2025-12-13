from abc import ABC, abstractmethod
from typing import List

from domain.customer.customer import Customer


class CustomerRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[Customer]:
        pass

    @abstractmethod
    async def get_by_id(self, customer_id: int) -> Customer | None:
        pass
