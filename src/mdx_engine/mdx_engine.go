package mdx_engine

import (
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
