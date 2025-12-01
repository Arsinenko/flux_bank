using Core.Context;
using Core.Interfaces;
using Core.Models;

namespace Core.Repositories;

public class PaymentTemplateRepository : GenericRepository<PaymentTemplate, int>, IPaymentTemplateRepository
{
    public PaymentTemplateRepository(MyDbContext context) : base(context)
    {
    }
}
