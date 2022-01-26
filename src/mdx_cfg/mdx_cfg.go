package mdx_cfg

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
	DefaultEventCfg.TimeCfgList = make([]*TimeCfg, 1)
	DefaultEventCfg.TimeCfgList[0].Type = TDEver
	DefaultEventCfg.TimeCfgList[0].LiftTime = -1
	DefaultEventCfg.TimeCfgList[0].CalcList = []CalcType{CTCount, CTValue}
}

type Cfg struct {
	KeyPrefix string //数据key前缀
	EventCfgs map[string]*EventCfg
}

type EventCfg struct {
	EventType   string
	TimeCfgList []*TimeCfg
}

type TimeCfg struct {
	Type     TimeDimension
	LiftTime int32
	CalcList []CalcType
}
