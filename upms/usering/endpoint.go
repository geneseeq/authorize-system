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

type deleteMutliUserRequest struct {
	ListID []string
}

type putUserRequest struct {
	ID   string
	User User
}

type listUserRequest struct{}

type postUserRequest struct {
	User []User
}

type postUserResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

// userResponse User must equal User type
type userResponse struct {
	User []User `json:"content"`
	Err  error  `json:"error,omitempty"`
}

func (r userResponse) error() error { return r.Err }

func (r postUserResponse) error() error { return r.Err }

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
		sucessed, failed, err := s.PostUser(req.User)
		if err == nil {
			return postUserResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add user sucessed"}, nil
		}
		return postUserResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makeDeleteUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserRequest)
		err := s.DeleteUser(req.ID)
		if err == nil {
			return postUserResponse{
				Err:     err,
				Status:  200,
				Content: "delete user sucessed"}, nil
		}

		return postUserResponse{
			Err:     err,
			Status:  300,
			Content: "delete user failed"}, nil
	}
}

func makeDeleteMultiUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteMutliUserRequest)
		sucessed, failed, err := s.DeleteMultiUser(req.ListID)
		if err == nil {
			return postUserResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "delete mutli user sucessed"}, nil
		}
		return postUserResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "delete mutli user failed"}, nil
	}
}

func makePutUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(putUserRequest)
		err := s.PutUser(req.ID, req.User)
		if err == nil {
			return postUserResponse{
				Err:     err,
				Status:  200,
				Content: "update user sucessed"}, nil
		}
		return postUserResponse{
			Err:     err,
			Status:  300,
			Content: "update user failed"}, nil
	}
}

func makePutMultiUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postUserRequest)
		sucessed, failed, err := s.PutMultiUser(req.User)
		if err == nil {
			return postUserResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "update user sucessed"}, nil
		}
		return postUserResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "update user failed"}, nil
	}
}
