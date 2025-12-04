using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class LoanService(ILoanRepository loanRepository, IMapper mapper)
    : Core.LoanService.LoanServiceBase
{
    public override async Task<GetAllLoansResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var loans = await loanRepository.GetAllAsync(request.PageN, request.PageSize);

        return new GetAllLoansResponse
        {
            Loans = { mapper.Map<IEnumerable<LoanModel>>(loans) }
        };
    }

    public override async Task<LoanModel> Add(AddLoanRequest request, ServerCallContext context)
    {
        var loan = mapper.Map<Loan>(request);

        await loanRepository.AddAsync(loan);

        return mapper.Map<LoanModel>(loan);
    }

    public override async Task<LoanModel> GetById(GetLoanByIdRequest request, ServerCallContext context)
    {
        var loan = await loanRepository.GetByIdAsync(request.LoanId);

        if (loan == null)
            throw new RpcException(new Status(StatusCode.NotFound, "Loan not found"));

        return mapper.Map<LoanModel>(loan);
    }

    public override async Task<Empty> Update(UpdateLoanRequest request, ServerCallContext context)
    {
        var loan = await loanRepository.GetByIdAsync(request.LoanId);

        if (loan == null)
            throw new RpcException(new Status(StatusCode.NotFound, "Loan not found"));

        mapper.Map(request, loan);

        await loanRepository.UpdateAsync(loan);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteLoanRequest request, ServerCallContext context)
    {
        await loanRepository.DeleteAsync(request.LoanId);
        return new Empty();
    }
}
