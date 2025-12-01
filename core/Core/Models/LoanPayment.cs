using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class LoanPayment
{
    public int PaymentId { get; set; }

    public int? LoanId { get; set; }

    public decimal? Amount { get; set; }

    public DateOnly? PaymentDate { get; set; }

    public bool? IsPaid { get; set; }

    public virtual Loan? Loan { get; set; }
}
