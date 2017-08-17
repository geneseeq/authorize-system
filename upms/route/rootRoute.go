package route

import (
	"net/http"

	"github.com/geneseeq/authorize-system/upms/association/groups"
	"github.com/geneseeq/authorize-system/upms/association/roles"
	"github.com/geneseeq/authorize-system/upms/association/users"
	"github.com/geneseeq/authorize-system/upms/dataing"
	"github.com/geneseeq/authorize-system/upms/grouping"
	"github.com/geneseeq/authorize-system/upms/roleing"
	"github.com/geneseeq/authorize-system/upms/servicing"
	"github.com/geneseeq/authorize-system/upms/usering"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// InitRouter init router
func InitRouter(logger log.Logger, httpLogger log.Logger, fieldKeys []string) {
	gs := initGroupRouter(logger, fieldKeys)
	as := initRelationRouter(logger, fieldKeys)
	us := initUserRouter(logger, fieldKeys)
	rs := initRoleRouter(logger, fieldKeys)
	gus := initUserRelationRouter(logger, fieldKeys)
	grs := initRoleRelationRouter(logger, fieldKeys)
	ras := initAuthorityRelationRouter(logger, fieldKeys)
	ss := initSetRouter(logger, fieldKeys)
	service := initServiceRouter(logger, fieldKeys)
	mux := http.NewServeMux()

	mux.Handle("/grouping/v1/", grouping.MakeHandler(gs, httpLogger))
	mux.Handle("/usering/v1/", usering.MakeHandler(us, httpLogger))
	mux.Handle("/roleing/v1/", roleing.MakeHandler(rs, httpLogger))
	mux.Handle("/releation/v1/user/", users.MakeHandler(as, httpLogger))
	mux.Handle("/releation/v1/group/", groups.MakeHandler(gus, grs, httpLogger))
	mux.Handle("/releation/v1/role/", roles.MakeHandler(ras, httpLogger))
	mux.Handle("/seting/v1/", dataing.MakeHandler(ss, httpLogger))
	mux.Handle("/servicing/v1/", servicing.MakeHandler(service, httpLogger))

	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
