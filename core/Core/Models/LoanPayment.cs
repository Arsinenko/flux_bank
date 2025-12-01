using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class LoanPayment
{
    public int PaymentId { get; set; }

    public required int LoanId { get; set; }

    public required decimal Amount { get; set; }

    public required DateOnly PaymentDate { get; set; }

    public required bool IsPaid { get; set; }

    public required Loan Loan { get; set; }
}
