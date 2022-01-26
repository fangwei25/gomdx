package redis

import (
	"github.com/go-redis/redis/v8"
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
	return
}

// UpdateMax 更新最大值
func (s DataSource) UpdateMax(key, field string, v int64, expireTime time.Duration) (newV int64, err error) {
	return
}

// UpdateMinus 更新最小值
func (s DataSource) UpdateMinus(key, field string, v int64, expireTime time.Duration) (newV int64, err error) {
	return
}

// QueryOne 查询指定一个数据
func (s DataSource) QueryOne(key, field string) (v int64, err error) {
	return
}

// QueryByList 根据指定域查询多个数据
func (s DataSource) QueryByList(key string, fields ...string) (values map[string]int64, err error) {
	values = make(map[string]int64)
	return
}

// QueryAll 查询指定key的所有域数据
func (s DataSource) QueryAll(key string) (values map[string]int64, err error) {
	values = make(map[string]int64)
	return
}
