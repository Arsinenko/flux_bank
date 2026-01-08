using AutoMapper;
using Core.Exceptions;
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
            throw new NotFoundException("Address not found");
        }
        await repository.DeleteAsync(request.AddressId);
        return new Empty();
    }

    public override async Task<GetAllCustomerAddressesResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var result = await repository.GetAllAsync(request.PageN, request.PageSize);
        return new GetAllCustomerAddressesResponse()
        {
            CustomerAddresses = { mapper.Map<IEnumerable<CustomerAddressModel>>(result) }
        };
    }

    public override async Task<Empty> Update(UpdateCustomerAddressRequest request, ServerCallContext context)
    {
        var address = await repository.GetByIdAsync(request.AddressId);
        if (address == null) throw new NotFoundException("Address not found");
        mapper.Map(request, address);
        await repository.UpdateAsync(address);
        return new Empty();
    }

    public override async Task<CustomerAddressModel> GetById(GetCustomerAddressByIdRequest request, ServerCallContext context)
    {
        var address = await repository.GetByIdAsync(request.AddressId);
        if (address == null) throw new NotFoundException("Address not found");
        return mapper.Map<CustomerAddressModel>(address);
    }

    public override async Task<GetAllCustomerAddressesResponse> GetByIds(GetCustomerAddressByIdsRequest request, ServerCallContext context)
    {
        var addresses = await repository.GetByIdsAsync(request.AddressIds);
        return new GetAllCustomerAddressesResponse()
        {
            CustomerAddresses = { mapper.Map<IEnumerable<CustomerAddressModel>>(addresses) }
        };
    }

    public override async Task<Empty> AddBulk(AddCustomerAddressBulkRequest request, ServerCallContext context)
    {
        var addresses = request.CustomerAddresses.Select(mapper.Map<CustomerAddress>).ToList();
        await repository.AddRangeAsync(addresses);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateCustomerAddressBulkRequest request, ServerCallContext context)
    {
        var addresses = request.CustomerAddresses.Select(mapper.Map<CustomerAddress>).ToList();
        if (!addresses.Any())
        {
            throw new ValidationException("No addresses to update");
        }
        await repository.UpdateRangeAsync(addresses);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteCustomerAddressBulkRequest request, ServerCallContext context)
    {
        var addresses = (await repository.GetByIdsAsync(request.CustomerAddresses.Select(a => a.AddressId))).ToList();
        if (addresses.Count != request.CustomerAddresses.Count)
        {
            throw new ValidationException("One or more addresses not found");
        }

        await repository.DeleteRangeAsync(addresses!);
        return new Empty();
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var count = await repository.GetCountAsync();
        return new CountResponse()
        {
            Count = count
        };
    }
}