package mdx_engine

import (
	"errors"
	"github.com/fangwei25/gomdx/src/iface"
	"github.com/fangwei25/gomdx/src/mdx_cfg"
)

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

// ReloadCfg 重新加载配置文件
func (e *Engine) ReloadCfg(cfg *mdx_cfg.Cfg) error {
	if nil == cfg {
		return errors.New("cfg is nil")
	}
	e.Cfg = cfg
	return nil
}
