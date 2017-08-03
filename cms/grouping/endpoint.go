package grouping

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getGroupRequest struct {
	ID string
}

type listGroupRequest struct{}

type postGroupRequest struct {
	Group []Group
}

// groupResponse User must equal User type
type groupResponse struct {
	Group []Group `json:"content,omitempty"`
	Err   error   `json:"error,omitempty"`
}

type postGroupResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedId []string `json:"sucessedid"`
	Err        error    `json:"err"`
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
