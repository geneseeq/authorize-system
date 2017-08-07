package route

import (
	"github.com/geneseeq/authorize-system/cms/action"
	"github.com/geneseeq/authorize-system/cms/usering"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// InitUserRouter init router
func InitUserRouter(logger log.Logger, fieldKeys []string) usering.Service {
	var users = action.NewUserDBRepository("test", "user_infos")
	var us usering.Service
	us = usering.NewService(users)
	us = usering.NewLoggingService(log.With(logger, "component", "usering"), us)
	us = usering.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "usering_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "usering_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		us,
	)
	return us
}
