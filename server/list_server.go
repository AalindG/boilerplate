package server

import (
	"context"

	todo_listv1 "github.com/aalindg/boilerplate/proto/todo_list/v1"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (be *Backend) CreateList(ctx context.Context, tlReq *todo_listv1.CreateListRequest) (*todo_listv1.CreateListResponse, error) {
	// be.mu.Lock()
	// defer be.mu.Unlock()

	if tlReq.Title == "" || tlReq.UserId == "" {
		return nil, status.Error(codes.FailedPrecondition, "Invalid parameters")
	}

	newTodoList := &todo_listv1.List{
		ListId:    uuid.Must(uuid.NewV4()).String(),
		Title:     tlReq.GetTitle(),
		CreatedBy: tlReq.GetUserId(),
	}

	be.lists = append(be.lists, *newTodoList)
	return &todo_listv1.CreateListResponse{
		List: newTodoList,
	}, nil
}

func (be *Backend) AddListElement(ctx context.Context, tlReq *todo_listv1.AddListElementRequest) (*todo_listv1.AddListElementResponse, error) {
	be.mu.Lock()
	defer be.mu.Unlock()

	if tlReq.GetListId() == "" {
		return nil, status.Error(codes.Internal, " List id is mandatory")
	}

	listElement := &todo_listv1.Element{
		Text:      tlReq.GetElement().Text,
		IsDone:    false,
		Priority:  tlReq.GetElement().GetPriority(),
		CreatedAt: timestamppb.Now(),
	}

	var _list *todo_listv1.List
	for _, list := range be.lists {
		if list.ListId == tlReq.GetListId() {
			_list = &list
			break
		}
	}

	if _list == nil {
		return nil, status.Error(codes.NotFound, "List not found: 404")
	}

	_list.Elements = append(_list.Elements, listElement)

	return &todo_listv1.AddListElementResponse{
		List: _list,
	}, nil
}

func (be *Backend) GetList(ctx context.Context, tlReq *todo_listv1.GetListRequest) (*todo_listv1.GetListResponse, error) {
	be.mu.RLock()
	defer be.mu.RUnlock()

	var _list *todo_listv1.List
	for _, list := range be.lists {
		if list.ListId == tlReq.GetListId() {
			_list = &list
			break
		}
	}

	if _list == nil {
		return nil, status.Error(codes.NotFound, "list  not found for listID: "+tlReq.GetListId())
	}

	return &todo_listv1.GetListResponse{
		List: _list,
	}, nil
}

// (_ *userv1.ListUsersRequest, srv userv1.UserService_ListUsersServer)
func (be *Backend) FetchAllList(_ *todo_listv1.FetchAllListsRequest, srv todo_listv1.TodoListService_FetchAllListsServer) error {
	be.mu.RLock()
	defer be.mu.RUnlock()

	for _, _list := range be.lists {
		err := srv.Send(&todo_listv1.FetchAllListsResponse{
			List: &_list,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
