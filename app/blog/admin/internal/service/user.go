package service

import (
	"context"
	v1 "dy201.com/doyle-blog/api/blog/admin/v1"
	"dy201.com/doyle-blog/app/blog/admin/internal/biz"
	"dy201.com/doyle-blog/app/blog/admin/internal/pkg/util"
	"github.com/golang/protobuf/ptypes/timestamp"
	"gorm.io/gorm"
	"time"
)

/**
BlogAdminServer 删除用户
 */
func (b *BlogAdminServer) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {
	ok, err := b.uc.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserReply{
		Ok: ok,
	},nil
}

/**
BlogAdminServer 根据id查询用户
 */
func (b *BlogAdminServer) SearchUserByID(ctx context.Context, req *v1.SearchUserByIDReq) (*v1.SearchUserByIDReply, error) {
	id := req.Id
	user, err := b.uc.SearchUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	reply := &v1.SearchUserByIDReply{User: &v1.User{
		Id:       uint32(user.Model.ID),
		UserName: user.UserName,
		Password: user.Password,
		Nickname: user.Nickname,
		Status:   user.Status,
		Email:    user.Email,
		Phone:    user.Phone,
		RoleId:   user.RoleID,
		Age:      user.Age,
	}}

	return reply , nil

}

func (b *BlogAdminServer) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	user := &biz.User{
		UserName: req.User.UserName,
		Password: req.User.Password,
		Nickname: req.User.Nickname,
		Status:   req.User.Status,
		Email:    req.User.Email,
		Phone:    req.User.Phone,
		Age:      req.User.Age,
		RoleID:   req.User.RoleId,
	}
	user, err := b.uc.CreateUser(ctx, *user)

	if err != nil {
		return nil, err
	}

	return &v1.CreateUserReply{
		User: &v1.User{
			Id:       uint32(user.ID),
			UserName: user.UserName,
			Password: user.Password,
			Nickname: user.Nickname,
			Email:    user.Email,
			Phone:    user.Phone,
			RoleId:   user.RoleID,
			Age:      user.Age,
		},
	},nil
}

func (b *BlogAdminServer) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {
	user := biz.User{
		Model:    gorm.Model{
			ID:        uint(req.User.Id),
			UpdatedAt: time.Now(),
		},
		UserName: req.User.UserName,
		Password: req.User.Password,
		Nickname: req.User.Nickname,
		Email:    req.User.Email,
		Phone:    req.User.Phone,
		Age:      req.User.Age,
		RoleID:   req.User.RoleId,
	}
	result, err := b.uc.UpdateUser(ctx, user)

	if err != nil{
		return nil, err
	}

	reply := &v1.UpdateUserReply{
		User: &v1.User{
			Id:       uint32(result.ID),
			UserName: result.UserName,
			Password: result.Password,
			Nickname: result.Nickname,
			Status:   result.Status,
			Email:    result.Email,
			Phone:    result.Phone,
			RoleId:   result.RoleID,
			Age:      result.Age,
		},
	}
	return reply, nil
}



func (b *BlogAdminServer) Login(context.Context, *v1.LoginReq) (*v1.LoginReply, error) {
	panic("implement me")
}

func (b *BlogAdminServer) GetUserList(ctx context.Context,req *v1.GetUserListReq) (*v1.GetUserListReply, error) {
	pageNum := req.PageInfo.PageNum
	pageSize := req.PageInfo.PageSize
	userList, err := b.uc.GetUserList(ctx, pageNum, pageSize)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetUserListReply{
		Users: make([]*v1.User, 0),
	}

	for _, user := range userList {
		item := &v1.User{
			Id:       uint32(user.ID),
			UserName: user.UserName,
			Password: user.Password,
			Nickname: user.Nickname,
			Status:   user.Status,
			Email:    user.Email,
			Phone:    user.Phone,
			RoleId:   user.RoleID,
			Age:      user.Age,
			CreateAt: &timestamp.Timestamp{
				Seconds: user.CreatedAt.Unix(),
				Nanos:   int32(user.CreatedAt.Nanosecond()),
			},
			UpdateAt: &timestamp.Timestamp{
				Seconds: user.UpdatedAt.Unix(),
				Nanos:   int32(user.UpdatedAt.Nanosecond()),
			},
		}
		reply.Users = append(reply.Users, item)
	}

	return reply, nil
}

func (b *BlogAdminServer) Logout(context.Context, *v1.LogoutReq) (*v1.LogoutReply, error) {
	panic("implement me")
}

func (b *BlogAdminServer) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	user := &biz.User{
		UserName: req.UserName,
		Password: util.PasswordWithSalt(req.Password),
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Age:      req.Age,
	}
	ok, err := b.uc.Register(ctx, user)
	if err != nil {
		return &v1.RegisterReply{
			Ok: ok,
		}, err
	}

	return &v1.RegisterReply{
		Ok: ok,
	} , nil
}
