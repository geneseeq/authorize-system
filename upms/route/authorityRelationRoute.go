package route

import (
	"github.com/geneseeq/authorize-system/upms/action"
	"github.com/geneseeq/authorize-system/upms/association/roles"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initAuthorityRelationRouter(logger log.Logger, fieldKeys []string) roles.Service {
	var relations = action.NewRoleAuthorityRelationRepository("test", "role_own_permissions")

	var as roles.Service
	as = roles.NewService(relations)
	as = roles.NewLoggingService(log.With(logger, "component", "roleAuthorityAssociation"), as)
	as = roles.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "role_authority_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "role_authority_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		as,
	)
	return as
}
