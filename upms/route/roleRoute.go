package route

import (
	"github.com/geneseeq/authorize-system/upms/action"
	"github.com/geneseeq/authorize-system/upms/roleing"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initRoleRouter(logger log.Logger, fieldKeys []string) roleing.Service {
	var roles = action.NewRoleDBRepository("test", "role_infos")
	var rs roleing.Service
	rs = roleing.NewService(roles)
	rs = roleing.NewLoggingService(log.With(logger, "component", "roleing"), rs)
	rs = roleing.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "roleing_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "roleing_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		rs,
	)
	return rs
}
