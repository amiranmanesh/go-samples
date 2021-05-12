package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"time"
)

type iCache interface {
	InitRedisClientInstance() *redis.Client

	Save(ctx context.Context, ttl time.Duration, key string, value interface{}) error
	SaveForEver(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string, value interface{}) (interface{}, error)

	LPush(ctx context.Context, key string, value interface{}) error
	LPop(ctx context.Context, key string, value interface{}) (interface{}, error)
	LRPop(ctx context.Context, key string, value interface{}) (interface{}, error)
	LIndex(ctx context.Context, key string, index int64, value interface{}) (interface{}, error)
	LLen(ctx context.Context, key string) int64

	Close()
}

type cache struct{}

var Cache iCache = &cache{}

var (
	redisClient *redis.Client
)

func (cache) InitRedisClientInstance() *redis.Client {
	return initRedisClientInstance()
}

func (cache) Save(ctx context.Context, ttl time.Duration, key string, value interface{}) error {
	return save(ctx, ttl, key, value)
}
func (cache) SaveForEver(ctx context.Context, key string, value interface{}) error {
	return save(ctx, -1, key, value)
}

func (c cache) LPush(ctx context.Context, key string, value interface{}) error {
	nilChecking()
	data, err := json.Marshal(value)
	if err != nil {
		return errors.Trace(err)
	}
	if err := redisClient.LPush(ctx, key, data).Err(); err != nil {
		logrus.Error(errors.Trace(err))
		return err
	}
	return nil
}

func save(ctx context.Context, ttl time.Duration, key string, value interface{}) error {
	nilChecking()
	data, err := json.Marshal(value)
	if err != nil {
		return errors.Trace(err)
	}
	if err := redisClient.Set(ctx, key, data, ttl).Err(); err != nil {
		return errors.Trace(err)
	}
	return nil
}

/**
for objects should be use like this
cache.GetObject(context.Background(), "key", &obj)
for others could be use like this
cache.GetObject(context.Background(), "key", nil)
*/
func (cache) Get(ctx context.Context, key string, value interface{}) (interface{}, error) {
	nilChecking()
	val, err := redisClient.Get(ctx, key).Result()
	return gettingProcess(val, err, value)
}

func (c cache) LPop(ctx context.Context, key string, value interface{}) (interface{}, error) {
	nilChecking()
	val, err := redisClient.LPop(ctx, key).Result()
	return gettingProcess(val, err, value)
}

func (c cache) LRPop(ctx context.Context, key string, value interface{}) (interface{}, error) {
	nilChecking()
	val, err := redisClient.RPop(ctx, key).Result()
	return gettingProcess(val, err, value)
}

func (c cache) LIndex(ctx context.Context, key string, index int64, value interface{}) (interface{}, error) {
	nilChecking()
	val, err := redisClient.LIndex(ctx, key, index).Result()
	return gettingProcess(val, err, value)
}

func gettingProcess(val string, err error, value interface{}) (interface{}, error) {
	if err != nil {
		return nil, errors.Trace(err)
	}
	if err2 := json.Unmarshal([]byte(val), &value); err2 != nil {
		return nil, errors.Trace(err2)
	}
	return value, nil
}

func (c cache) LLen(ctx context.Context, key string) int64 {
	nilChecking()
	length, err := redisClient.LLen(ctx, key).Result()
	if err != nil {
		return -1
	}
	return length
}

func (cache) Close() {
	redisClient.Close()
}

func initRedisClientInstance() *redis.Client {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}
	return redisClient
}

func nilChecking() {
	if redisClient == nil {
		initRedisClientInstance()
	}
}
