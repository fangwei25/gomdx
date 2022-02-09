package memery

import (
	"errors"
	"time"
)

func CreateDataSource() *DataSource {
	s := &DataSource{
		Cache: make(map[string]FiledData),
	}
	return s
}

type FiledData map[string]int64

type DataSource struct {
	Cache map[string]FiledData
}

func (s DataSource) Expire(key string, expireTime time.Duration) (bool, error) {
	return true, nil
}

func (s DataSource) Increase(key, field string, v int64) (newV int64, err error) {
	fData, ok := s.Cache[key]
	if !ok {
		fData = make(map[string]int64)
		s.Cache[key] = fData
	}

	fData[field] += v
	newV = fData[field]
	return
}

// UpdateMax 更新最大值
func (s DataSource) UpdateMax(key, field string, v int64) (newV int64, err error) {
	fData, ok := s.Cache[key]
	if !ok {
		fData = make(map[string]int64)
		s.Cache[key] = fData
	}

	oldValue, ok := fData[field]
	if !ok || v > oldValue {
		fData[field] = v
	}
	newV = fData[field]
	return
}

// UpdateMinus 更新最小值
func (s DataSource) UpdateMinus(key, field string, v int64) (newV int64, err error) {
	fData, ok := s.Cache[key]
	if !ok {
		fData = make(map[string]int64)
		s.Cache[key] = fData
	}

	oldValue, ok := fData[field]
	if !ok || v < oldValue {
		fData[field] = v
	}
	newV = fData[field]
	return
}

// QueryOne 查询指定一个数据
func (s DataSource) QueryOne(key, field string) (v int64, err error) {
	fData, ok := s.Cache[key]
	if !ok {
		err = errors.New("key not found")
		return
	}
	v, ok = fData[field]
	if !ok {
		err = errors.New("field not found")
		return
	}
	return
}

// QueryByList 根据指定域查询多个数据
func (s DataSource) QueryByList(key string, fields ...string) (values map[string]int64, err error) {
	values = make(map[string]int64)
	for _, field := range fields {
		v, err2 := s.QueryOne(key, field)
		if nil != err2 {
			if err2.Error() == "key not found" {
				err = err2
				return
			}
			continue
		}
		values[field] = v
	}
	return
}

// QueryAll 查询指定key的所有域数据
func (s DataSource) QueryAll(key string) (values map[string]int64, err error) {
	values = make(map[string]int64)
	fData, ok := s.Cache[key]
	if !ok {
		err = errors.New("key not found")
		return
	}

	for k, v := range fData {
		values[k] = v
	}
	return
}
