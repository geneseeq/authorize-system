package labeling

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

func (s *loggingService) PostLabel(label []Label) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "data",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostLabel(label)
}

func (s *loggingService) GetLabel(id string) (label Label, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetLabel(id)
}

func (s *loggingService) GetAllLabel() ([]Label, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllLabel()
}

func (s *loggingService) PutMultiLabel(label []Label) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiLabel(label)
}

func (s *loggingService) DeleteMultiLabel(listid []string) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiLabel(listid)
}
