class FeeType:
    def __init__(self,
                 fee_id: int,
                 name: str | None,
                 description: str | None):
        self.fee_id = fee_id
        self.name = name
        self.description = description
