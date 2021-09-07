package server

import (
	v1 "dy201.com/doyle-blog/api/blog/admin/v1"
	conf "dy201.com/doyle-blog/app/blog/admin/internal/config"
	"dy201.com/doyle-blog/app/blog/admin/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func NewHTTPServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.BlogAdminServer) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			validate.Validator(),
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
	v1.RegisterBlogAdminServiceHTTPServer(srv, s)
	return srv
}
