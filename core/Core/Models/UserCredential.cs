using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class UserCredential
{
    public int CustomerId { get; set; }

    public string Username { get; set; } = null!;

    public string PasswordHash { get; set; } = null!;

    public DateTime? UpdatedAt { get; set; }

    public virtual Customer Customer { get; set; } = null!;
}
