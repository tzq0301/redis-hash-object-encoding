package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{})

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	hashKey := fmt.Sprintf("hash-%d\n", time.Now().UnixMilli())

	// ==================================================================

	if err := rdb.HMSet(ctx, hashKey, "hello", "world").Err(); err != nil {
		panic(err)
	}

	result, err := rdb.ObjectEncoding(ctx, hashKey).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("object encoding of hash object with 1 field: %s\n", result)

	if err := rdb.Del(ctx, hashKey).Err(); err != nil {
		panic(err)
	}

	// ==================================================================

	if err := rdb.HMSet(ctx, hashKey, "bigfieldbigfieldbigfieldbigfieldbigfieldbigfieldbigfieldbigfield", "world").Err(); err != nil {
		panic(err)
	}

	result, err = rdb.ObjectEncoding(ctx, hashKey).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("object encoding of hash object with 64-byte field: %s\n", result)

	if err := rdb.Del(ctx, hashKey).Err(); err != nil {
		panic(err)
	}

	// ==================================================================

	if err := rdb.HMSet(ctx, hashKey, "bigfieldbigfieldbigfieldbigfieldbigfieldbigfieldbigfieldbigfieldd", "world").Err(); err != nil {
		panic(err)
	}

	result, err = rdb.ObjectEncoding(ctx, hashKey).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("object encoding of hash object with 65-byte field: %s\n", result)

	if err := rdb.Del(ctx, hashKey).Err(); err != nil {
		panic(err)
	}

	// ==================================================================

	if err := rdb.HMSet(ctx, hashKey, "field", "bigvaluebigvaluebigvaluebigvaluebigvaluebigvaluebigvaluebigvalue").Err(); err != nil {
		panic(err)
	}

	result, err = rdb.ObjectEncoding(ctx, hashKey).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("object encoding of hash object with 64-byte value: %s\n", result)

	if err := rdb.Del(ctx, hashKey).Err(); err != nil {
		panic(err)
	}

	// ==================================================================

	if err := rdb.HMSet(ctx, hashKey, "field", "bigvaluebigvaluebigvaluebigvaluebigvaluebigvaluebigvaluebigvaluee").Err(); err != nil {
		panic(err)
	}

	result, err = rdb.ObjectEncoding(ctx, hashKey).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("object encoding of hash object with 65-byte value: %s\n", result)

	if err := rdb.Del(ctx, hashKey).Err(); err != nil {
		panic(err)
	}

	// ==================================================================

	var fvPairs []any

	for i := 0; i < 512; i++ {
		fvPairs = append(fvPairs, "field", "value")
	}

	if err := rdb.HMSet(ctx, hashKey, fvPairs...).Err(); err != nil {
		panic(err)
	}

	result, err = rdb.ObjectEncoding(ctx, hashKey).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("object encoding of hash object with 512 field: %s\n", result)

	if err := rdb.Del(ctx, hashKey).Err(); err != nil {
		panic(err)
	}

	// ==================================================================

	fvPairs = append(fvPairs, "field", "value")

	if err := rdb.HMSet(ctx, hashKey, fvPairs...).Err(); err != nil {
		panic(err)
	}

	result, err = rdb.ObjectEncoding(ctx, hashKey).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("object encoding of hash object with 513 field: %s\n", result)

	if err := rdb.Del(ctx, hashKey).Err(); err != nil {
		panic(err)
	}
}
