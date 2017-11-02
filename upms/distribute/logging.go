package distribute

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

func (s *loggingService) GetRoleDistribute(id string) (distribute RoleDistribute, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetRoleDistribute(id)
}

func (s *loggingService) GetAllRoleDistribute() ([]RoleDistribute, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllRoleDistribute()
}

func (s *loggingService) PostRoleDistribute(distribute []RoleDistribute) (sucessed []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "distribute",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostRoleDistribute(distribute)
}

func (s *loggingService) DeleteMultiRoleDistribute(listid []string) (sucessed []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiRoleDistribute(listid)
}

func (s *loggingService) PutMultiRoleDistribute(distribute []RoleDistribute) (sucessed []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiRoleDistribute(distribute)
}
