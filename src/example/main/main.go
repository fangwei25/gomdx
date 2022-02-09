package main

import (
	"encoding/json"
	"fmt"
	"github.com/fangwei25/gomdx/src/data_source/memery"
	"github.com/fangwei25/gomdx/src/mdx_cfg"
	"github.com/fangwei25/gomdx/src/mdx_engine"
	"time"
)

func main() {
	//engine := mdx_engine.CreateEngine(redis.CreateDataSource(), &mdx_cfg.Cfg{EventCfgs: map[string]*mdx_cfg.EventCfg{}})
	fmt.Println("example start")
	engine := mdx_engine.CreateEngine(memery.CreateDataSource(), &mdx_cfg.Cfg{EventCfgs: map[string]*mdx_cfg.EventCfg{}, KeyPrefix: "BB"})
	ownerId := int32(123)
	eventType1 := "test"
	engine.Update(ownerId, eventType1, "first", 1)
	engine.Update(ownerId, eventType1, "first", 2)
	engine.Update(ownerId, eventType1, "first", 3)

	engine.Update(ownerId, eventType1, "second", 4)
	engine.Update(ownerId, eventType1, "second", 5)
	engine.Update(ownerId, eventType1, "second", 6)

	eventType2 := "hello"
	engine.Update(ownerId, eventType2, "one", 1)
	engine.Update(ownerId, eventType2, "one", 2)
	engine.Update(ownerId, eventType2, "one", 3)

	engine.Update(ownerId, eventType2, "two", 4)
	engine.Update(ownerId, eventType2, "two", 5)
	engine.Update(ownerId, eventType2, "two", 6)

	//动态添加一个配置
	eventType3 := "god"
	eventCfg := &mdx_cfg.EventCfg{
		EventType:   eventType3,
		TimeCfgList: make(map[mdx_cfg.TimeDimension]*mdx_cfg.TimeCfg),
	}
	eventCfg.TimeCfgList[mdx_cfg.TDMinute] = &mdx_cfg.TimeCfg{
		Type:     mdx_cfg.TDMinute,
		LiftTime: 1,
		CalcList: map[mdx_cfg.CalcType]bool{mdx_cfg.CTCount: true, mdx_cfg.CTValue: true, mdx_cfg.CTMax: true, mdx_cfg.CTMin: true},
	}

	engine.Cfg.Add(eventCfg)
	engine.Update(ownerId, eventType3, "new", 1)
	engine.Update(ownerId, eventType3, "new", 2)
	engine.Update(ownerId, eventType3, "new", 5)

	res, _ := engine.QueryOne(ownerId, eventType3, "new", mdx_cfg.TDMinute, time.Now())
	resByte, _ := json.Marshal(res)
	fmt.Println(string(resByte))

	pause := make(chan bool)
	go func() {
		time.Sleep(100 * time.Second)
		pause <- true
	}()
	<-pause
}
