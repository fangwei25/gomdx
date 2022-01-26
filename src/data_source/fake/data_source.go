package redis

import (
	"fmt"
	"time"
)

func CreateDataSource() *DataSource {
	s := &DataSource{}
	return s
}

type DataSource struct {
	LastValue int64 //上一次更新的数据
}

func (s DataSource) Increase(key, field string, v int64, expireTime time.Duration) (newV int64, err error) {
	fmt.Printf("Increase key=%s, field=%s, v=%d, expire(minutes): %d", key, field, v, expireTime/time.Minute)
	s.LastValue = v
	newV = v
	return
}

// UpdateMax 更新最大值
func (s DataSource) UpdateMax(key, field string, v int64, expireTime time.Duration) (newV int64, err error) {
	fmt.Printf("UpdateMax key=%s, field=%s, v=%d, expire(minutes): %d", key, field, v, expireTime/time.Minute)
	s.LastValue = v
	newV = v
	return
}

// UpdateMinus 更新最小值
func (s DataSource) UpdateMinus(key, field string, v int64, expireTime time.Duration) (newV int64, err error) {
	fmt.Printf("UpdateMinus key=%s, field=%s, v=%d, expire(minutes): %d", key, field, v, expireTime/time.Minute)
	s.LastValue = v
	s.LastValue = v
	newV = v
	return
}

// QueryOne 查询指定一个数据
func (s DataSource) QueryOne(key, field string) (v int64, err error) {
	fmt.Printf("QueryOne key=%s, field=%s, v=%d", key, field, s.LastValue)
	v = s.LastValue
	return
}

// QueryByList 根据指定域查询多个数据
func (s DataSource) QueryByList(key string, fields ...string) (values map[string]int64, err error) {
	values = make(map[string]int64)
	fmt.Printf("QueryOne key=%s, fields=%v, v=%d", key, fields, s.LastValue)
	for _, field := range fields {
		values[field] = s.LastValue
	}
	return
}

// QueryAll 查询指定key的所有域数据
func (s DataSource) QueryAll(key string) (values map[string]int64, err error) {
	values = make(map[string]int64)
	fmt.Printf("QueryOne key=%s, field=field, v=%d", key, s.LastValue)
	values["field"] = s.LastValue
	return
}
