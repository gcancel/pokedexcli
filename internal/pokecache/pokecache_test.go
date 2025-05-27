package pokecache

import(
	"testing"
)

TestNewCache(t *testing.T){
	cases := struct{
		input time.Duration
		expected *Cache
	}{
		{
			input: (5 * time.Seconds),
			expected: &Cache{CacheEntries: make(map[string]CacheEntry, 0)}
		},
	}
}

for _,case := range cases{
	actual := NewCache(case.input)
	if actual != case.expected{
		t.Errorf("Cache struct of actual does not equal expected... %v, %v", actual, case.expected)
	}
}

TestAddCacheEntry(t *testing.T){
	cases := struct{
		key string
		val []byte 
	}{
		{
			key: "www.google.com/",
			val: byte("testing data")
		},
		{
			key: "www.amazon.com/prime/",
			val: byte("testing amazon.com")
		},
	}
	
	for _,case := range cases{
		cache := NewCache(5 * time.Seconds)
		cache.Add(case.key, case.val)
		val, ok := cache.Get("google.com")
		if !ok{
			fmt.Errorf("expected to find a cache key")
			return
		}
		if string(val) != string(case.val){
			fmt.Errorf("expected to a cache value")
		}  
	}
}
