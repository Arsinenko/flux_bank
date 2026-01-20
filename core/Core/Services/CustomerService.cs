using AutoMapper;
using Core.Exceptions;
using Core.Interfaces;
using Core.Models;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Core.Services;

public class CustomerService(ICustomerRepository repository, IMapper mapper, ICacheService cacheService, IStatsService statsService) : Core.CustomerService.CustomerServiceBase
{
    public override async Task<CustomerModel> Add(AddCustomerRequest request, ServerCallContext context)
    {
        var customer = mapper.Map<Customer>(request);
        await repository.AddAsync(customer);
        cacheService.Remove("BankStats");
        return mapper.Map<CustomerModel>(customer);
    }

    public override async Task<Empty> Delete(DeleteCustomerRequest request, ServerCallContext context)
    {
        var customer = await repository.GetByIdAsync(request.CustomerId);
        if (customer == null)
        {
            throw new NotFoundException("Customer not found");
        }
        await repository.DeleteAsync(request.CustomerId);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<GetAllCustomersResponse> GetAll(GetAllRequest request, ServerCallContext context)
    {
        var result = await repository.GetAllAsync(request.PageN, request.PageSize, request.OrderBy, request.IsDesc ?? false);
        return new GetAllCustomersResponse()
        {
            Customers = { mapper.Map<IEnumerable<CustomerModel>>(result) }
        };
    }

    public override async Task<CustomerModel> GetById(GetCustomerByIdRequest request, ServerCallContext context)
    {
        var customer = await repository.GetByIdAsync(request.CustomerId);
        if (customer == null) throw new NotFoundException("Customer not found");
        return mapper.Map<CustomerModel>(customer);
    }

    public override async Task<GetAllCustomersResponse> GetBySubstring(GetBySubstringRequest request, ServerCallContext context)
    {
        var customers = await repository.GetBySubstringAsync(request.SubStr, request.PageN, request.PageSize, request.Order,
            request.Desc);
        return new GetAllCustomersResponse()
        {
            Customers = { mapper.Map<IEnumerable<CustomerModel>>(customers) }
        };
    }

    public override async Task<GetAllCustomersResponse> GetByDateRange(GetByDateRangeRequest request, ServerCallContext context)
    {
        var customers = await repository.GetByDateRange(request.FromDate.ToDateTime(), request.ToDate.ToDateTime(), request.PageN, request.PageSize);
        return new GetAllCustomersResponse()
        {
            Customers = { mapper.Map<IEnumerable<CustomerModel>>(customers) }
        };
    }

    public override async Task<Empty> Update(UpdateCustomerRequest request, ServerCallContext context)
    {
        var customer = await repository.GetByIdAsync(request.CustomerId);
        if (customer == null) throw new NotFoundException("Customer not found");
        mapper.Map(request, customer);
        await repository.UpdateAsync(customer);
        return new Empty();
    }

    public override async Task<Empty> DeleteBulk(DeleteCustomerBulkRequest request, ServerCallContext context)
    {
        var ids = request.Customers.Select(c => c.CustomerId).ToList();
        if (ids.Count == 0)
        {
            throw new ValidationException("No customers to delete");
        }
        var customers = await repository.GetByIdsAsync(ids);
        var foundCustomers = customers.Where(c => c != null).ToList();
        if (foundCustomers.Count != ids.Count)
        {
            throw new ValidationException("Some customers not found");
        }
        await repository.DeleteRangeAsync(foundCustomers!);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<Empty> UpdateBulk(UpdateCustomerBulkRequest request, ServerCallContext context)
    {
        var customers = request.Customers.Select(mapper.Map<Customer>).ToList();
        if (!customers.Any())
        {
            throw new ValidationException("No customers to update");
        }

        await repository.UpdateRangeAsync(customers);
        return new Empty();
    }

    public override async Task<GetAllCustomersResponse> GetByIds(GetCustomerByIdsRequest request, ServerCallContext context)
    {
        var customers = await repository.GetByIdsAsync(request.CustomerIds);
        return new GetAllCustomersResponse()
        {
            Customers = { mapper.Map<IEnumerable<CustomerModel>>(customers) }
        };
    }

    public override async Task<CountResponse> GetCount(Empty request, ServerCallContext context)
    {
        var stats = await statsService.GetStatsAsync();
        return new CountResponse()
        {
            Count = stats.CustomerCount
        };
    }

    public override async Task<CountResponse> GetCountByDateRange(GetByDateRangeRequest request, ServerCallContext context)
    {
        var count = await repository.GetCountAsync(c => c.CreatedAt >= request.FromDate.ToDateTime() && c.CreatedAt <= request.ToDate.ToDateTime());
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<CountResponse> GetCountBySubstring(GetBySubstringRequest request, ServerCallContext context)
    {
        var count = await repository.GetCountAsync(c => c.FirstName.Contains(request.SubStr) || c.LastName.Contains(request.SubStr) ||
                                                        c.Email.Contains(request.SubStr) ||
                                                        c.Phone.Contains(request.SubStr));
        return new CountResponse()
        {
            Count = count
        };
    }

    public override async Task<Empty> AddBulk(AddCustomerBulkRequest request, ServerCallContext context)
    {
        var customers = request.Customers.Select(mapper.Map<Customer>).ToList();
        await repository.AddRangeAsync(customers);
        cacheService.Remove("BankStats");
        return new Empty();
    }

    public override async Task<GetAllCustomersResponse> GetInactive(GetInactiveCustomersRequest request, ServerCallContext context)
    {
        var threshold = request.ThresholdTime.ToDateTime();
        var result = await repository.FindAsync(c => !c.LoginLogs.Any(l => l.LoginTime > threshold), request.PageN, request.PageSize);
        return new GetAllCustomersResponse()
        {
            Customers = { mapper.Map<IEnumerable<CustomerModel>>(result) }
        };
    }
}
