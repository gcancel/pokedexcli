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