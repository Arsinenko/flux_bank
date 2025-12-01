using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Deposit
{
    public int DepositId { get; set; }

    public int? CustomerId { get; set; }

    public decimal? Amount { get; set; }

    public decimal? InterestRate { get; set; }

    public DateOnly? StartDate { get; set; }

    public DateOnly? EndDate { get; set; }

    public string? Status { get; set; }

    public virtual Customer? Customer { get; set; }
}
