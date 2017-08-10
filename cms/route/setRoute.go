package route

import (
	"github.com/geneseeq/authorize-system/cms/action"
	"github.com/geneseeq/authorize-system/cms/dataing"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initSetRouter(logger log.Logger, fieldKeys []string) dataing.Service {
	var sets = action.NewSetDBRepository("test", "data_sets")
	var ss dataing.Service
	ss = dataing.NewService(sets)
	ss = dataing.NewLoggingService(log.With(logger, "component", "seting"), ss)
	ss = dataing.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "set_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "set_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		ss,
	)
	return ss
}
