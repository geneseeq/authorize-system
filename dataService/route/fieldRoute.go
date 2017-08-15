package route

import (
	"github.com/geneseeq/authorize-system/dataService/action"
	"github.com/geneseeq/authorize-system/dataService/fielding"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initFieldRouter(logger log.Logger, fieldKeys []string) fielding.Service {
	var datas = action.NewFieldDBRepository("test", "field_infos")
	var rs fielding.Service
	rs = fielding.NewService(datas)
	rs = fielding.NewLoggingService(log.With(logger, "component", "fielding"), rs)
	rs = fielding.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "fielding_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "fielding_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		rs,
	)
	return rs
}
