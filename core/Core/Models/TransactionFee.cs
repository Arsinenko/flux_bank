using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class TransactionFee
{
    public int Id { get; set; }

    public int? TransactionId { get; set; }

    public int? FeeId { get; set; }

    public decimal? Amount { get; set; }

    public virtual FeeType? Fee { get; set; }

    public virtual Transaction? Transaction { get; set; }
}
