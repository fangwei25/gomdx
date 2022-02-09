package engine

import (
	"fmt"
	"github.com/fangwei25/gomdx/cfg"
	"time"
)

const KeyPatton = "mdx-%s-%d-%s-%s" //mdx-key前缀-ownerId-eventType-time(根据time的格式可以识别时间维度)
const totalCounterField = "total"

const (
	timeFormatEver         = "0"
	timeFormatPattonYear   = "2006"
	timeFormatPattonMonth  = "2006-01"
	timeFormatPattonDay    = "2006-01-02"
	timeFormatPattonHour   = "2006-01-02-15"
	timeFormatPattonMinute = "2006-01-02-15-04"
)

var TFPMap map[cfg.TimeDimension]string

func init() {
	TFPMap = make(map[cfg.TimeDimension]string)
	//TFPMap[cfg.TDEver] = timeFormatEver
	TFPMap[cfg.TDYear] = timeFormatPattonYear
	TFPMap[cfg.TDMonth] = timeFormatPattonMonth
	TFPMap[cfg.TDDay] = timeFormatPattonDay
	TFPMap[cfg.TDHour] = timeFormatPattonHour
	TFPMap[cfg.TDMinute] = timeFormatPattonMinute
}

func (e *Engine) GenKey(ownerId int32, eventType string, timeDimension cfg.TimeDimension, t time.Time) string {
	var key string
	if timeDimension == cfg.TDEver {
		key = fmt.Sprintf(KeyPatton, e.Cfg.KeyPrefix, ownerId, eventType, timeFormatEver)
	} else {
		tfp := TFPMap[timeDimension]
		key = fmt.Sprintf(KeyPatton, e.Cfg.KeyPrefix, ownerId, eventType, t.Format(tfp))
	}
	return key
}

func (e *Engine) GenField(field string, calcType cfg.CalcType) string {
	return field + "-" + string(calcType)
}

func (e *Engine) GetLifeTime(timeDimension cfg.TimeDimension, cfgValue int32) time.Duration {
	switch timeDimension {
	case cfg.TDEver:
		return -1
	case cfg.TDYear:
		return time.Duration(cfgValue) * time.Hour * 24 * 365
	case cfg.TDMonth:
		return time.Duration(cfgValue) * time.Hour * 24 * 30
	case cfg.TDDay:
		return time.Duration(cfgValue) * time.Hour * 24
	case cfg.TDHour:
		return time.Duration(cfgValue) * time.Hour
	case cfg.TDMinute:
		return time.Duration(cfgValue) * time.Minute
	}
	fmt.Printf("GetLifeTime failed, no timeDimension hit: %d", timeDimension)
	return -1
}
