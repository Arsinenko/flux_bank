using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class ExchangeRate
{
    public int RateId { get; set; }

    public string? BaseCurrency { get; set; }

    public string? TargetCurrency { get; set; }

    public decimal? Rate { get; set; }

    public DateTime? UpdatedAt { get; set; }
}
