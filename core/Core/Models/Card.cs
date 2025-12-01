using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Card
{
    public int CardId { get; set; }

    public int? AccountId { get; set; }

    public string CardNumber { get; set; } = null!;

    public string Cvv { get; set; } = null!;

    public DateOnly ExpiryDate { get; set; }

    public string? Status { get; set; }

    public virtual Account? Account { get; set; }
}
