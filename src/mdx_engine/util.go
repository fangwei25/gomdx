package mdx_engine

import (
	"fmt"
	"github.com/fangwei25/gomdx/src/mdx_cfg"
	"time"
)

const KeyPatton = "mdx-%s-%d-%s-%s" //mdx-key前缀-ownerId-eventType-time(根据time的格式可以识别时间维度)
const totalCounterField = "total"

const (
	timeFormatPattonEver   = "0"
	timeFormatPattonYear   = "2006"
	timeFormatPattonMonth  = "2006-01"
	timeFormatPattonDay    = "2006-01-02"
	timeFormatPattonHour   = "2006-01-02-15"
	timeFormatPattonMinute = "2006-01-02-15-04"
)

var TFPMap map[mdx_cfg.TimeType]string

func init() {
	TFPMap = make(map[mdx_cfg.TimeType]string)
	TFPMap[mdx_cfg.Ever] = timeFormatPattonEver
	TFPMap[mdx_cfg.Year] = timeFormatPattonYear
	TFPMap[mdx_cfg.Month] = timeFormatPattonMonth
	TFPMap[mdx_cfg.Day] = timeFormatPattonDay
	TFPMap[mdx_cfg.Hour] = timeFormatPattonHour
	TFPMap[mdx_cfg.Minute] = timeFormatPattonMinute
}

func (e *Engine) GenKey(ownerId int32, eventType string, timeType mdx_cfg.TimeType, t time.Time) string {
	tfp := TFPMap[timeType]
	key := fmt.Sprintf(e.Cfg.KeyPrefix, KeyPatton, ownerId, eventType, t.Format(tfp))
	return key
}

func (e *Engine) GetLifeTime(timeType mdx_cfg.TimeType, cfgValue int32) time.Duration {
	switch timeType {
	case mdx_cfg.Ever:
		return -1
	case mdx_cfg.Year:
		return time.Duration(cfgValue) * time.Hour * 24 * 365
	case mdx_cfg.Month:
		return time.Duration(cfgValue) * time.Hour * 24 * 30
	case mdx_cfg.Day:
		return time.Duration(cfgValue) * time.Hour * 24
	case mdx_cfg.Hour:
		return time.Duration(cfgValue) * time.Hour
	case mdx_cfg.Minute:
		return time.Duration(cfgValue) * time.Minute
	}
	fmt.Printf("GetLifeTime failed, no timeType hit: %d", timeType)
	return -1
}
