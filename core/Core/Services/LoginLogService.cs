using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class LoginLogService(ILoginLogRepository loginLogRepository, IMapper mapper)
    : Core.LoginLogService.LoginLogServiceBase
{
    public override async Task<GetAllLoginLogsResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var loginLogs = await loginLogRepository.GetAllAsync(request.PageN, request.PageSize);

        return new GetAllLoginLogsResponse
        {
            LoginLogs = { mapper.Map<IEnumerable<LoginLogModel>>(loginLogs) }
        };
    }

    public override async Task<LoginLogModel> Add(AddLoginLogRequest request, ServerCallContext context)
    {
        var loginLog = mapper.Map<LoginLog>(request);

        await loginLogRepository.AddAsync(loginLog);

        return mapper.Map<LoginLogModel>(loginLog);
    }

    public override async Task<LoginLogModel> GetById(GetLoginLogByIdRequest request, ServerCallContext context)
    {
        var loginLog = await loginLogRepository.GetByIdAsync(request.LogId);

        if (loginLog == null)
            throw new NotFoundException("LoginLog not found");

        return mapper.Map<LoginLogModel>(loginLog);
    }

    public override async Task<Empty> Update(UpdateLoginLogRequest request, ServerCallContext context)
    {
        var loginLog = await loginLogRepository.GetByIdAsync(request.LogId);

        if (loginLog == null)
            throw new NotFoundException("LoginLog not found");

        mapper.Map(request, loginLog);

        await loginLogRepository.UpdateAsync(loginLog);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteLoginLogRequest request, ServerCallContext context)
    {
        await loginLogRepository.DeleteAsync(request.LogId);
        return new Empty();
    }

    public override async Task<GetAllLoginLogsResponse> GetByCustomer(GetLoginLogsByCustomerRequest request, ServerCallContext context)
    {
        var logs = await loginLogRepository.FindAsync(l => l.CustomerId == request.CustomerId);
        return new GetAllLoginLogsResponse()
        {
            LoginLogs = { mapper.Map<LoginLogModel>(logs) }
        };
    }

    public override async Task<GetAllLoginLogsResponse> GetInTimeRange(GetLoginLogsInTimeRangeRequest request, ServerCallContext context)
    {
        var logs = await loginLogRepository.FindAsync(l =>
            l.LoginTime >= request.StartTime.ToDateTime() && l.LoginTime <= request.EndTime.ToDateTime());
        return new GetAllLoginLogsResponse()
        {
            LoginLogs = { mapper.Map<LoginLogModel>(logs) }
        };
    }

    public override async Task<GetAllLoginLogsResponse> GetByIds(GetLoginLogByIdsRequest request, ServerCallContext context)
    {
        var logs = await loginLogRepository.GetByIdsAsync(request.LogIds);
        return new GetAllLoginLogsResponse()
        {
            LoginLogs = { mapper.Map<LoginLogModel>(logs) }
        };
    }
}
