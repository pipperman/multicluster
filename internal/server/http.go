package server

import (
	"context"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	cluster "multicluster/api/cluster/v1"
	component "multicluster/api/component/v1"
	"multicluster/internal/conf"
	"multicluster/internal/middleware/auth"
	"multicluster/internal/service"
)

func NewSkipRoutersMatcher() selector.MatchFunc {

	skipRouters := map[string]struct{}{
		"/multicluster.v1.Cluster/Create":  {},
		"/multicluster.v1.Cluster/Get":     {},
		"/multicluster.v1.Cluster/List":    {},
		"/multicluster.v1.Cluster/Update":  {},
		"/multicluster.v1.Cluster/Delete":  {},
		"/multicluster.v1.Cluster/Log":     {},
		"/multicluster.v1.Cluster/Managed": {},
	}

	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, jwtc *conf.JWT, cls *service.MultiClusterService, logger log.Logger, tp *tracesdk.TracerProvider) *http.Server {
	otel.SetTracerProvider(tp)
	var opts = []http.ServerOption{
		http.ErrorEncoder(errorEncoder),
		http.Middleware(
			recovery.Recovery(),
			// 开放白名单
			//selector.Server(auth.JWTAuth(jwtc.Secret)).Match(NewSkipRoutersMatcher()).Build(),
			tracing.Server(),
			logging.Server(logger),
			metadata.Server(),
			auth.JWTAuth(jwtc.Secret),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
				metrics.WithRequests(prom.NewCounter(_metricRequests)),
			),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIHandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIHandler)
	srv.Handle("/metrics", promhttp.Handler())
	cluster.RegisterClusterHTTPServer(srv, cls)
	component.RegisterComponentHTTPServer(srv, cls)

	return srv
}
