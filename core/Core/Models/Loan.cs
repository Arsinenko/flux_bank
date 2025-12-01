using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Loan
{
    public int LoanId { get; set; }

    public int? CustomerId { get; set; }

    public decimal? Principal { get; set; }

    public decimal? InterestRate { get; set; }

    public DateOnly? StartDate { get; set; }

    public DateOnly? EndDate { get; set; }

    public string? Status { get; set; }

    public virtual Customer? Customer { get; set; }

    public virtual ICollection<LoanPayment> LoanPayments { get; set; } = new List<LoanPayment>();
}
