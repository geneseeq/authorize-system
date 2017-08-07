package route

import (
	"github.com/geneseeq/authorize-system/cms/action"
	"github.com/geneseeq/authorize-system/cms/grouping"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// InitGroupRouter init router
func InitGroupRouter(logger log.Logger, fieldKeys []string) grouping.Service {
	var groups = action.NewGroupDBRepository("test", "group_infos")

	var gs grouping.Service
	gs = grouping.NewService(groups)
	gs = grouping.NewLoggingService(log.With(logger, "component", "grouping"), gs)
	gs = grouping.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "grouping_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "grouping_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		gs,
	)
	return gs
}
