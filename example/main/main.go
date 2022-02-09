package main

import (
	"encoding/json"
	"fmt"
	"github.com/fangwei25/gomdx"
	"github.com/fangwei25/gomdx/datasource/memery"
	"time"
)

func main() {
	//engine := engine.CreateEngine(redis.CreateDataSource(), &cfg.Cfg{EventCfgs: map[string]*cfg.EventCfg{}})
	fmt.Println("example start")
	engine := gomdx.CreateEngine(memery.CreateDataSource(), &gomdx.Cfg{EventCfgs: map[string]*gomdx.EventCfg{}, KeyPrefix: "BB"})
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
	eventCfg := &gomdx.EventCfg{
		EventType:   eventType3,
		TimeCfgList: make(map[gomdx.TimeDimension]*gomdx.TimeCfg),
	}
	eventCfg.TimeCfgList[gomdx.TDMinute] = &gomdx.TimeCfg{
		Type:     gomdx.TDMinute,
		LiftTime: 1,
		CalcList: map[gomdx.CalcType]bool{gomdx.CTCount: true, gomdx.CTValue: true, gomdx.CTMax: true, gomdx.CTMin: true},
	}

	engine.Cfg.Add(eventCfg)
	engine.Update(ownerId, eventType3, "new", 1)
	engine.Update(ownerId, eventType3, "new", 2)
	engine.Update(ownerId, eventType3, "new", 5)

	res, _ := engine.QueryOne(ownerId, eventType3, "new", gomdx.TDMinute, time.Now())
	resByte, _ := json.Marshal(res)
	fmt.Println(string(resByte))

	pause := make(chan bool)
	go func() {
		time.Sleep(100 * time.Second)
		pause <- true
	}()
	<-pause
}
