package example

import (
	"github.com/fangwei25/gomdx/src/data_source/redis"
	"github.com/fangwei25/gomdx/src/mdx_cfg"
	"github.com/fangwei25/gomdx/src/mdx_engine"
)

func main() {
	engine := mdx_engine.CreateEngine(redis.CreateDataSource(), &mdx_cfg.Cfg{EventCfgs: map[string]*mdx_cfg.EventCfg{}})
	engine.Update(123, "test1", "first", 1)
	engine.Update(123, "test1", "first", 2)
	engine.Update(123, "test1", "first", 3)

	engine.Update(123, "test1", "second", 4)
	engine.Update(123, "test1", "second", 5)
	engine.Update(123, "test1", "second", 6)

	engine.Update(123, "dev", "one", 1)
	engine.Update(123, "dev", "one", 2)
	engine.Update(123, "dev", "one", 3)

	engine.Update(123, "dev", "two", 4)
	engine.Update(123, "dev", "two", 5)
	engine.Update(123, "dev", "two", 6)
}
