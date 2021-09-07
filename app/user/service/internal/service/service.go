package service

import (
	v1 "dy201.com/doyle-blog/api/user/service/v1"
	"dy201.com/doyle-blog/app/user/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServiceServer

	uc *biz.UserUseCase
	log *log.Helper

}

func NewUserService(uc *biz.UserUseCase, logger log.Logger) (*UserService) {
	return &UserService{
		uc: uc,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}