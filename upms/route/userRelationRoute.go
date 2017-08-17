package route

import (
	"github.com/geneseeq/authorize-system/upms/action"
	"github.com/geneseeq/authorize-system/upms/association/users"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initRelationRouter(logger log.Logger, fieldKeys []string) users.Service {
	var relations = action.NewUserRelationRoleRepository("test", "user_own_roles")

	var as users.Service
	as = users.NewService(relations)
	as = users.NewLoggingService(log.With(logger, "component", "usersAssociation"), as)
	as = users.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "users_association_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "users_association_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		as,
	)
	return as
}
