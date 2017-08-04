package roleing

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

func (s *loggingService) GetRole(id string) (role Role, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetRole(id)
}

func (s *loggingService) GetAllRole() ([]Role, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllRole()
}

func (s *loggingService) PostRole(role []Role) (sucessedIds []string, failedIds []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "post",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostRole(role)
}

func (s *loggingService) DeleteRole(id string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "delete",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteRole(id)
}

func (s *loggingService) DeleteMultiRole(listid []string) (sucessedIds []string, failedIds []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiRole(listid)
}

func (s *loggingService) PutRole(id string, role Role) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutRole(id, role)
}

func (s *loggingService) PutMultiRole(role []Role) (sucessedIds []string, failedIds []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiRole(role)
}
