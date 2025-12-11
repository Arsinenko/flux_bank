from datetime import date
from decimal import Decimal


class Deposit:
    def __init__(self,
                 deposit_id: int,
                 customer_id: int | None,
                 amount: Decimal | None,
                 interest_rate: Decimal | None,
                 start_date: date | None,
                 end_date: date | None,
                 status: str | None):
        self.deposit_id = deposit_id
        self.customer_id = customer_id
        self.amount = amount
        self.interest_rate = interest_rate
        self.start_date = start_date
        self.end_date = end_date
        self.status = status
