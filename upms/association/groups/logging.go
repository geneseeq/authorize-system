package groups

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

func (s *loggingService) GetDataFromGroup(id string) (group Groups, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetDataFromGroup(id)
}

func (s *loggingService) GetAllData() ([]Groups, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllData()
}

func (s *loggingService) PostData(group []Groups) (sucessedIds []string, failedIds []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "post",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostData(group)
}

func (s *loggingService) DeleteMultiData(group []Groups) (sucessedIds []string, failedIds []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiData(group)
}

func (s *loggingService) PutMultiData(g []Groups) (sucessedIds []string, failedIds []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiData(g)
}
