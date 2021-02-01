package tracer

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

//Trace  跟踪 一个trace 代表了一个食物或者流程在（分布式）系统中的执行过程
//Span   跨度 代表了一个食物中的每个工作单元，通常多个Span将会组成一个完成的Trace
//SpanContext   跨度上下文   代表一个食物的相关跟踪信息，不同的Span会根据OpenTracing 规范封装不同的属性，
//包含操作名称 开始时间和结束时间 标签信息 日志信息 上下文信息等

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	//jaeger client 配置项 主要设置应用的基本信息
	cfg := &config.Configuration{
		ServiceName: serviceName,
		//固定采样，对所有数据都逆行采样
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		//是否启用loggingReporter 刷新缓冲区的频率，上报的agent地址
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: 1 * time.Second,
			LogSpans:            true,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	//根据配置项初始化Tracer 对象
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	//设置全局的tracer对象
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
