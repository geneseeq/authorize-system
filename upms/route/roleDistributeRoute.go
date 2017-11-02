package route

import (
	"github.com/geneseeq/authorize-system/upms/action"
	"github.com/geneseeq/authorize-system/upms/distribute"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initRoleDistributeRouter(logger log.Logger, fieldKeys []string) distribute.Service {
	var groups = action.NewnewRoleDistributeDBRepository("test", "group_user_role_distribute")

	var gs distribute.Service
	gs = distribute.NewService(groups)
	gs = distribute.NewLoggingService(log.With(logger, "component", "distribute"), gs)
	gs = distribute.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "distribute_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "distribute_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		gs,
	)
	return gs
}
