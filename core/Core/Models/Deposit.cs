using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Deposit
{
    public int DepositId { get; set; }

    public required int CustomerId { get; set; }

    public required decimal Amount { get; set; }

    public required decimal InterestRate { get; set; }

    public required DateOnly StartDate { get; set; }

    public DateOnly? EndDate { get; set; }

    public required string Status { get; set; }

    public required Customer Customer { get; set; }
}
