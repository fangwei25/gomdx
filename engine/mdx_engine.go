package engine

import (
	"errors"
	"github.com/fangwei25/gomdx/cfg"
	"github.com/fangwei25/gomdx/iface"
)

type Engine struct {
	DS  iface.DataSourcer
	Cfg *cfg.Cfg
}

func CreateEngine(ds iface.DataSourcer, cfg *cfg.Cfg) *Engine {
	return &Engine{
		DS:  ds,
		Cfg: cfg,
	}
}

// ReloadCfg 重新加载配置文件
func (e *Engine) ReloadCfg(cfg *cfg.Cfg) error {
	if nil == cfg {
		return errors.New("cfg is nil")
	}
	e.Cfg = cfg
	return nil
}
