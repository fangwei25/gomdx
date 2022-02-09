package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fangwei25/gomdx/utils"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var ctx = context.Background()

func createClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     utils.GlobalObject.RedisAddr,
		Password: utils.GlobalObject.RedisPWD,        // no password set
		DB:       int(utils.GlobalObject.RedisDBIdx), // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if nil != err {
		panic(fmt.Sprintf("connect redis failed, err: %v", err))
	}

	return rdb
}

// Set a key/value
func (s DataSource) Set(key string, data interface{}, expireSeconds time.Duration) error {

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = s.RedisConn.Set(ctx, key, value, expireSeconds).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s DataSource) HGetAll(key string) (map[string]string, error) {
	return s.RedisConn.HGetAll(ctx, key).Result()
}

func (s DataSource) HDel(key string, fields ...string) (int64, error) {
	return s.RedisConn.HDel(ctx, key, fields...).Result()
}

func (s DataSource) HGet(key string, field string) (string, error) {
	return s.RedisConn.HGet(ctx, key, field).Result()
}

func (s DataSource) HExist(key string, field string) (bool, error) {
	return s.RedisConn.HExists(ctx, key, field).Result()
}

func (s DataSource) HGetInt64(key string, field string) (int64, error) {
	result, err := s.RedisConn.HGet(ctx, key, field).Result()
	if nil != err {
		if err == redis.Nil {
			return 0, nil
		} else {
			return 0, err //查询失败，返回0，不限制生成，避免业务不能持续
		}
	}
	if result == "" {
		//一般查询的结果为空字符串,减少了parse的开销
		return 0, nil
	}
	totalNum, err := strconv.ParseInt(result, 10, 32)
	return totalNum, err
}
func (s DataSource) HGetInt32(key string, field string) (int32, error) {
	totalNum, err := s.HGetInt64(key, field)
	return int32(totalNum), err
}
func (s DataSource) HMGet(key string, fields ...string) ([]interface{}, error) {
	return s.RedisConn.HMGet(ctx, key, fields...).Result()
}
func (s DataSource) HMGetInt64(key string, fields ...string) (map[string]int64, error) {
	res, err := s.RedisConn.HMGet(ctx, key, fields...).Result()
	if nil != err {
		return nil, err
	}
	resMap := make(map[string]int64)
	for idx, oneRes := range res {
		if s, ok := oneRes.(string); ok {
			num, err := strconv.ParseInt(s, 10, 64)
			if err == nil {
				resMap[fields[idx]] = num
			}
		}
	}
	return resMap, nil
}

func (s DataSource) HSet(key string, field string, val interface{}) (int64, error) {
	return s.RedisConn.HSet(ctx, key, field, val).Result()
}
func (s DataSource) HSetNx(key string, field string, val interface{}) (bool, error) {
	return s.RedisConn.HSetNX(ctx, key, field, val).Result()
}

func (s DataSource) HMSet(key string, kvs map[string]interface{}) (bool, error) {
	return s.RedisConn.HMSet(ctx, key, kvs).Result()
}

func (s DataSource) HIncrBy(key string, field string, incValue int64) (int64, error) {
	return s.RedisConn.HIncrBy(ctx, key, field, incValue).Result()
}

// IncrBy 增加
func (s DataSource) IncrBy(key string, incValue int64) (int64, error) {
	return s.RedisConn.IncrBy(ctx, key, incValue).Result()
}

// Get a key
func (s DataSource) Get(key string) (string, error) {
	return s.RedisConn.Get(ctx, key).Result()
}

func (s DataSource) GetInt64(key string) (int64, error) {
	result, err := s.RedisConn.Get(ctx, key).Result()
	if nil != err {
		if err == redis.Nil {
			return 0, nil
		} else {
			return 0, err //查询失败，返回0，不限制生成，避免业务不能持续
		}
	}
	if result == "" {
		//一般查询的结果为空字符串,减少了parse的开销
		return 0, nil
	}
	totalNum, err := strconv.ParseInt(result, 10, 64)
	return totalNum, err
}

func (s DataSource) LPush(key string, values ...interface{}) (int64, error) {
	return s.RedisConn.LPush(ctx, key, values).Result()
}

func (s DataSource) LRange(key string, start, stop int64) ([]string, error) {
	return s.RedisConn.LRange(ctx, key, start, stop).Result()
}

// Delete delete a kye
func (s DataSource) Delete(key string) (int64, error) {
	return s.RedisConn.Del(ctx, key).Result()
}

func (s DataSource) Expire(key string, seconds time.Duration) (bool, error) {
	return s.RedisConn.Expire(ctx, key, seconds).Result()
}

func (s DataSource) Exists(key string) bool {
	i, e := s.RedisConn.Exists(ctx, key).Result()
	if nil != e {
		return false
	}
	return 1 == i
}

func (s DataSource) TTL(key string) time.Duration {
	t, err := s.RedisConn.TTL(ctx, key).Result()
	if nil != err {
		return 0 * time.Second
	}
	return t
}
