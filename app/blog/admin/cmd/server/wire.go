// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"dy201.com/doyle-blog/app/blog/admin/internal/biz"
	conf "dy201.com/doyle-blog/app/blog/admin/internal/config"
	"dy201.com/doyle-blog/app/blog/admin/internal/data"
	"dy201.com/doyle-blog/app/blog/admin/internal/server"
	"dy201.com/doyle-blog/app/blog/admin/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}