package roles

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

func (s *loggingService) GetAuthorityFromRole(id string) (a Authority, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetAuthorityFromRole(id)
}

func (s *loggingService) GetAllAuthority() ([]Authority, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllAuthority()
}

func (s *loggingService) PostAuthority(a []Authority) (sucessedIds []string, failedIds []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "post",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostAuthority(a)
}

func (s *loggingService) DeleteMultiAuthority(a []DeleteData) (sucessedIds []string, failedIds []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiAuthority(a)
}

func (s *loggingService) PutMultiAuthority(a []Authority) (sucessedIds []string, failedIds []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiAuthority(a)
}
