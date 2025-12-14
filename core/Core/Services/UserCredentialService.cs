using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class UserCredentialService(IUserCredentialRepository userCredentialRepository, IMapper mapper)
    : Core.UserCredentialService.UserCredentialServiceBase
{
    public override async Task<GetAllUserCredentialsResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var userCredentials = await userCredentialRepository.GetAllAsync(request.PageN, request.PageSize);

        return new GetAllUserCredentialsResponse
        {
            UserCredentials = { mapper.Map<IEnumerable<UserCredentialModel>>(userCredentials) }
        };
    }

    public override async Task<UserCredentialModel> Add(AddUserCredentialRequest request, ServerCallContext context)
    {
        var userCredential = mapper.Map<UserCredential>(request);

        await userCredentialRepository.AddAsync(userCredential);

        return mapper.Map<UserCredentialModel>(userCredential);
    }

    public override async Task<UserCredentialModel> GetById(GetUserCredentialByIdRequest request, ServerCallContext context)
    {
        var userCredential = await userCredentialRepository.GetByIdAsync(request.CustomerId);

        if (userCredential == null)
            throw new RpcException(new Status(StatusCode.NotFound, "UserCredential not found"));

        return mapper.Map<UserCredentialModel>(userCredential);
    }

    public override async Task<Empty> Update(UpdateUserCredentialRequest request, ServerCallContext context)
    {
        var userCredential = await userCredentialRepository.GetByIdAsync(request.CustomerId);

        if (userCredential == null)
            throw new RpcException(new Status(StatusCode.NotFound, "UserCredential not found"));

        mapper.Map(request, userCredential);

        await userCredentialRepository.UpdateAsync(userCredential);

        return new Empty();
    }

    public override async Task<Empty> Delete(DeleteUserCredentialRequest request, ServerCallContext context)
    {
        await userCredentialRepository.DeleteAsync(request.CustomerId);
        return new Empty();
    }

    public override async Task<UserCredentialModel> GetByUsername(GetUserCredentialByUsernameRequest request, ServerCallContext context)
    {
        var creds = await userCredentialRepository.FindAsync(c => c.Username == request.Username);
        return mapper.Map<UserCredentialModel>(creds.FirstOrDefault());
    }

    public override async Task<GetAllUserCredentialsResponse> GetByIds(GetUserCredentialByIdsRequest request, ServerCallContext context)
    {
        var userCredentials = await userCredentialRepository.GetByIdsAsync(request.CustomerIds);
        return new GetAllUserCredentialsResponse()
        {
            UserCredentials = { mapper.Map<IEnumerable<UserCredentialModel>>(userCredentials) }
        };
    }
}
