package servicing

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

func (s *loggingService) PostService(service []Services) (sucessed []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "service",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostService(service)
}

func (s *loggingService) GetService(id string) (servcie Services, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetService(id)
}

func (s *loggingService) GetAllService() ([]Services, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllService()
}

func (s *loggingService) PutMultiService(service []Services) (sucessed []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiService(service)
}

func (s *loggingService) DeleteMultiService(listid []string) (sucessed []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiService(listid)
}
