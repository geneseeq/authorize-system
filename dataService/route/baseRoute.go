package route

import (
	"github.com/geneseeq/authorize-system/dataService/action"
	"github.com/geneseeq/authorize-system/dataService/baseing"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initDataRouter(logger log.Logger, fieldKeys []string) baseing.Service {
	var datas = action.NewBaseDataDBRepository("test", "data_infos")
	var rs baseing.Service
	rs = baseing.NewService(datas)
	rs = baseing.NewLoggingService(log.With(logger, "component", "baseing"), rs)
	rs = baseing.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "baseing_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "baseing_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		rs,
	)
	return rs
}
