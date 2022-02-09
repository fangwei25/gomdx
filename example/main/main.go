package main

import (
	"encoding/json"
	"fmt"
	"github.com/fangwei25/gomdx/cfg"
	"github.com/fangwei25/gomdx/datasource/memery"
	"github.com/fangwei25/gomdx/engine"
	"time"
)

func main() {
	//engine := engine.CreateEngine(redis.CreateDataSource(), &cfg.Cfg{EventCfgs: map[string]*cfg.EventCfg{}})
	fmt.Println("example start")
	engine := engine.CreateEngine(memery.CreateDataSource(), &cfg.Cfg{EventCfgs: map[string]*cfg.EventCfg{}, KeyPrefix: "BB"})
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
	eventCfg := &cfg.EventCfg{
		EventType:   eventType3,
		TimeCfgList: make(map[cfg.TimeDimension]*cfg.TimeCfg),
	}
	eventCfg.TimeCfgList[cfg.TDMinute] = &cfg.TimeCfg{
		Type:     cfg.TDMinute,
		LiftTime: 1,
		CalcList: map[cfg.CalcType]bool{cfg.CTCount: true, cfg.CTValue: true, cfg.CTMax: true, cfg.CTMin: true},
	}

	engine.Cfg.Add(eventCfg)
	engine.Update(ownerId, eventType3, "new", 1)
	engine.Update(ownerId, eventType3, "new", 2)
	engine.Update(ownerId, eventType3, "new", 5)

	res, _ := engine.QueryOne(ownerId, eventType3, "new", cfg.TDMinute, time.Now())
	resByte, _ := json.Marshal(res)
	fmt.Println(string(resByte))

	pause := make(chan bool)
	go func() {
		time.Sleep(100 * time.Second)
		pause <- true
	}()
	<-pause
}
