package redis

import (
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func CreateDataSource() *DataSource {
	s := &DataSource{
		RedisConn: createClient(),
	}
	return s
}

type DataSource struct {
	RedisConn *redis.Client
}

func (s DataSource) Increase(key, field string, v int64, expireTime time.Duration) (newV int64, err error) {
	return s.HIncrBy(key, field, v)
}

// UpdateMax 更新最大值
func (s DataSource) UpdateMax(key, field string, v int64, expireTime time.Duration) (newV int64, err error) {
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
func (s DataSource) UpdateMinus(key, field string, v int64, expireTime time.Duration) (newV int64, err error) {
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
