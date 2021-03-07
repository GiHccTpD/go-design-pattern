package _1_factory_method

// 定义一个Cache接口，作为夫类
type Cache interface {
	Set(key, value string)
	Get(key string) string
}

// 实现具体的Cache:RedisCache
type RedisCache struct {
	data map[string]string
}

func NewRedisCache() *RedisCache {
	return &RedisCache{
		data: make(map[string]string),
	}
}

func (redis *RedisCache) Set(key, value string) {
	redis.data[key] = value
}

func (redis *RedisCache) Get(key string) string {
	return "redis data: " + redis.data[key]
}

// 实现具体的Cache:RedisCache
type MemCache struct {
	data map[string]string
}

func NewMemCache() *MemCache {
	return &MemCache{
		data: make(map[string]string),
	}
}

func (mem *MemCache) Set(key, value string) {
	mem.data[key] = value
}

func (mem *MemCache) Get(key string) string {
	return "mem data: " + mem.data[key]
}

// 定义一个抽象的cache工厂
type CacheFactory interface {
	Create() Cache
}

// 实现具体的工厂 redis
type RedisCacheFactory struct {
}

func (rf *RedisCacheFactory) Create() Cache {
	return NewRedisCache()
}

// 实现具体的工厂 mem
type MemCacheFactory struct {
}

func (mf *MemCacheFactory) Create() Cache {
	return NewMemCache()
}
