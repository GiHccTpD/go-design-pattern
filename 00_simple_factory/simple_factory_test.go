package _0_simple_factory

import (
	"fmt"
	"testing"
)

func TestCacheFactory_Create(t *testing.T) {
	cacheFactory := &CacheFactory{}
	redis, err := cacheFactory.Create(redis)
	if err != nil {
		t.Error(err)
	}

	redis.Set("r1", "v1")

	fmt.Println(redis.Get("r1"))

	mem, err := cacheFactory.Create(mem)
	if err != nil {
		t.Error(err)
	}

	mem.Set("m1", "m1")

	fmt.Println(mem.Get("m1"))
}
