using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Loan
{
    public int LoanId { get; set; }

    public required int CustomerId { get; set; }

    public required decimal Principal { get; set; }

    public required decimal InterestRate { get; set; }

    public required DateOnly StartDate { get; set; }

    public DateOnly? EndDate { get; set; }

    public required string Status { get; set; }

    public required Customer Customer { get; set; }

    public virtual ICollection<LoanPayment> LoanPayments { get; set; } = new List<LoanPayment>();
}
