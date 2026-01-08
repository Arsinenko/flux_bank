using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class LoanPaymentService(ILoanPaymentRepository loanPaymentRepository, IMapper mapper)
    : Core.LoanPaymentService.LoanPaymentServiceBase
{
    public override async Task<GetAllLoanPaymentsResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var loanPayments = await loanPaymentRepository.GetAllAsync(request.PageN, request.PageSize);

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
            throw new NotFoundException("LoanPayment not found");

        return mapper.Map<LoanPaymentModel>(loanPayment);
    }

    public override async Task<Empty> Update(UpdateLoanPaymentRequest request, ServerCallContext context)
    {
        var loanPayment = await loanPaymentRepository.GetByIdAsync(request.PaymentId);

        if (loanPayment == null)
            throw new NotFoundException("LoanPayment not found");

        mapper.Map(request, loanPayment);

        await loanPaymentRepository.UpdateAsync(loanPayment);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteLoanPaymentRequest request, ServerCallContext context)
    {
        await loanPaymentRepository.DeleteAsync(request.PaymentId);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteLoanPaymentBulkRequest request, ServerCallContext context)
    {
        var ids = request.Payments.Select(p => p.PaymentId).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No payments to delete");
        }
        var loanPayments = await loanPaymentRepository.GetByIdsAsync(ids);
        var foundLoanPayments = loanPayments.Where(lp => lp != null).ToList();
        if (foundLoanPayments.Count != ids.Count)
        {
            throw new ValidationException("Some loan payments not found");
        }
        await loanPaymentRepository.DeleteRangeAsync(foundLoanPayments!);
        return new Empty();
    }

    public override async Task<GetAllLoanPaymentsResponse> GetByLoan(GetLoanPaymentsByLoanRequest request, ServerCallContext context)
    {
        var loanPayments = await loanPaymentRepository.FindAsync(lp => lp.LoanId == request.LoanId);
        return new GetAllLoanPaymentsResponse()
        {
            LoanPayments = { mapper.Map<LoanPaymentModel>(loanPayments) }
        };
    }
    public override async Task<Empty> UpdateBulk(UpdateLoanPaymentBulkRequest request, ServerCallContext context)
    {
        var loanPayments = request.Payments.Select(mapper.Map<LoanPayment>).ToList();
        if (!loanPayments.Any())
        {
            throw new ValidationException("No loan payments to update");
        }
        await loanPaymentRepository.UpdateRangeAsync(loanPayments);
        return new Empty();
    }

    public override async Task<Empty> AddBulk(AddLoanPaymentBulkRequest request, ServerCallContext context)
    {
        var loanPayments = request.Payments.Select(mapper.Map<LoanPayment>).ToList();
        await loanPaymentRepository.AddRangeAsync(loanPayments);
        return new Empty();
    }

    public override async Task<GetAllLoanPaymentsResponse> GetByIds(GetLoanPaymentByIdsRequest request, ServerCallContext context)
    {
        var loanPayments = await loanPaymentRepository.GetByIdsAsync(request.PaymentIds);
        return new GetAllLoanPaymentsResponse()
        {
            LoanPayments = { mapper.Map<IEnumerable<LoanPaymentModel>>(loanPayments) }
        };
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var count = await loanPaymentRepository.GetCountAsync();
        return new CountResponse()
        {
            Count = count
        };
    }
    
}
