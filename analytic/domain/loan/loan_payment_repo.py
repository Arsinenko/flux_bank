from abc import ABC, abstractmethod
from typing import List

from domain.loan.loan_payment import LoanPayment


class LoanPaymentRepositoryAbc(ABC):
    @abstractmethod
    async def get_all(self, page_n: int, page_size: int) -> List[LoanPayment]:
        pass

    @abstractmethod
    async def get_by_id(self, payment_id: int) -> LoanPayment | None:
        pass

    @abstractmethod
    async def get_by_loan_id(self, loan_id: int) -> List[LoanPayment]:
        pass
