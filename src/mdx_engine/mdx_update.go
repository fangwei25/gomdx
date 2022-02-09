package mdx_engine

import (
	"fmt"
	"github.com/fangwei25/gomdx/src/mdx_cfg"
	"time"
)

// Update 数值统计 按照给定的数值进行累加统计
func (e *Engine) Update(ownerId int32, eventType string, field string, value int64) {
	//根据配置决定统计的计算维度（累加，计数，记录最大值，记录最小值）和时间维度(总，年，月，日，小时，分)
	eventCfg := e.Cfg.Get(eventType)
	if nil == eventCfg {
		//未配置的事件类型，则用默认配置： 时间维度=永久，计算维度=累加值、累加计数
		eventCfg = mdx_cfg.DefaultEventCfg
	}

	nonce := time.Now()
	for _, timeCfg := range eventCfg.TimeCfgList {
		key := e.GenKey(ownerId, eventType, timeCfg.Type, nonce)
		lifeTime := e.GetLifeTime(timeCfg.Type, timeCfg.LiftTime)
		e.UpdateByTimeDimension(key, field, value, lifeTime, timeCfg.CalcList)
	}
}

func (e Engine) UpdateByTimeDimension(key, field string, value int64, expire time.Duration, calcTypes map[mdx_cfg.CalcType]bool) {
	var err error
	for calcType, _ := range calcTypes {
		fieldExt := e.GenField(field, calcType)
		switch calcType {
		case mdx_cfg.CTCount:
			_, err = e.DS.Increase(key, fieldExt, 1, expire)
		case mdx_cfg.CTValue:
			_, err = e.DS.Increase(key, fieldExt, value, expire)
		case mdx_cfg.CTMax:
			_, err = e.DS.UpdateMax(key, fieldExt, value, expire)
		case mdx_cfg.CTMin:
			_, err = e.DS.UpdateMinus(key, fieldExt, value, expire)
		}
		if err != nil {
			fmt.Printf("Engine.UpdateByTimeDimension failed, key=%s, field=%s, value=%d, err: %v", key, field, value, err)
		}
	}
}
