package biz

import (
	"context"
	v1 "dy201.com/doyle-blog/api/user/service/v1"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"'user_name' not null comment('用户名') index VARCHAR(255)"`
	Password string `gorm:"'password' not null comment('加密用&分开') VARCHAR(255)"`
	Nickname string `gorm:"'nickname' comment('昵称') index VARCHAR(255)"`
	Status   int32  `gorm:"'status' comment('状态，1-启用，0-弃用') TINYINT"`
	Email    string `gorm:"'email' comment('邮箱') index VARCHAR(255)"`
	Phone    string `gorm:"'phone' comment('手机号') index VARCHAR(11)"`
	Age      uint32  `gorm:"'age' comment('年龄') INT"`
	RoleID   int32  `gorm:"'role_id' comment('角色id') INT"`
}

type UserRepo interface {
	GetUserList(ctx context.Context, pageNum uint32, pageSize uint32) ([]User, error)
	UpdateUser(ctx context.Context, user User) (*User, error)
	SearchUserByID(ctx context.Context, id uint32) (*User, error)
	CreateUser(ctx context.Context, user User)(*User, error)
	Register(ctx context.Context, user *User) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	us   v1.UserServiceClient
	log  *log.Helper
}

// 调用 user rpc服务，删除用户
func (c *UserUseCase) DeleteUser(ctx context.Context, id uint32) (bool ,error) {
	result, err := c.us.DeleteUser(ctx, &v1.DeleteUserReq{Id: id})

	if err != nil {
		return false, err
	}

	return result.Ok, nil
}


func (c *UserUseCase) UpdateUser(ctx context.Context, user User) (*User, error) {
	return c.repo.UpdateUser(ctx, user)
}

func (c *UserUseCase) GetUserList(ctx context.Context, pageNum uint32, pageSize uint32) ([]User, error) {
	return c.repo.GetUserList(ctx, pageNum, pageSize)
}

func (c *UserUseCase) SearchUserByID(ctx context.Context,id uint32) (*User , error) {
	return c.repo.SearchUserByID(ctx , id)
}

func (c *UserUseCase) CreateUser(ctx context.Context, user User) (*User , error) {
	return c.repo.CreateUser(ctx, user)
}

func (c *UserUseCase) Register(ctx context.Context, user *User)(bool, error) {
	return c.repo.Register(ctx, user)
}

func NewUserUseCase(repo UserRepo, logger log.Logger, us v1.UserServiceClient) *UserUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &UserUseCase{
		repo: repo,
		us:   us,
		log:  log,
	}
}
