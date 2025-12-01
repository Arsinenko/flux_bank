using System;
using System.Collections.Generic;

namespace Core.Models;

public partial class Notification
{
    public int NotificationId { get; set; }

    public required int CustomerId { get; set; }

    public required string Message { get; set; }

    public required DateTime CreatedAt { get; set; }

    public required bool IsRead { get; set; }

    public required Customer Customer { get; set; }
}
