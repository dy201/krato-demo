package data

import (
	"context"
	v1 "dy201.com/doyle-blog/api/user/service/v1"
	conf "dy201.com/doyle-blog/app/blog/admin/internal/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserServiceClient,
	NewUserRepo,
)

// Data .
type Data struct {
	log *log.Helper
	uc  v1.UserServiceClient
}

// NewData .
func NewData(
	conf *conf.Data,
	logger log.Logger,
	uc v1.UserServiceClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, uc: uc,}, nil
}

func NewUserServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) v1.UserServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///doyle.blog.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := v1.NewUserServiceClient(conn)
	return c
}
