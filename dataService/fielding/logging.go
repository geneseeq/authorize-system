package fielding

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

func (s *loggingService) PostField(field []Field) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "data",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostField(field)
}

func (s *loggingService) GetField(id string) (field Field, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetField(id)
}

func (s *loggingService) GetAllField() ([]Field, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllField()
}

func (s *loggingService) PutMultiField(field []Field) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiField(field)
}

func (s *loggingService) DeleteMultiField(listid []string) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiField(listid)
}
