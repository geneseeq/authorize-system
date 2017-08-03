package roleing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getRoleRequest struct {
	ID string
}

// roleResponse User must equal User type
type roleResponse struct {
	Role []Role `json:"content"`
	Err  error  `json:"error,omitempty"`
}

func (r roleResponse) error() error { return r.Err }

func makeGetRoleEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getRoleRequest)
		result, err := s.GetRole(req.ID)
		var roles []Role
		roles = append(roles, result)
		return roleResponse{Role: roles, Err: err}, nil
	}
}

// func makeGetAllGroupEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		_ = request.(listGroupRequest)
// 		result, err := s.GetAllGroup()
// 		return groupResponse{Group: result, Err: err}, nil
// 	}
// }

// func makePostGroupEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(postGroupRequest)
// 		ids, err := s.PostGroup(req.Group)
// 		if err == nil {
// 			return postGroupResponse{SucessedId: ids, Err: err, Status: 200, Content: "add user sucessed"}, nil
// 		}
// 		return postGroupResponse{SucessedId: ids, Err: err, Status: 300, Content: "add user failed"}, nil
// 	}
// }

// func makeDeleteGroupEndpoint(s Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(deleteGroupRequest)
// 		err := s.DeleteGroup(req.ID)
// 		if err == nil {
// 			return postGroupResponse{Err: err, Status: 200, Content: "delete user sucessed"}, nil
// 		}
// 		return postGroupResponse{Err: err, Status: 300, Content: "delete user failed"}, nil
// 	}
// }

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
