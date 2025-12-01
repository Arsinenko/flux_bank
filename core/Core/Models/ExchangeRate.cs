using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class ExchangeRate
{
    public int RateId { get; set; }

    public required string BaseCurrency { get; set; }

    public required string TargetCurrency { get; set; }

    public required decimal Rate { get; set; }

    public DateTime? UpdatedAt { get; set; }
}
