package mdx_cfg

type TimeType int

const (
	Total = TimeType(iota)
	Year
	Month
	Day
	Hour
	Minute
)

type CalcType string

const (
	Count = CalcType("count")
	Value = CalcType("value")
	Max   = CalcType("max")
	Min   = CalcType("min")
)

var DefaultEventCfg = &EventCfg{}

func init() {
	DefaultEventCfg.TimeCfgList = make([]*TimeCfg, 1)
	DefaultEventCfg.TimeCfgList[0].Type = Total
	DefaultEventCfg.TimeCfgList[0].LiftTime = -1
	DefaultEventCfg.TimeCfgList[0].CalcList = []CalcType{Count, Value}
}

type Cfg struct {
	EventCfgs map[string]*EventCfg
}

type EventCfg struct {
	EventType   string
	TimeCfgList []*TimeCfg
}

type TimeCfg struct {
	Type     TimeType
	LiftTime int32
	CalcList []CalcType
}
