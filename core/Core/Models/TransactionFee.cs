using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class TransactionFee
{
    public required int Id { get; set; }

    public required int TransactionId { get; set; }

    public required int FeeId { get; set; }

    public required decimal Amount { get; set; }

    public required FeeType Fee { get; set; }

    public virtual Transaction? Transaction { get; set; }
}
