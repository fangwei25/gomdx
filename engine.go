package gomdx

import (
	"errors"
	"github.com/fangwei25/gomdx/datasource"
)

type Engine struct {
	DS  datasource.DataSourcer
	Cfg *Cfg
}

func CreateEngine(ds datasource.DataSourcer, cfg *Cfg) *Engine {
	return &Engine{
		DS:  ds,
		Cfg: cfg,
	}
}

// ReloadCfg 重新加载配置文件
func (e *Engine) ReloadCfg(cfg *Cfg) error {
	if nil == cfg {
		return errors.New("cfg is nil")
	}
	e.Cfg = cfg
	return nil
}
