package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCacheSaveAndGetThings(t *testing.T) {
	type Object struct {
		Str string
		Num int
	}
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}
	var obj2 Object
	if err := Cache.SaveForEver(context.Background(), "key", obj); err != nil {
		logrus.Error("1", err)
	}
	if _, err := Cache.Get(context.Background(), "key", &obj2); err != nil {
		logrus.Error("2", err)
	}

	if err := Cache.Save(context.Background(), -1, "key", "value"); err != nil {
		logrus.Error("3", err)
	}
	data, err := Cache.Get(context.Background(), "key", nil)
	if err != nil {
		logrus.Error("4", err)
	}

	if err := Cache.Save(context.Background(), -1, "key", true); err != nil {
		logrus.Error("5", err)
	}
	data2, err := Cache.Get(context.Background(), "key", nil)
	if err != nil {
		logrus.Error("6", err)
	}

	defer Cache.Close()
	assert.True(t, data == "value")
	assert.True(t, data2 == true)
	assert.True(t, obj2.Num == 42)

}

func TestRedisPipeline(t *testing.T) {
	s := Cache.InitRedisClientInstance()
	for i := 0; i < 10; i++ {
		s.Set(context.Background(), "key"+strconv.Itoa(i), "hoge"+strconv.Itoa(i), time.Hour)
	}
	client := s

	// 普通にループ
	result := map[string]string{}
	for i := 0; i < 10; i++ {
		key := "key" + strconv.Itoa(i)
		res, _ := client.Get(context.Background(), key).Result()
		result[key] = res
	}

	// Pipelineを使ってループ
	m := map[string]*redis.StringCmd{}
	pipe := client.Pipeline()
	for i := 0; i < 10; i++ {
		m["key"+strconv.Itoa(i)] = pipe.Get(context.Background(), "key"+strconv.Itoa(i))
	}
	_, err := pipe.Exec(context.Background())
	if err != nil {
		panic(err)
	}

	result2 := map[string]string{}
	for k, v := range m {
		res, _ := v.Result()
		result2[k] = res
	}
	logrus.Info(result)
	logrus.Info(result2)
}

func TestRedisWatch(t *testing.T) {

	key := "pipe1"

	client := Cache.InitRedisClientInstance()
	ctx := context.Background()
	txf := func(tx *redis.Tx) error {
		// Phase 1:
		var getPipe *redis.StringCmd

		cmds, err := client.Pipelined(context.Background(), func(pipe redis.Pipeliner) error {
			getPipe = pipe.Get(context.Background(), "getPipe")
			pipe.Set(context.Background(), "pipe1", "p1", -1)
			return nil
		})
		fmt.Println(getPipe)
		fmt.Println(cmds)
		val, _ := getPipe.Result()
		fmt.Println("Value read for 'getPipe':", val)

		// Phase 2: Prepare new data based on read data

		// Phase 3
		_, err = tx.Pipelined(context.Background(), func(pipe redis.Pipeliner) error {
			// pipe handles the error case
			pipe.Set(context.Background(), key, "value22", -1)
			return nil
		})
		return err
	}

	err := client.Watch(context.Background(), txf, key)
	fmt.Println(client.Get(context.Background(), key), err)

	pipe := client.Pipeline()
	//incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)

	_, err = pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
}
func TestRedisPublish(t *testing.T) {

	rdb := Cache.InitRedisClientInstance()
	ctx := context.Background()
	go func() {
		pubsub := rdb.Subscribe(ctx, "mychannel1")
		ch := pubsub.Channel()
		for msg := range ch {
			fmt.Println(msg.Channel, msg.Payload)
		}
	}()

	go func() {
		err := rdb.Publish(ctx, "mychannel1", "payload").Err()
		if err != nil {
			panic(err)
		}
	}()
}
