from abc import ABC, abstractmethod
from typing import List

from domain.customer.customer_address import CustomerAddress


class CustomerAddressRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int, order_by: str = None, is_desc: bool = False) -> List[CustomerAddress]:
        pass

    @abstractmethod
    async def get_by_id(self, address_id: int) -> CustomerAddress | None:
        pass
