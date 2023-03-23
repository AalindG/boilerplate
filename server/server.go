package server

import (
	"sync"

	todo_listv1 "github.com/aalindg/boilerplate/proto/todo_list/v1"
	userv1 "github.com/aalindg/boilerplate/proto/user/v1"
)

// Backend implements the protobuf interface
type Backend struct {
	mu    *sync.RWMutex
	users []*userv1.User
	lists []todo_listv1.List
	userv1.UnimplementedUserServiceServer
	todo_listv1.UnimplementedTodoListServiceServer
}

// New initializes a new Backend struct.
func New() *Backend {
	return &Backend{
		mu: &sync.RWMutex{},
	}
}

// func (be *Backend) AddUser(ctx context.Context, ureq *userv1.AddUserRequest) (*userv1.AddUserResponse, error) {
// 	be.mu.Lock()
// 	defer be.mu.Unlock()

// 	newUser := &userv1.User{
// 		UserId:     uuid.Must(uuid.NewV4()).String(),
// 		Name:       ureq.GetName(),
// 		Email:      ureq.GetEmail(),
// 		CreateTime: timestamppb.Now(),
// 	}

// 	be.users = append(be.users, newUser)

// 	return &userv1.AddUserResponse{
// 		User: newUser,
// 	}, nil
// }

// func (be *Backend) ListUser(_ *userv1.ListUsersRequest, srv userv1.UserService_ListUsersServer) error {
// 	be.mu.RLock()
// 	defer be.mu.RUnlock()

// 	for _, _user := range be.users {
// 		err := srv.Send(&userv1.ListUsersResponse{User: _user})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (be *Backend) GetUser(ctx context.Context, ureq *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
// 	be.mu.RLock()
// 	defer be.mu.RUnlock()

// 	for _, _user := range be.users {
// 		if ureq.GetUserId() == _user.UserId {
// 			return &userv1.GetUserResponse{
// 				User: _user,
// 			}, nil
// 		}
// 	}
// 	return nil, status.Errorf(codes.NotFound, "user with ID %q could not be found", ureq.GetUserId())
// }

// func (be *Backend) CreateList(ctx context.Context, tlReq *todo_listv1.CreateListRequest) (*todo_listv1.CreateListResponse, error) {
// 	// be.mu.Lock()
// 	// defer be.mu.Unlock()

// 	if tlReq.Title == "" || tlReq.UserId == "" {
// 		return nil, status.Error(codes.FailedPrecondition, "Invalid parameters")
// 	}

// 	newTodoList := &todo_listv1.List{
// 		ListId:    uuid.Must(uuid.NewV4()).String(),
// 		Title:     tlReq.GetTitle(),
// 		CreatedBy: tlReq.GetUserId(),
// 	}

// 	be.lists = append(be.lists, *newTodoList)
// 	return &todo_listv1.CreateListResponse{
// 		List: newTodoList,
// 	}, nil
// }

// func (be *Backend) AddListElement(ctx context.Context, tlReq *todo_listv1.AddListElementRequest) (*todo_listv1.AddListElementResponse, error) {
// 	be.mu.Lock()
// 	defer be.mu.Unlock()

// 	if tlReq.GetListId() == "" {
// 		return nil, status.Error(codes.Internal, " List id is mandatory")
// 	}

// 	listElement := &todo_listv1.Element{
// 		Text:      tlReq.GetElement().Text,
// 		IsDone:    false,
// 		Priority:  tlReq.GetElement().GetPriority(),
// 		CreatedAt: timestamppb.Now(),
// 	}

// 	var _list *todo_listv1.List
// 	for _, list := range be.lists {
// 		if list.ListId == tlReq.GetListId() {
// 			_list = &list
// 			break
// 		}
// 	}

// 	if _list == nil {
// 		return nil, status.Error(codes.NotFound, "List not found: 404")
// 	}

// 	_list.Elements = append(_list.Elements, listElement)

// 	return &todo_listv1.AddListElementResponse{
// 		List: _list,
// 	}, nil
// }

// func (be *Backend) GetList(ctx context.Context, tlReq *todo_listv1.GetListRequest) (*todo_listv1.GetListResponse, error) {
// 	be.mu.RLock()
// 	defer be.mu.RUnlock()

// 	var _list *todo_listv1.List
// 	for _, list := range be.lists {
// 		if list.ListId == tlReq.GetListId() {
// 			_list = &list
// 			break
// 		}
// 	}

// 	if _list == nil {
// 		return nil, status.Error(codes.NotFound, "list  not found for listID: "+tlReq.GetListId())
// 	}

// 	return &todo_listv1.GetListResponse{
// 		List: _list,
// 	}, nil
// }

// // (_ *userv1.ListUsersRequest, srv userv1.UserService_ListUsersServer)
// func (be *Backend) FetchAllList(_ *todo_listv1.FetchAllListsRequest, srv todo_listv1.TodoListService_FetchAllListsServer) error {
// 	be.mu.RLock()
// 	defer be.mu.RUnlock()

// 	for _, _list := range be.lists {
// 		err := srv.Send(&todo_listv1.FetchAllListsResponse{
// 			List: &_list,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
