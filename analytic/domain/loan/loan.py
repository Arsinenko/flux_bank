from datetime import date
from decimal import Decimal


class Loan:
    def __init__(self,
                 loan_id: int,
                 customer_id: int | None,
                 principal: Decimal | None,
                 interest_rate: Decimal | None,
                 start_date: date | None,
                 end_date: date | None,
                 status: str | None):
        self.loan_id = loan_id
        self.customer_id = customer_id
        self.principal = principal
        self.interest_rate = interest_rate
        self.start_date = start_date
        self.end_date = end_date
        self.status = status
