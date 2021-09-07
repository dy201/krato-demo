// +build wireinject


package main

import (
	"dy201.com/doyle-blog/app/user/service/internal/biz"
	"dy201.com/doyle-blog/app/user/service/internal/conf"
	"dy201.com/doyle-blog/app/user/service/internal/data"
	"dy201.com/doyle-blog/app/user/service/internal/server"
	"dy201.com/doyle-blog/app/user/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}