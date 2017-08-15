package route

import (
	"github.com/geneseeq/authorize-system/dataService/action"
	"github.com/geneseeq/authorize-system/dataService/labeling"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initLabelRouter(logger log.Logger, fieldKeys []string) labeling.Service {
	var datas = action.NewLabelDBRepository("test", "label_own_id")
	var rs labeling.Service
	rs = labeling.NewService(datas)
	rs = labeling.NewLoggingService(log.With(logger, "component", "labeling"), rs)
	rs = labeling.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "labeling_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "labeling_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		rs,
	)
	return rs
}
