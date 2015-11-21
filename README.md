bloomis
====

Distributed bloom filters for very large sets using redis for persistence and consensus.

## Goals:

+ Plugable as a module with a minimal, clean interface
+ Transparent for the client using the standard semantics for bloom filters
+ Agents can join and leave the distributed bloom filter anytime without effort
+ Minimal data transference over the network

## Usage:

Import the required libraries

```
import (
	"gopkg.in/redis.v3"
	"github.com/kpacha/bloomis"
)
```

Create your redis client

```
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
```

Calculate the best `m` & `k` values for your use case with one of the currently available online like [thomas hurst's implementation](http://hur.st/bloomfilter?n=100000000&p=1.0E-9) and create a single bloom filter (this example uses `n = 100,000,000` & `p = 1e-9`)

```
	m := 4313276270
	k := 30
	bf := bloomis.NewSingleFilter(m, k, client)
```

Start adding and testing values

```
	fmt.Println(bf.AddToDefault([]byte("foo")))
	fmt.Println(bf.TestToDefault([]byte("foo")))
	fmt.Println(bf.TestToDefault([]byte("baz")))
	// Output:
	// <nil>
	// true <nil>
	// false <nil>
```

Check the [bloom_test.go](https://github.com/kpacha/bloomis/blob/master/bloom_test.go) for more examples

## Customization

Since the `BloomFilter` is just a facade, you can create your custom filters with your own `Hasher` implementations and/or personalize how to store/load the filter metadata, etc.