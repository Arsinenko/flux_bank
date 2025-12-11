from datetime import date
from decimal import Decimal


class LoanPayment:
    def __init__(self,
                 payment_id: int,
                 loan_id: int | None,
                 amount: Decimal | None,
                 payment_date: date | None,
                 is_paid: bool | None):
        self.payment_id = payment_id
        self.loan_id = loan_id
        self.amount = amount
        self.payment_date = payment_date
        self.is_paid = is_paid
