package iface

import "time"

// Updater 更新
type Updater interface {
	// Increase 累加值（通过field和v可实现累加数值和累加计数两种作用）
	Increase(key, field string, v int64, expireTime time.Duration) (newV int64, err error)
	// UpdateMax 更新最大值
	UpdateMax(key, field string, v int64, expireTime time.Duration) (newV int64, err error)
	// UpdateMinus 更新最小值
	UpdateMinus(key, field string, v int64, expireTime time.Duration) (newV int64, err error)
}

// Querier 查询
type Querier interface {
	// QueryOne 查询指定一个数据
	QueryOne(key, field string) (v int64, err error)
	// QueryByList 根据指定域查询多个数据
	QueryByList(key string, fields ...string) (values map[string]int64, err error)
	// QueryAll 查询指定key的所有域数据
	QueryAll(key string) (values map[string]int64, err error)
}

type DataSourcer interface {
	Updater
	Querier
}
