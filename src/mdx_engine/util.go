package mdx_engine

import (
	"fmt"
	"github.com/fangwei25/gomdx/src/mdx_cfg"
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

var TFPMap map[mdx_cfg.TimeDimension]string

func init() {
	TFPMap = make(map[mdx_cfg.TimeDimension]string)
	//TFPMap[mdx_cfg.TDEver] = timeFormatEver
	TFPMap[mdx_cfg.TDYear] = timeFormatPattonYear
	TFPMap[mdx_cfg.TDMonth] = timeFormatPattonMonth
	TFPMap[mdx_cfg.TDDay] = timeFormatPattonDay
	TFPMap[mdx_cfg.TDHour] = timeFormatPattonHour
	TFPMap[mdx_cfg.TDMinute] = timeFormatPattonMinute
}

func (e *Engine) GenKey(ownerId int32, eventType string, timeDimension mdx_cfg.TimeDimension, t time.Time) string {
	var key string
	if timeDimension == mdx_cfg.TDEver {
		key = fmt.Sprintf(KeyPatton, e.Cfg.KeyPrefix, ownerId, eventType, timeFormatEver)
	} else {
		tfp := TFPMap[timeDimension]
		key = fmt.Sprintf(KeyPatton, e.Cfg.KeyPrefix, ownerId, eventType, t.Format(tfp))
	}
	return key
}

func (e *Engine) GenField(field string, calcType mdx_cfg.CalcType) string {
	return field + "-" + string(calcType)
}

func (e *Engine) GetLifeTime(timeDimension mdx_cfg.TimeDimension, cfgValue int32) time.Duration {
	switch timeDimension {
	case mdx_cfg.TDEver:
		return -1
	case mdx_cfg.TDYear:
		return time.Duration(cfgValue) * time.Hour * 24 * 365
	case mdx_cfg.TDMonth:
		return time.Duration(cfgValue) * time.Hour * 24 * 30
	case mdx_cfg.TDDay:
		return time.Duration(cfgValue) * time.Hour * 24
	case mdx_cfg.TDHour:
		return time.Duration(cfgValue) * time.Hour
	case mdx_cfg.TDMinute:
		return time.Duration(cfgValue) * time.Minute
	}
	fmt.Printf("GetLifeTime failed, no timeDimension hit: %d", timeDimension)
	return -1
}
