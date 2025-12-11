from datetime import datetime


class Atm:
    def __init__(self,
                 atm_id: int,
                 location: str,
                 status: str,
                 branch_id: int):
        self.atm_id: int = atm_id
        self.location: str = location
        self.status: str = status
        self.branch_id = branch_id