using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class PaymentTemplate
{
    public int TemplateId { get; set; }

    public int? CustomerId { get; set; }

    public string? Name { get; set; }

    public string? TargetIban { get; set; }

    public decimal? DefaultAmount { get; set; }

    public virtual Customer? Customer { get; set; }
}
