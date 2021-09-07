package data

import (
	"context"
	"dy201.com/doyle-blog/app/user/service/internal/biz"
	"dy201.com/doyle-blog/app/user/service/internal/pkg"
	"github.com/go-kratos/kratos/v2/errors"
	log "github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

var (
	USER_NOT_EXIST = errors.New(500, "USER_NOT_EXIST", "delete user , but user is not exist")
)
// 把 Data 包装到 Repo 中
// userRepo 实现 biz 的接口
type userRepo struct {
	data *Data
	log  *log.Helper
}

func (ur *userRepo) GetUserList(ctx context.Context, pageSize uint32, pageNum uint32) ([]biz.User, error) {
	var users []biz.User
	result := ur.data.db.Table("user").Scopes(pkg.Paginate(int(pageNum),int(pageSize))).Find(&users)

	if result.Error != nil{
		return nil, result.Error
	}

	log.DefaultLogger.Log(log.LevelInfo,"数据库取值", len(users), users)

	return users, nil
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	result := ur.data.db.Table("user").Create(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (ur *userRepo) GetUser(ctx context.Context, id uint32) (*biz.User, error) {
	var user biz.User
	result := ur.data.db.Table("user").First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user , nil
}

func (ur *userRepo) UpdateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	var user biz.User

	result := ur.data.db.Model(&user).Updates(&u)

	if result.Error != nil {
		return nil , result.Error
	}

	return  &user, nil
}

func (ur *userRepo) DeleteUser(ctx context.Context, id uint32) (bool, error) {
	result := ur.data.db.Table("user").Delete(&biz.User{}, id)
	if result.Error != nil {
		return false , result.Error
	}

	if result.RowsAffected == 0 {
		return false, USER_NOT_EXIST
	}

	return  true , nil
}

