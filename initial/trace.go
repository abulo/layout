package initial

import (
	"github.com/abulo/ratel/v3/trace"
	"github.com/abulo/ratel/v3/trace/jaeger"
	"github.com/abulo/ratel/v3/util"
)

func (initial *Initial) InitTrace() {
	opt := jaeger.NewJaeger()
	conf := initial.Config.Get("trace")
	res := conf.(map[string]interface{})
	opt.EnableRPCMetrics = util.ToBool(res["EnableRPCMetrics"])
	opt.LocalAgentHostPort = util.ToString(res["LocalAgentHostPort"])
	opt.LogSpans = util.ToBool(res["LogSpans"])
	opt.Param = util.ToFloat64(res["Param"])
	opt.PanicOnError = util.ToBool(res["PanicOnError"])
	client := opt.Build().Build()
	trace.SetGlobalTracer(client)
}
