package gomdx

import (
	"time"
)

type ResultData struct {
	Count int64 `json:"count,omitempty"`
	Value int64 `json:"value,omitempty"`
	Max   int64 `json:"max,omitempty"`
	Min   int64 `json:"min,omitempty"`
}

// QueryOne 查询一个数据
func (e Engine) QueryOne(ownerId int32, eventType, subType string, timeDimension TimeDimension, t time.Time) (res *ResultData, err error) {
	key := e.GenKey(ownerId, eventType, timeDimension, t)
	fieldCount := e.GenField(subType, CTCount)
	fieldValue := e.GenField(subType, CTValue)
	fieldMax := e.GenField(subType, CTMax)
	fieldMin := e.GenField(subType, CTMin)
	results, err2 := e.DS.QueryByList(key, fieldCount, fieldValue, fieldMax, fieldMin)
	if nil != err2 {
		err = err2
		return
	}

	res = &ResultData{
		Count: results[fieldCount],
		Value: results[fieldValue],
		Max:   results[fieldMax],
		Min:   results[fieldMin],
	}
	return
}
