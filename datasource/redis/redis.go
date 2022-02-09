package redis

import (
	"context"
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

func CreateDataSource() *DataSource {
	s := &DataSource{
		RedisConn: createClient(),
	}
	return s
}

type DataSource struct {
	RedisConn *redis.Client
}

// Increase 增加
func (s DataSource) Increase(key, field string, v int64) (newV int64, err error) {
	return s.HIncrBy(key, field, v)
}

// UpdateMax 更新最大值
func (s DataSource) UpdateMax(key, field string, v int64) (newV int64, err error) {
	oldV := v
	oldV, err = s.HGetInt64(key, field)
	if nil != err {
		return
	}
	if oldV >= v {
		return
	}
	newV, err = s.HSet(key, field, v)
	return
}

// UpdateMinus 更新最小值
func (s DataSource) UpdateMinus(key, field string, v int64) (newV int64, err error) {
	oldV := v
	oldV, err = s.HGetInt64(key, field)
	if nil != err {
		return
	}
	if oldV <= v {
		return
	}
	newV, err = s.HSet(key, field, v)
	return
}

// QueryOne 查询指定一个数据
func (s DataSource) QueryOne(key, field string) (v int64, err error) {
	v, err = s.HGetInt64(key, field)
	return
}

// QueryByList 根据指定域查询多个数据
func (s DataSource) QueryByList(key string, fields ...string) (values map[string]int64, err error) {
	values, err = s.HMGetInt64(key, fields...)
	return
}

// QueryAll 查询指定key的所有域数据
func (s DataSource) QueryAll(key string) (values map[string]int64, err error) {
	values = make(map[string]int64)
	var results map[string]string
	results, err = s.HGetAll(key)
	if nil != err {
		return
	}
	for k, v := range results {
		num, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			values[k] = num
		}
	}
	return
}

// HGetAll 获取key里所有域&值
func (s DataSource) HGetAll(key string) (map[string]string, error) {
	return s.RedisConn.HGetAll(ctx, key).Result()
}

// HGetInt64 获取key里指定域&值 并尝试将值转为int64
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

// HMGetInt64 获取key里指定一系列域&值 并尝试将值转为int64
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

// HSet 设置key里指定的域&值
func (s DataSource) HSet(key string, field string, val interface{}) (int64, error) {
	return s.RedisConn.HSet(ctx, key, field, val).Result()
}

// HIncrBy 增加
func (s DataSource) HIncrBy(key string, field string, incValue int64) (int64, error) {
	return s.RedisConn.HIncrBy(ctx, key, field, incValue).Result()
}

// Expire 设置key的超时
func (s DataSource) Expire(key string, seconds time.Duration) (bool, error) {
	return s.RedisConn.Expire(ctx, key, seconds).Result()
}
