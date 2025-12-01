using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Card
{
    public int CardId { get; set; }

    public required int AccountId { get; set; }

    public required string CardNumber { get; set; }

    public required string Cvv { get; set; }

    public DateOnly? ExpiryDate { get; set; }

    public required string Status { get; set; }

    public Account? Account { get; set; }
}
