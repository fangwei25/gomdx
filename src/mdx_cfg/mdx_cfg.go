package mdx_cfg

import "sync"

type TimeDimension int

const (
	TDEver = TimeDimension(iota)
	TDYear
	TDMonth
	TDDay
	TDHour
	TDMinute
)

type CalcType string

const (
	CTCount = CalcType("count")
	CTValue = CalcType("value")
	CTMax   = CalcType("max")
	CTMin   = CalcType("min")
)

var DefaultEventCfg = &EventCfg{}

func init() {
	DefaultEventCfg.TimeCfgList = make([]*TimeCfg, 0)
	DefaultEventCfg.TimeCfgList = append(DefaultEventCfg.TimeCfgList, &TimeCfg{
		Type:     TDEver,
		LiftTime: -1,
		CalcList: []CalcType{CTCount, CTValue},
	})
}

type Cfg struct {
	KeyPrefix   string               `yaml:"key_prefix" json:"key_prefix,omitempty"` //数据key前缀
	EventCfgs   map[string]*EventCfg `yaml:"event_cfgs" json:"event_cfgs,omitempty"`
	EventLocker sync.RWMutex         `json:"-" yaml:"-"`
}

type EventCfg struct {
	EventType   string     `yaml:"event_type" json:"event_type,omitempty"`
	TimeCfgList []*TimeCfg `yaml:"time_cfg_list" json:"time_cfg_list,omitempty" `
}

type TimeCfg struct {
	Type     TimeDimension `yaml:"type" json:"type,omitempty"`
	LiftTime int32         `yaml:"lift_time" json:"lift_time,omitempty" `
	CalcList []CalcType    `yaml:"calc_list" json:"calc_list,omitempty" `
}

// Add 添加一个事件的配置 如果已存在该事件，则覆盖
func (c *Cfg) Add(eventCfg *EventCfg) {
	c.EventLocker.Lock()
	defer c.EventLocker.Unlock()
	c.EventCfgs[eventCfg.EventType] = eventCfg
}

// Remove 删除一个事件的配置
func (c *Cfg) Remove(EventType string) {
	c.EventLocker.Lock()
	defer c.EventLocker.Unlock()
	delete(c.EventCfgs, EventType)
}

//Get 获取一个事件配置
func (c *Cfg) Get(eventType string) *EventCfg {
	c.EventLocker.RLock()
	defer c.EventLocker.RUnlock()
	return c.EventCfgs[eventType]
}
