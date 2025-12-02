using AutoMapper;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class CustomerService(ICustomerRepository repository, IMapper mapper) : Core.CustomerService.CustomerServiceBase
{
    public override async Task<CustomerModel> Add(AddCustomerRequest request, ServerCallContext context)
    {
        var customer = mapper.Map<Customer>(request);
        await repository.AddAsync(customer);
        return mapper.Map<CustomerModel>(customer);
    }

    public override async Task<Empty> Delete(DeleteCustomerRequest request, ServerCallContext context)
    {
        var customer = await repository.GetByIdAsync(request.CustomerId);
        if (customer == null)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Customer not found."));
        }
        await repository.DeleteAsync(request.CustomerId);
        return new Empty();
    }

    public override async Task<GetAllCustomersResponse> GetAll(Empty request, ServerCallContext context)
    {
        var result = await repository.GetAllAsync();
        return new GetAllCustomersResponse()
        {
            Customers = { mapper.Map<IEnumerable<CustomerModel>>(result) }
        };
    }

    public override async Task<CustomerModel> GetById(GetCustomerByIdRequest request, ServerCallContext context)
    {
        var customer = await repository.GetByIdAsync(request.CustomerId);
        if (customer == null) throw new RpcException(new Status(StatusCode.NotFound, "Customer not found"));
        return mapper.Map<CustomerModel>(customer);
    }

    public override async Task<Empty> Update(UpdateCustomerRequest request, ServerCallContext context)
    {
        var customer = await repository.GetByIdAsync(request.CustomerId);
        if (customer == null) throw new RpcException(new Status(StatusCode.NotFound, "Customer not found"));
        mapper.Map(request, customer);
        await repository.UpdateAsync(customer);
        return new Empty();
    }
}
