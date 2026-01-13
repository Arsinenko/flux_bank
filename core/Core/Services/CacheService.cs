using Core.Interfaces;
using Microsoft.Extensions.Caching.Memory;

namespace Core.Services;

public class CacheService : ICacheService
{
    private readonly IMemoryCache _cache;
    private readonly TimeSpan _defaultExpiration = TimeSpan.FromMinutes(10);

    public CacheService(IMemoryCache cache)
    {
        _cache = cache;
    }

    public T? Get<T>(string key)
    {
        return _cache.Get<T>(key);
    }

    public void Set<T>(string key, T value, TimeSpan? expiration = null)
    {
        var options = new MemoryCacheEntryOptions
        {
            AbsoluteExpirationRelativeToNow = expiration ?? _defaultExpiration
        };
        _cache.Set(key, value, options);
    }

    public void Remove(string key)
    {
        _cache.Remove(key);
    }
}
