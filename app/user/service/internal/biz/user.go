package biz

import (
	"context"
	log2 "github.com/go-kratos/kratos/v2/log"
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
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id uint32) (*User, error)
	UpdateUser(ctx context.Context, u *User) (*User, error)
	DeleteUser(ctx context.Context, id uint32) (bool, error)
	GetUserList(ctx context.Context, size uint32, num uint32) ([]User, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log2.Helper
}

func NewUserUseCase(repo UserRepo, logger log2.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  log2.NewHelper(log2.With(logger, "module", "usecase/user")),
	}
}

func (u *UserUseCase) GetUser(ctx context.Context, id uint32) (*User, error) {
	return u.repo.GetUser(ctx, id)
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUseCase) Update(ctx context.Context, user *User) (*User, error) {
	return u.repo.UpdateUser(ctx, user)
}

func (u *UserUseCase) Delete(ctx context.Context, id uint32) (bool, error) {

	return u.repo.DeleteUser(ctx, id)
}

func (u *UserUseCase) GetUserList(ctx context.Context, pageSize uint32, pageNum uint32) ([]User, error) {
	return u.repo.GetUserList(ctx , pageSize, pageNum)
}
