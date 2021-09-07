package data

import (
	"context"
	v1 "dy201.com/doyle-blog/api/user/service/v1"
	"dy201.com/doyle-blog/app/blog/admin/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) Register(ctx context.Context, user *biz.User) (bool, error) {
	_, err := u.data.uc.CreateUser(ctx, &v1.CreateUserReq{
		Username: user.UserName,
		Password: user.Password,
		Nickname: user.Nickname,
		Status:   user.Status,
		Email:    user.Email,
		Phone:    user.Phone,
		Age:      user.Age,
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *userRepo) CreateUser(ctx context.Context, user biz.User) (*biz.User, error) {
	result, err := u.data.uc.CreateUser(ctx, &v1.CreateUserReq{
		Username: user.UserName,
		Password: user.Password,
		Nickname: user.Nickname,
		Status:   user.Status,
		Email:    user.Email,
		Phone:    user.Phone,
		Age:      user.Age,
	})
	if err != nil {
		return nil, err
	}

	reply := &biz.User{
		Model: gorm.Model{
			ID:        uint(result.User.Id),
			CreatedAt: time.Now(),
		},
		UserName: result.User.UserName,
		Password: result.User.Password,
		Nickname: result.User.Nickname,
		Status:   result.User.Status,
		Email:    result.User.Email,
		Phone:    result.User.Phone,
		Age:      result.User.Age,
		RoleID:   result.User.RoleId,
	}

	return reply, nil
}

func (u *userRepo) SearchUserByID(ctx context.Context, id uint32) (*biz.User, error) {
	result, err := u.data.uc.GetUser(ctx, &v1.GetUserReq{Id: id})
	if err != nil {
		return nil, err
	}

	reply := &biz.User{
		Model: gorm.Model{
			ID:        uint(result.User.Id),
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		UserName: result.User.UserName,
		Password: result.User.Password,
		Nickname: result.User.Nickname,
		Email:    result.User.Email,
		Phone:    result.User.Phone,
		RoleID:   result.User.RoleId,
		Age:      result.User.Age,
	}
	return reply, nil
}

func (u userRepo) GetUserList(ctx context.Context, pageNum uint32, pageSize uint32) ([]biz.User, error) {

	list, err := u.data.uc.GetUserList(ctx, &v1.GetUserListReq{
		PageInfo: &v1.PageInfo{
			PageNum:  pageNum,
			PageSize: pageSize,
		},
	})

	log.DefaultLogger.Log(log.LevelInfo, "admin user 调用", list.Users, "长度", len(list.Users))

	if err != nil {
		return nil, err
	}

	result := make([]biz.User, 0)

	for _, user := range list.Users {
		log.DefaultLogger.Log(log.LevelInfo, "测试时间",  user.CreateAt.AsTime())
		u := biz.User{
			Model: gorm.Model{
				ID:        uint(user.Id),
				CreatedAt: user.CreateAt.AsTime(),
				UpdatedAt: user.UpdateAt.AsTime(),
			},
			UserName: user.UserName,
			Password: user.Password,
			Nickname: user.Nickname,
			Status:   user.Status,
			Email:    user.Email,
			Phone:    user.Phone,
			Age:      user.Age,
			RoleID:   user.RoleId,
		}
		result = append(result, u)
	}

	return result, nil
}

func (u userRepo) UpdateUser(ctx context.Context, user biz.User) (*biz.User, error) {
	panic("implement me")
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
