using Core.Interfaces;
using Core.Models;
using Google.Protobuf.Collections;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class CustomerAddressService(ICustomerAddressRepository repository) : Core.CustomerAddressService.CustomerAddressServiceBase
{
    public override async Task<CustomerAddressModel> Add(AddCustomerAddressRequest request, ServerCallContext context)
    {
        var address = new CustomerAddress
        {
            CustomerId = request.CustomerId,
            Country = request.Country,
            City = request.City,
            Street = request.Street,
            ZipCode = request.ZipCode,
            IsPrimary = request.IsPrimary,
        };
        await repository.AddAsync(address);
        return new CustomerAddressModel
        {
            AddressId = address.AddressId,
            CustomerId = address.CustomerId,
            Country = address.Country,
            City = address.City,
            Street = address.Street,
            ZipCode = address.ZipCode,
            IsPrimary = address.IsPrimary,
        };
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

    public override async Task<GetAllCustomerAddressesResponse> GetAll(Empty request, ServerCallContext context)
    {
        var result = await repository.GetAllAsync();
        var addresses = new RepeatedField<CustomerAddressModel>();
        foreach (var address in result)
        {
            addresses.Add(new CustomerAddressModel
            {
                AddressId = address.AddressId,
                CustomerId = address.CustomerId,
                Country = address.Country,
                City = address.City,
                Street = address.Street,
                ZipCode = address.ZipCode,
                IsPrimary = address.IsPrimary
            });
        }
        return new GetAllCustomerAddressesResponse{CustomerAddresses = { addresses }};
    }

    public override async Task<Empty> Update(UpdateCustomerAddressRequest request, ServerCallContext context)
    {
        var address = await repository.GetByIdAsync(request.AddressId);
        if (address == null) throw new RpcException(new Status(StatusCode.NotFound, "Address not found"));
        address.Country = request.Country;
        address.City = request.City;
        address.Street = request.Street;
        address.ZipCode = request.ZipCode;
        address.IsPrimary = request.IsPrimary;
        await repository.UpdateAsync(address);
        return new Empty();
    }

    public override async Task<CustomerAddressModel> GetById(GetCustomerAddressByIdRequest request, ServerCallContext context)
    {
        var address = await repository.GetByIdAsync(request.AddressId);
        if (address == null) throw new RpcException(new Status(StatusCode.NotFound, "Address not found"));
        return new CustomerAddressModel
        {
            AddressId = address.AddressId,
            CustomerId = address.CustomerId,
            Country = address.Country,
            City = address.City,
            Street = address.Street,
            ZipCode = address.ZipCode,
            IsPrimary = address.IsPrimary,
        };  
    }
    
}