from datetime import datetime, date


class Customer:
    def __init__(self,
                 customer_id: int,
                 first_name: str,
                 last_name: str,
                 email: str,
                 phone: str | None,
                 birth_date: date | None,
                 created_at: datetime | None):
        self.customer_id = customer_id
        self.first_name = first_name
        self.last_name = last_name
        self.email = email
        self.phone = phone
        self.birth_date = birth_date
        self.created_at = created_at
