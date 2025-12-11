class CustomerAddress:
    def __init__(self,
                 address_id: int,
                 customer_id: int | None,
                 country: str | None,
                 city: str | None,
                 street: str | None,
                 zip_code: str | None,
                 is_primary: bool | None):
        self.address_id = address_id
        self.customer_id = customer_id
        self.country = country
        self.city = city
        self.street = street
        self.zip_code = zip_code
        self.is_primary = is_primary
