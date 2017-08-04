package roleing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type baseRoleRequest struct {
	ID string
}

type listRoleRequest struct{}

type baseResponse struct {
	Role []Role `json:"content,omitempty"`
	Err  error  `json:"error,omitempty"`
}

type roleResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedId []string `json:"sucessedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

type postRoleRequest struct {
	Role []Role
}

func (r baseResponse) error() error { return r.Err }

func makeGetRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(baseRoleRequest)
		result, err := s.GetRole(req.ID)
		var roles []Role
		roles = append(roles, result)
		return baseResponse{Role: roles, Err: err}, nil
	}
}

func makeGetAllRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listRoleRequest)
		result, err := s.GetAllRole()
		return baseResponse{Role: result, Err: err}, nil
	}
}

func makePostRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postRoleRequest)
		ids, err := s.PostRole(req.Role)
		if err == nil {
			return roleResponse{SucessedId: ids, Err: err, Status: 200, Content: "add user sucessed"}, nil
		}
		return roleResponse{SucessedId: ids, Err: err, Status: 300, Content: "add user failed"}, nil
	}
}

func makeDeleteRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(baseRoleRequest)
		err := s.DeleteRole(req.ID)
		if err == nil {
			return roleResponse{Err: err, Status: 200, Content: "delete user sucessed"}, nil
		}
		return roleResponse{Err: err, Status: 300, Content: "delete user failed"}, nil
	}
}

// func makeDeleteMultiGroupEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(deleteMutliGroupRequest)
// 		ids, err := s.DeleteMultiGroup(req.ListId)
// 		if err == nil {
// 			return postGroupResponse{SucessedId: ids, Err: err, Status: 200, Content: "delete mutli user sucessed"}, nil
// 		}
// 		return postGroupResponse{SucessedId: ids, Err: err, Status: 300, Content: "delete mutli user failed"}, nil
// 	}
// }

// func makePutGroupEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(putGroupRequest)
// 		err := s.PutGroup(req.ID, req.Group)
// 		if err == nil {
// 			return postGroupResponse{Err: err, Status: 200, Content: "update user sucessed"}, nil
// 		}
// 		return postGroupResponse{Err: err, Status: 300, Content: "update user failed"}, nil
// 	}
// }

// func makePutMultiGroupEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(postGroupRequest)
// 		ids, err := s.PutMultiGroup(req.Group)
// 		if err == nil {
// 			return postGroupResponse{SucessedId: ids, Err: err, Status: 200, Content: "update user sucessed"}, nil
// 		}
// 		return postGroupResponse{SucessedId: ids, Err: err, Status: 300, Content: "update user failed"}, nil
// 	}
// }
