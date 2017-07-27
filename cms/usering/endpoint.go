package usering

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getUserRequest struct {
	ID string
}

type userResponse struct {
	content *User `json:"content,omitempty"`
	Err     error `json:"error,omitempty"`
}

func (r userResponse) error() error { return r.Err }

func makeGetUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		result, err := s.GetUser(req.ID)
		return userResponse{content: &result, Err: err}, nil
	}
}
