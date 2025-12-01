using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Notification
{
    public int NotificationId { get; set; }

    public int? CustomerId { get; set; }

    public string? Message { get; set; }

    public DateTime? CreatedAt { get; set; }

    public bool? IsRead { get; set; }

    public virtual Customer? Customer { get; set; }
}
