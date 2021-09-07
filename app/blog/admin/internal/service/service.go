package service

import (
	v1 "dy201.com/doyle-blog/api/blog/admin/v1"
	"dy201.com/doyle-blog/app/blog/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogAdminServer)

type BlogAdminServer struct {
	v1.UnimplementedBlogAdminServiceServer

	log *log.Helper
	uc  *biz.UserUseCase
}



func NewBlogAdminServer(uc *biz.UserUseCase, logger log.Logger) *BlogAdminServer {
	return &BlogAdminServer{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
	}
}