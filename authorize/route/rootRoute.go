package route

import (
	"net/http"

	"github.com/geneseeq/authorize-system/authorize/authing"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// InitRouter init router
func InitRouter(logger log.Logger, httpLogger log.Logger, fieldKeys []string) {
	tokenData := initTokenRouter(logger, fieldKeys)

	mux := http.NewServeMux()

	mux.Handle("/authing/v1/", authing.MakeHandler(tokenData, httpLogger))

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
