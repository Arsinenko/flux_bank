using System;
using System.Collections.Generic;
using DateOnly = global::System.DateOnly;

namespace Core.Models;

public partial class Deposit
{
    public int DepositId { get; set; }

    public required int CustomerId { get; set; }

    public required decimal Amount { get; set; }

    public required decimal InterestRate { get; set; }

    public required System.DateOnly StartDate { get; set; }

    public System.DateOnly? EndDate { get; set; }

    public required string Status { get; set; }

    public required Customer Customer { get; set; }
}
