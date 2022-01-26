package mdx_engine

import (
	"fmt"
	"github.com/fangwei25/gomdx/src/iface"
	"github.com/fangwei25/gomdx/src/mdx_cfg"
	"time"
)

const KeyPatton = "mdx-%d-%s-%s" //event-ownerId-eventType-time(根据time的格式可以识别时间维度)
const totalCounterField = "total"

type Engine struct {
	DS  iface.DataSourcer
	Cfg *mdx_cfg.Cfg
}

func CreateEngine(ds iface.DataSourcer, cfg *mdx_cfg.Cfg) *Engine {
	return &Engine{
		DS:  ds,
		Cfg: cfg,
	}
}

// Update 数值统计 按照给定的数值进行累加统计
func (e *Engine) Update(ownerId int32, eventType string, field string, value int64) {
	//根据配置决定统计的计算维度（累加，计数，记录最大值，记录最小值）和时间维度(总，年，月，日，小时，分)
	eventCfg, ok := e.Cfg.EventCfgs[eventType]
	if !ok {
		//未配置的事件类型，则用默认配置： 时间维度=永久，计算维度=累加值、累加计数
		eventCfg = mdx_cfg.DefaultEventCfg
	}

	for _, timeCfg := range eventCfg.TimeCfgList {
		switch timeCfg.Type {
		case mdx_cfg.Total:
			e.UpdateTotal(ownerId, eventType, field, value, timeCfg)
		case mdx_cfg.Year:
			e.UpdateTotal(ownerId, eventType, field, value, timeCfg)
		case mdx_cfg.Month:
			e.UpdateTotal(ownerId, eventType, field, value, timeCfg)
		case mdx_cfg.Day:
			e.UpdateTotal(ownerId, eventType, field, value, timeCfg)
		case mdx_cfg.Hour:
			e.UpdateTotal(ownerId, eventType, field, value, timeCfg)
		case mdx_cfg.Minute:
			e.UpdateTotal(ownerId, eventType, field, value, timeCfg)
		default:
			fmt.Printf("Engine.Update failed, unsupported time dimension: %d", timeCfg.Type)
		}
	}
}

func (e *Engine) UpdateTotal(ownerId int32, eventType string, field string, value int64, cfg *mdx_cfg.TimeCfg) {
	key := fmt.Sprintf(KeyPatton, ownerId, eventType, totalCounterField)
	lifeDuration := time.Duration(-1)
	e.UpdateByTimeDimension(key, field, value, lifeDuration, cfg.CalcList)
}

func (e *Engine) UpdateYear(ownerId int32, eventType string, field string, value int64, cfg *mdx_cfg.TimeCfg) {
	key := fmt.Sprintf(KeyPatton, ownerId, eventType, time.Now().Format("2006")) //年
	lifeDuration := time.Duration(cfg.LiftTime) * time.Hour * 24 * 365
	e.UpdateByTimeDimension(key, field, value, lifeDuration, cfg.CalcList)

}

func (e *Engine) UpdateMonth(ownerId int32, eventType string, field string, value int64, cfg *mdx_cfg.TimeCfg) {
	key := fmt.Sprintf(KeyPatton, ownerId, eventType, time.Now().Format("2006-01")) //年-月
	lifeDuration := time.Duration(cfg.LiftTime) * time.Hour * 24 * 30
	e.UpdateByTimeDimension(key, field, value, lifeDuration, cfg.CalcList)
}

func (e *Engine) UpdateDay(ownerId int32, eventType string, field string, value int64, cfg *mdx_cfg.TimeCfg) {
	key := fmt.Sprintf(KeyPatton, ownerId, eventType, time.Now().Format("2006-01-02")) //年-月-日
	lifeDuration := time.Duration(cfg.LiftTime) * time.Hour * 24
	e.UpdateByTimeDimension(key, field, value, lifeDuration, cfg.CalcList)
}

func (e *Engine) UpdateHour(ownerId int32, eventType string, field string, value int64, cfg *mdx_cfg.TimeCfg) {
	key := fmt.Sprintf(KeyPatton, ownerId, eventType, time.Now().Format("2006-01-02-15")) //年-月-日-时
	lifeDuration := time.Duration(cfg.LiftTime) * time.Hour
	e.UpdateByTimeDimension(key, field, value, lifeDuration, cfg.CalcList)
}

func (e *Engine) UpdateMinute(ownerId int32, eventType string, field string, value int64, cfg *mdx_cfg.TimeCfg) {
	key := fmt.Sprintf(KeyPatton, ownerId, eventType, time.Now().Format("2006-01-02-15-04")) //年-月-日-时-分
	lifeDuration := time.Duration(cfg.LiftTime) * time.Minute
	e.UpdateByTimeDimension(key, field, value, lifeDuration, cfg.CalcList)
}

func (e Engine) UpdateByTimeDimension(key, field string, value int64, expire time.Duration, calcTypes []mdx_cfg.CalcType) {
	var err error
	for _, calcType := range calcTypes {
		fieldExt := field + "-" + string(calcType)
		switch calcType {
		case mdx_cfg.Count:
			_, err = e.DS.Increase(key, fieldExt, 1, expire)
		case mdx_cfg.Value:
			_, err = e.DS.Increase(key, fieldExt, value, expire)
		case mdx_cfg.Max:
			_, err = e.DS.UpdateMax(key, fieldExt, value, expire)
		case mdx_cfg.Min:
			_, err = e.DS.UpdateMinus(key, fieldExt, value, expire)
		}
		if err != nil {
			fmt.Printf("Engine.UpdateByTimeDimension failed, key=%s, field=%s, value=%d, err: %v", key, field, value, err)
		}
	}
}
