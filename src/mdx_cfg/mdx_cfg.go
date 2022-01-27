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
	DefaultEventCfg.TimeCfgList = make([]*TimeCfg, 0)
	DefaultEventCfg.TimeCfgList = append(DefaultEventCfg.TimeCfgList, &TimeCfg{
		Type:     TDEver,
		LiftTime: -1,
		CalcList: []CalcType{CTCount, CTValue},
	})
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
