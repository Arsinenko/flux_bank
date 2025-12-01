using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class PaymentTemplate
{
    public int TemplateId { get; set; }

    public required int CustomerId { get; set; }

    public required string Name { get; set; }

    public required string TargetIban { get; set; }

    public decimal? DefaultAmount { get; set; }

    public required Customer Customer { get; set; }
}
