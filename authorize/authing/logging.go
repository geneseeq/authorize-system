package authing

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) PostToken(token []Token) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "data",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostToken(token)
}

func (s *loggingService) GetToken(id string) (token Token, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetToken(id)
}

func (s *loggingService) GetAllToken() ([]Token, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllToken()
}

func (s *loggingService) PutMultiToken(token []Token) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiToken(token)
}

func (s *loggingService) DeleteMultiToken(listid []string) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiToken(listid)
}
