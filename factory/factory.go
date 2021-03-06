package factory

import "errors"

// 定义一个Cache接口，作为夫类
type Cache interface {
	Set(key, value string)
	Get(key string) string
}

// 实现具体的Cache:RedisCache
type RedisCache struct {
	data map[string]string
}

func (redis * RedisCache) Set(key, value string) {
	redis.data[key] = value
}

func (redis *RedisCache) Get(key string) string {
	return "redis data: " + redis.data[key]
}

// 实现具体的Cache:RedisCache
type MemCache struct {
	data map[string]string
}

func (mem * MemCache) Set(key, value string) {
	mem.data[key] = value
}

func (mem *MemCache) Get(key string) string {
	return "mem data: " + mem.data[key]
}

type cacheType int

const (
	redis cacheType = iota
	mem
)

// 实现Cache的简单工厂
type CacheFactory struct {}

func (factory *CacheFactory) Create(cacheType cacheType) (Cache, error) {
	if cacheType == redis {
		return &RedisCache{
			data: map[string]string{},
		}, nil
	}
	if cacheType == mem {
		return &MemCache{
			data: map[string]string{},
		}, nil
	}

	return nil, errors.New("error cache type")
}