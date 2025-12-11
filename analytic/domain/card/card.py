from datetime import datetime


class Card:
    def __init__(self,
                 card_id: int,
                 account_id: int,
                 card_number: str,
                 cvv: str,
                 expiry_date: datetime,
                 status: str):
        self.card_id = card_id
        self.account_id = account_id
        self.card_number = card_number
        self.cvv = cvv
        self.expiry_date = expiry_date
        self.status = status
