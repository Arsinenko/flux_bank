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

    public override async Task<GetAllCustomersResponse>GetAll( GetAllRequest request, ServerCallContext context)
    {
        var result = await repository.GetAllAsync(request.PageN, request.PageSize);
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

    public override async Task<Empty> DeleteBulk(DeleteCustomerBulkRequest request, ServerCallContext context)
    {
        var ids = request.Customers.Select(c => c.CustomerId).ToList();
        if (ids.Count == 0)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No customers to delete"));
        }
        var customers = await repository.GetByIdsAsync(ids);
        var foundCustomers = customers.Where(c => c != null).ToList();
        if (foundCustomers.Count != ids.Count)
        {
            throw new RpcException(new Status(StatusCode.NotFound, "Some customers not found"));
        }
        await repository.DeleteRangeAsync(foundCustomers!);
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateCustomerBulkRequest request, ServerCallContext context)
    {
        var customers = request.Customers.Select(mapper.Map<Customer>).ToList();
        if (!customers.Any())
        {
            throw new RpcException(new Status(StatusCode.NotFound, "No customers to update"));
        }

        await repository.UpdateRangeAsync(customers);
        return new Empty();
    }
}
