package server

import (
	"context"

	userv1 "github.com/aalindg/boilerplate/proto/user/v1"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (be *Backend) AddUser(ctx context.Context, ureq *userv1.AddUserRequest) (*userv1.AddUserResponse, error) {
	be.mu.Lock()
	defer be.mu.Unlock()

	newUser := &userv1.User{
		UserId:     uuid.Must(uuid.NewV4()).String(),
		Name:       ureq.GetName(),
		Email:      ureq.GetEmail(),
		CreateTime: timestamppb.Now(),
	}

	be.users = append(be.users, newUser)

	return &userv1.AddUserResponse{
		User: newUser,
	}, nil
}

func (be *Backend) ListUser(_ *userv1.ListUsersRequest, srv userv1.UserService_ListUsersServer) error {
	be.mu.RLock()
	defer be.mu.RUnlock()

	for _, _user := range be.users {
		err := srv.Send(&userv1.ListUsersResponse{User: _user})
		if err != nil {
			return err
		}
	}
	return nil
}

func (be *Backend) GetUser(ctx context.Context, ureq *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	be.mu.RLock()
	defer be.mu.RUnlock()

	for _, _user := range be.users {
		if ureq.GetUserId() == _user.UserId {
			return &userv1.GetUserResponse{
				User: _user,
			}, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "user with ID %q could not be found", ureq.GetUserId())
}
