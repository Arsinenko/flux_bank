class PaymentTemplate:
    def __init__(self, template_id: int, customer_id: int, name: str, target_iban: str, default_amount: str):
        self.default_amount = default_amount
        self.target_iban = target_iban
        self.name = name
        self.customer_id = customer_id
        self.template_id = template_id