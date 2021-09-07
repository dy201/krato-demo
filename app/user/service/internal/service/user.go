package service

import (
	"context"
	v1 "dy201.com/doyle-blog/api/user/service/v1"
	"dy201.com/doyle-blog/app/user/service/internal/biz"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func (u *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	user, err := u.uc.GetUser(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &v1.GetUserReply{
		User: &v1.User{
			Id:       uint32(user.ID),
			UserName: user.UserName,
			Password: user.Password,
			Nickname: user.Nickname,
			Status:   user.Status,
			Email:    user.Email,
			Phone:    user.Phone,
			RoleId:   user.RoleID,
			Age:      user.Age,
		},
	} , nil
}

func (u *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {

	user, err := u.uc.CreateUser(ctx, &biz.User{
		UserName: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Status:   req.Status,
		Email:    req.Email,
		Phone:    req.Phone,
		Age:      req.Age,
	})

	if err != nil {
		return nil, err
	}

	return &v1.CreateUserReply{
		User: &v1.User{
			Id:       uint32(user.ID),
			UserName: user.UserName,
			Password: user.Password,
			Nickname: user.Nickname,
			Status:   user.Status,
			Email:    user.Email,
			Phone:    user.Phone,
			RoleId:   user.RoleID,
			Age:      user.Age,
		},
	},nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserReply, error) {

	user, err := u.uc.Update(ctx, &biz.User{
		UserName: req.User.UserName,
		Password: req.User.Password,
		Nickname: req.User.Nickname,
		Status:   req.User.Status,
		Email:  req.User.Email,
		Phone:  req.User.Phone,
		Age:    req.User.Age,
		RoleID: req.User.RoleId,
	})

	if err != nil {
		return nil, err
	}

	return &v1.UpdateUserReply{
		User: &v1.User{
			Id:       uint32(user.ID),
			UserName: user.UserName,
			Password: user.Password,
			Nickname: user.Nickname,
			Status:   user.Status,
			Email:    user.Email,
			Phone:    user.Phone,
			RoleId:   user.RoleID,
			Age:      user.Age,
		},
	},nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (*v1.DeleteUserReply, error) {

	ok, err := u.uc.Delete(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &v1.DeleteUserReply{Ok:ok}, nil
}

func (u *UserService) GetUserList(ctx context.Context,req *v1.GetUserListReq) (*v1.GetUserListReply, error) {
	pageNum := req.PageInfo.PageNum
	pageSize := req.PageInfo.PageSize
	userList, err := u.uc.GetUserList(ctx, pageSize, pageNum)
	if err != nil {
		return nil, err
	}

	reply := &v1.GetUserListReply{}
	for _, user := range userList {

		var replyUser = &v1.User{
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
				Seconds: user.Model.CreatedAt.Unix(),
				Nanos:   int32(user.CreatedAt.Nanosecond()),
			},
			UpdateAt: &timestamp.Timestamp{
				Seconds: user.Model.UpdatedAt.Unix(),
				Nanos:   int32(user.UpdatedAt.Nanosecond()),
			},
			DeleteAt: &timestamp.Timestamp{
				Seconds: user.Model.DeletedAt.Time.Unix(),
				Nanos:   int32(user.DeletedAt.Time.Nanosecond()),
			},
		}
		reply.Users = append(reply.Users, replyUser)
	}
	return reply, nil
}
