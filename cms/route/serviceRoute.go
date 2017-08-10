package route

import (
	"github.com/geneseeq/authorize-system/cms/action"
	"github.com/geneseeq/authorize-system/cms/servicing"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initServiceRouter(logger log.Logger, fieldKeys []string) servicing.Service {
	var sets = action.NewServiceDBRepository("test", "service_infos")
	var ss servicing.Service
	ss = servicing.NewService(sets)
	ss = servicing.NewLoggingService(log.With(logger, "component", "servicing"), ss)
	ss = servicing.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "servicing_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "servicing_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		ss,
	)
	return ss
}
