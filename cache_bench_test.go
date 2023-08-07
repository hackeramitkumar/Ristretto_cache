package main

// import (
// 	"errors"
// 	"fmt"
// 	"math/rand"
// 	"time"

// 	"github.com/dgraph-io/ristretto"
// )

// type Cache interface {
// 	Get(key string, rule string, image string) (string, bool)
// 	Set(key string, rule string, image string) (bool, error)
// }

// type val struct {
// 	m map[string][]string
// }

// const (
// 	maxKeyLength = 128
// 	workloadSize = 2 << 20
// 	charset      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 	numset       = "0123456789"
// )

// var (
// 	errKeyNotFound  = errors.New("key not found")
// 	errInvalidValue = errors.New("invalid value")
// )

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// func generateRandomString(length int) string {
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[rand.Intn(len(charset))]
// 	}
// 	return string(b)
// }

// func generateRandomId(length int) string {
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = numset[rand.Intn(len(numset))]
// 	}
// 	return string(b)
// }

// func genratekey3() []string {
// 	keys := make([]string, workloadSize)
// 	for i := 0; i < workloadSize; i++ {
// 		policy := generateRandomId(10)
// 		img := generateRandomString(30)
// 		rule := generateRandomString(10)
// 		keys[i] = (policy + ";" + img + ";" + rule)
// 	}
// 	return keys
// }

// func genratekey2() []string {
// 	keys := make([]string, workloadSize)
// 	for i := 0; i < workloadSize; i++ {
// 		policy := generateRandomId(10)
// 		img := generateRandomString(30)
// 		keys[i] = (policy + ";" + img)
// 	}
// 	return keys
// }

// // Ristretto with TTL

// type RistrettoCacheTTL struct {
// 	c          *ristretto.Cache
// 	defaultTTL time.Duration
// }

// func (r *RistrettoCacheTTL) Get(policy string, rule string, image string) (interface{}, bool) {
// 	key := policy + ";" + rule + ";" + image
// 	val, g := r.c.Get(key)
// 	return val, g
// }

// func (r *RistrettoCacheTTL) Set(policy string, rule string, image string) (bool, error) {
// 	key := policy + ";" + rule + ";" + image
// 	if r.defaultTTL <= 0 {
// 		return r.c.Set(key, "val", 1), nil
// 	}
// 	return r.c.SetWithTTL(key, "vall", 1, r.defaultTTL), nil
// }

// func newRistrettoTTL(keysInWindow int, ttl time.Duration) *RistrettoCacheTTL {
// 	cache, err := ristretto.NewCache(&ristretto.Config{
// 		NumCounters: int64(keysInWindow * 10),
// 		MaxCost:     int64(keysInWindow),
// 		BufferItems: 64,
// 		OnEvict: func(item *ristretto.Item) {
// 			fmt.Println("Amit is key ko hatay gaya  : ", item.Key)
// 		},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	return &RistrettoCacheTTL{cache, ttl}
// }

// // // Default Ristretto without

// // type RistrettoCache struct {
// // 	c *ristretto.Cache
// // }

// // func (r *RistrettoCache) Get(policy string,rule string, image string) (string, error) {
// // 	key := policy+";"+rule+";"+image
// // 	return r.c.Get()
// // }

// // func (r *RistrettoCache) Set(key string, value int64) error {
// // 	_ = r.c.Set(key, value, 1)
// // 	return nil
// // }

// // func newRistretto(keysInWindow int) *RistrettoCache {
// // 	cache, err := ristretto.NewCache(&ristretto.Config{
// // 		NumCounters: int64(keysInWindow * 10),
// // 		MaxCost:     int64(keysInWindow),
// // 		BufferItems: 64,
// // 	})
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	return &RistrettoCache{cache}
// // }

// /// ristretto with TTL

// // Benchmarking

// // func runCacheBenchmark(b *testing.B, cache Cache, keys []string, pctWrites uint64) {
// // 	b.ReportAllocs()

// // 	size := len(keys)
// // 	mask := size - 1
// // 	rc := uint64(0)

// // 	// initialize cache
// // 	for i := 0; i < size; i++ {
// // 		_ = cache.Set("img",)
// // 	}

// // 	b.ResetTimer()
// // 	b.RunParallel(func(pb *testing.PB) {
// // 		index := rand.Int() & mask
// // 		mc := atomic.AddUint64(&rc, 1)

// // 		if pctWrites*mc/100 != pctWrites*(mc-1)/100 {
// // 			for pb.Next() {
// // 				_ = cache.Set(keys[index&mask], 545421454)
// // 				index = index + 1
// // 			}
// // 		} else {
// // 			for pb.Next() {
// // 				_, _ = cache.Get(keys[index&mask])
// // 				index = index + 1
// // 			}
// // 		}
// // 	})
// // }

// // func BenchmarkCaches(b *testing.B) {
// // 	string_3_key := genratekey3()
// // 	time1 := 1e2 * time.Millisecond
// // 	time2 := 1e5 * time.Millisecond

// // 	benchmarks := []struct {
// // 		name      string
// // 		cache     Cache
// // 		keys      []string
// // 		pctWrites uint64
// // 	}{
// // 		{"RistrettoWithoutTTLRead", newRistretto(b.N), string_3_key, 0},
// // 		{"RistrettoWithSmallTTLRead", newRistrettoTTL(b.N, time1), string_3_key, 0},
// // 		{"RistrettoWithLargeTTLRead", newRistrettoTTL(b.N, time2), string_3_key, 0},
// // 		{"RistrettoWithoutTTLWrite", newRistretto(b.N), string_3_key, 100},
// // 		{"RistrettoWithSmallTTLWrite", newRistrettoTTL(b.N, time1), string_3_key, 100},
// // 		{"RistrettoWithLargeTTLWrite", newRistrettoTTL(b.N, time2), string_3_key, 100},
// // 		{"RistrettoWithoutTTLMixed", newRistretto(b.N), string_3_key, 30},
// // 		{"RistrettoWithSmallTTLMixed", newRistrettoTTL(b.N, time1), string_3_key, 30},
// // 		{"RistrettoWithLargeTTLMixed", newRistrettoTTL(b.N, time2), string_3_key, 30},
// // 	}

// // 	for _, bm := range benchmarks {
// // 		b.Run(bm.name, func(b *testing.B) {
// // 			runCacheBenchmark(b, bm.cache, bm.keys, bm.pctWrites)
// // 		})
// // 	}
// // }

// func main() {
// 	ca := newRistrettoTTL(5, 30*time.Second)
// 	v1, e := ca.Set("p1", "r1", "i1")
// 	fmt.Println(v1)
// 	v2, e := ca.Set("p2", "r2", "i2")
// 	fmt.Println(v2)

// 	v3, e := ca.Set("p3", "r3", "i3")
// 	fmt.Println(v3)

// 	v4, e := ca.Set("p4", "r4", "i4")
// 	fmt.Println(v4)

// 	v5, e := ca.Set("p5", "r5", "i5")
// 	fmt.Println(v5)

// 	v6, e := ca.Set("p6", "r6", "i6")
// 	fmt.Println(v6)

// 	v7, e := ca.Set("p7", "r7", "i7")
// 	fmt.Println(v7)

// 	fmt.Println(e)

// 	s1, b1 := ca.Get("p1", "r1", "i1")
// 	fmt.Println(s1)
// 	s2, b1 := ca.Get("p2", "r2", "i2")
// 	fmt.Println(s2)

// 	s3, b1 := ca.Get("p3", "r3", "i3")
// 	fmt.Println(s3)

// 	s4, b1 := ca.Get("p4", "r4", "i4")
// 	fmt.Println(s4)

// 	s5, b1 := ca.Get("p5", "r5", "i5")
// 	fmt.Println(s5)

// 	s6, b1 := ca.Get("p6", "r6", "i6")
// 	fmt.Println(s6)

// 	s7, b1 := ca.Get("p7", "r7", "i7")
// 	fmt.Println(s7)

// 	fmt.Println(b1)

// }
