package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		fixtures := []struct {
			key     Key
			value   interface{}
			survive bool
		}{
			{key: "aaa", value: 1, survive: false},
			{key: "bbb", value: 3, survive: false},
			{key: "ccc", value: 5, survive: true},
			{key: "ddd", value: 7, survive: true},
			{key: "eee", value: 9, survive: true},
		}
		for _, f := range fixtures {
			c.Set(f.key, f.value)
		}
		for _, f := range fixtures {
			val, ok := c.Get(f.key)
			if f.survive {
				require.True(t, ok)
				require.Equal(t, f.value, val)
			} else {
				require.False(t, ok)
				require.Nil(t, val)
			}
		}
	})

	t.Run("overwriting logic", func(t *testing.T) {
		var k Key
		c := NewCache(3)
		k = "aaa"
		fixtures := []struct {
			key     Key
			value   interface{}
			survive bool
		}{
			{key: k, value: 1, survive: true},
			{key: "bbb", value: 3, survive: false},
			{key: "ccc", value: 5, survive: false},
			{key: "ddd", value: 7, survive: true},
			{key: "eee", value: 9, survive: true},
		}
		for _, f := range fixtures {
			c.Set(f.key, f.value)
			c.Get(k)
		}
		for _, f := range fixtures {
			val, ok := c.Get(f.key)
			if f.survive {
				require.True(t, ok)
				require.Equal(t, f.value, val)
			} else {
				require.False(t, ok)
				require.Nil(t, val)
			}
		}
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
