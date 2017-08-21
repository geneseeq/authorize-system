package authing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getTokenRequest struct {
	ID string
}

type deleteMutliTokenRequest struct {
	ListID []string
}

type listTokenRequest struct{}

type postTokenRequest struct {
	Token []Token
}

type postTokenResponse struct {
	//omitempty表示字段值为空，则不输出到json串
	Status     int      `json:"status"`
	Content    string   `json:"content"`
	SucessedID []string `json:"sucessedid,omitempty"`
	FailedID   []string `json:"failedid,omitempty"`
	Err        error    `json:"err,omitempty"`
}

// tokenResponse User must equal User type
type tokenResponse struct {
	Token []Token `json:"content"`
	Err   error   `json:"error,omitempty"`
}

func (r tokenResponse) error() error { return r.Err }

func (r postTokenResponse) error() error { return r.Err }

func makeGetTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getTokenRequest)
		result, err := s.GetToken(req.ID)
		var datas []Token
		datas = append(datas, result)
		return tokenResponse{Token: datas, Err: err}, nil
	}
}

func makeGetAllTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(listTokenRequest)
		result, err := s.GetAllToken()
		return tokenResponse{Token: result, Err: err}, nil
	}
}

func makePostTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postTokenRequest)
		sucessed, failed, err := s.PostToken(req.Token)
		if err == nil {
			return postTokenResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postTokenResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makeDeleteMultiTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteMutliTokenRequest)
		sucessed, failed, err := s.DeleteMultiToken(req.ListID)
		if err == nil {
			return postTokenResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postTokenResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}

func makePutMultiTokenEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postTokenRequest)
		sucessed, failed, err := s.PutMultiToken(req.Token)
		if err == nil {
			return postTokenResponse{
				SucessedID: sucessed,
				FailedID:   failed,
				Err:        err,
				Status:     200,
				Content:    "add data sucessed"}, nil
		}
		return postTokenResponse{
			SucessedID: sucessed,
			FailedID:   failed,
			Err:        err,
			Status:     300,
			Content:    "add user failed"}, nil
	}
}
