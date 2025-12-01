using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class FeeType
{
    public int FeeId { get; set; }

    public string? Name { get; set; }

    public string? Description { get; set; }

    public virtual ICollection<TransactionFee> TransactionFees { get; set; } = new List<TransactionFee>();
}
