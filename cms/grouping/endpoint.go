package grouping

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getGroupRequest struct {
	ID string
}

type deleteGroupRequest struct {
	ID string
}

type deleteMutliGroupRequest struct {
	ListId []string
}

type listGroupRequest struct{}

type postGroupRequest struct {
	Group []Group
}

type putGroupRequest struct {
	ID    string
	Group Group
}

// groupResponse User must equal User type
type groupResponse struct {
	Group []Group `json:"content"`
	Err   error   `json:"error,omitempty"`
}

type postGroupResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedId []string `json:"sucessedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

func (r groupResponse) error() error { return r.Err }

func makeGetGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getGroupRequest)
		result, err := s.GetGroup(req.ID)
		var groups []Group
		groups = append(groups, result)
		return groupResponse{Group: groups, Err: err}, nil
	}
}

func makeGetAllGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listGroupRequest)
		result, err := s.GetAllGroup()
		return groupResponse{Group: result, Err: err}, nil
	}
}

func makePostGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postGroupRequest)
		ids, err := s.PostGroup(req.Group)
		if err == nil {
			return postGroupResponse{SucessedId: ids, Err: err, Status: 200, Content: "add user sucessed"}, nil
		}
		return postGroupResponse{SucessedId: ids, Err: err, Status: 300, Content: "add user failed"}, nil
	}
}

func makeDeleteGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteGroupRequest)
		err := s.DeleteGroup(req.ID)
		if err == nil {
			return postGroupResponse{Err: err, Status: 200, Content: "delete user sucessed"}, nil
		}
		return postGroupResponse{Err: err, Status: 300, Content: "delete user failed"}, nil
	}
}

func makeDeleteMultiGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteMutliGroupRequest)
		ids, err := s.DeleteMultiGroup(req.ListId)
		if err == nil {
			return postGroupResponse{SucessedId: ids, Err: err, Status: 200, Content: "delete mutli user sucessed"}, nil
		}
		return postGroupResponse{SucessedId: ids, Err: err, Status: 300, Content: "delete mutli user failed"}, nil
	}
}

func makePutGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(putGroupRequest)
		err := s.PutGroup(req.ID, req.Group)
		if err == nil {
			return postGroupResponse{Err: err, Status: 200, Content: "update user sucessed"}, nil
		}
		return postGroupResponse{Err: err, Status: 300, Content: "update user failed"}, nil
	}
}

func makePutMultiGroupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postGroupRequest)
		ids, err := s.PutMultiGroup(req.Group)
		if err == nil {
			return postGroupResponse{SucessedId: ids, Err: err, Status: 200, Content: "update user sucessed"}, nil
		}
		return postGroupResponse{SucessedId: ids, Err: err, Status: 300, Content: "update user failed"}, nil
	}
}
