using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class LoanPaymentService(ILoanPaymentRepository loanPaymentRepository, IMapper mapper)
    : Core.LoanPaymentService.LoanPaymentServiceBase
{
    public override async Task<GetAllLoanPaymentsResponse> GetAll(Empty request, ServerCallContext context)
    {
        var loanPayments = await loanPaymentRepository.GetAllAsync();

        return new GetAllLoanPaymentsResponse
        {
            LoanPayments = { mapper.Map<IEnumerable<LoanPaymentModel>>(loanPayments) }
        };
    }

    public override async Task<LoanPaymentModel> Add(AddLoanPaymentRequest request, ServerCallContext context)
    {
        var loanPayment = mapper.Map<LoanPayment>(request);

        await loanPaymentRepository.AddAsync(loanPayment);

        return mapper.Map<LoanPaymentModel>(loanPayment);
    }

    public override async Task<LoanPaymentModel> GetById(GetLoanPaymentByIdRequest request, ServerCallContext context)
    {
        var loanPayment = await loanPaymentRepository.GetByIdAsync(request.PaymentId);

        if (loanPayment == null)
            throw new RpcException(new Status(StatusCode.NotFound, "LoanPayment not found"));

        return mapper.Map<LoanPaymentModel>(loanPayment);
    }

    public override async Task<Empty> Update(UpdateLoanPaymentRequest request, ServerCallContext context)
    {
        var loanPayment = await loanPaymentRepository.GetByIdAsync(request.PaymentId);

        if (loanPayment == null)
            throw new RpcException(new Status(StatusCode.NotFound, "LoanPayment not found"));

        mapper.Map(request, loanPayment);

        await loanPaymentRepository.UpdateAsync(loanPayment);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteLoanPaymentRequest request, ServerCallContext context)
    {
        await loanPaymentRepository.DeleteAsync(request.PaymentId);
        return new Empty();
    }
}
