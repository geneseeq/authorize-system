package usering

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

func (s *loggingService) PostUser(user User) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "user",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostUser(user)
}

func (s *loggingService) GetUser(id string) (user User, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetUser(id)
}

func (s *loggingService) GetAllUser() ([]User, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllUser()
}

func (s *loggingService) PutUser(id string, user User) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutUser(id, user)
}

func (s *loggingService) DeleteUser(id string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteUser(id)
}

func (s *loggingService) DeleteMultiUser(listid []string) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiUser(listid)
}
