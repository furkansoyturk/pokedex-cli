package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheStruct_Add_Get(t *testing.T) {
	const timeTolive = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "ex.com",
			val: []byte("testdata1"),
		},
		{
			key: "test.com",
			val: []byte("testdata2"),
		},
	}
	
	for i,c := range cases{
		t.Run(fmt.Sprintf("Test case %v",i), func(t *testing.T) {
			cache := NewCache(timeTolive)
			cache.Add(c.key,c.val)
			
			val,ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val){
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestCacheStruct_CleanExpired(t *testing.T) {
	const timeTolive = 5 * time.Second
	const expiredTTL = 2 * timeTolive
	
	cache := NewCache(timeTolive)
	cache.Add("mockkey",[]byte("mockdata"))
	
	_,ok := cache.Get("mockkey")
	if !ok {
		t.Errorf("expected to find key")
		return 
	}
	time.Sleep(expiredTTL)
	_,ok = cache.Get("mockkey")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
