package dataing

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

func (s *loggingService) PostDataSet(set []DataSet) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "user",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostDataSet(set)
}

func (s *loggingService) GetDataSet(id string) (set DataSet, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetDataSet(id)
}

func (s *loggingService) GetAllDataSet() ([]DataSet, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllDataSet()
}

func (s *loggingService) PutMultiDataSet(set []DataSet) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiDataSet(set)
}

func (s *loggingService) DeleteMultiDataSet(listid []string) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiDataSet(listid)
}
