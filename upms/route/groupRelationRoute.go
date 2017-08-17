package route

import (
	"github.com/geneseeq/authorize-system/upms/action"
	"github.com/geneseeq/authorize-system/upms/association/groups"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func initUserRelationRouter(logger log.Logger, fieldKeys []string) groups.Service {
	var relations = action.NewGroupUserRelationRoleRepository("test", "group_own_users_and_roles")

	var as groups.Service
	as = groups.NewService(relations)
	as = groups.NewLoggingService(log.With(logger, "component", "groupsUserAssociation"), as)
	as = groups.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "group_user_association_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "group_user_association_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		as,
	)
	return as
}

func initRoleRelationRouter(logger log.Logger, fieldKeys []string) groups.Service {
	var relations = action.NewGroupRoleRelationRoleRepository("test", "group_own_users_and_roles")

	var as groups.Service
	as = groups.NewService(relations)
	as = groups.NewLoggingService(log.With(logger, "component", "groupsRoleAssociation"), as)
	as = groups.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "group_role_association_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "group_role_association_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		as,
	)
	return as
}
