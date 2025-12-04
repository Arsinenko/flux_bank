using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class CustomerAddressService(ICustomerAddressRepository repository, IMapper mapper) : Core.CustomerAddressService.CustomerAddressServiceBase
{
    public override async Task<CustomerAddressModel> Add(AddCustomerAddressRequest request, ServerCallContext context)
    {
        var address = mapper.Map<CustomerAddress>(request);
        await repository.AddAsync(address);
        return mapper.Map<CustomerAddressModel>(address); 
    }

    public override async Task<Empty> Delete(DeleteCustomerAddressRequest request, ServerCallContext context)
    {
        var address = await repository.GetByIdAsync(request.AddressId);
        if (address == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Address not found."));
        }
        await repository.DeleteAsync(request.AddressId);
        return new Empty();
    }

    public override async Task<GetAllCustomerAddressesResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var result = await repository.GetAllAsync(request.PageN, request.PageSize);
        var addresses = new RepeatedField<CustomerAddressModel>();
        return new GetAllCustomerAddressesResponse()
        {
            CustomerAddresses = { mapper.Map<IEnumerable<CustomerAddressModel>>(result) }
        };
    }

    public override async Task<Empty> Update(UpdateCustomerAddressRequest request, ServerCallContext context)
    {
        var address = await repository.GetByIdAsync(request.AddressId);
        if (address == null) throw new RpcException(new Status(StatusCode.NotFound, "Address not found"));
        mapper.Map(request, address);
        await repository.UpdateAsync(address);
        return new Empty();
    }

    public override async Task<CustomerAddressModel> GetById(GetCustomerAddressByIdRequest request, ServerCallContext context)
    {
        var address = await repository.GetByIdAsync(request.AddressId);
        if (address == null) throw new RpcException(new Status(StatusCode.NotFound, "Address not found"));
        return mapper.Map<CustomerAddressModel>(address);
    }
}