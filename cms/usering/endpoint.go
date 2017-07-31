package usering

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getUserRequest struct {
	ID string
}

type deleteUserRequest struct {
	ID string
}

type listUserRequest struct{}

type postUserRequest struct {
	User User
}

type postUserResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status  int    `json:"status"`
	Content string `json:"content"`
	Err     error  `json:"err,omitempty"`
}

// userResponse User must equal User type
type userResponse struct {
	User []User `json:"content,omitempty"`
	Err  error  `json:"error,omitempty"`
}

func (r userResponse) error() error { return r.Err }

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		result, err := s.GetUser(req.ID)
		var users []User
		users = append(users, result)
		return userResponse{User: users, Err: err}, nil
	}
}

func makeGetAllUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listUserRequest)
		result, err := s.GetAllUser()
		// var users []User
		// users = append(users, result)
		return userResponse{User: result, Err: err}, nil
	}
}

func makePostUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postUserRequest)
		err := s.PostUser(req.User)
		if err == nil {
			return postUserResponse{Err: err, Status: 200, Content: "sucessed"}, nil
		}
		return postUserResponse{Err: err, Status: 300, Content: "sucessed"}, nil
	}
}

func makeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserRequest)
		err := s.DeleteUser(req.ID)
		if err == nil {
			return postUserResponse{Err: err, Status: 200, Content: "sucessed"}, nil
		}
		return postUserResponse{Err: err, Status: 300, Content: "sucessed"}, nil
	}
}
