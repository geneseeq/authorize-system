package baseing

import (
	"time"

	"github.com/geneseeq/authorize-system/dataService/data"
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

func (s *loggingService) PostBaseData(data []BaseData) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "data",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PostBaseData(data)
}

func (s *loggingService) GetBaseData(id string) (data BaseData, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetBaseData(id)
}

func (s *loggingService) GetAllBaseData() ([]BaseData, error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "request_routes",
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.GetAllBaseData()
}

func (s *loggingService) PutMultiBaseData(data []BaseData) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "assign_to_route",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.PutMultiBaseData(data)
}

func (s *loggingService) DeleteMultiBaseData(listid []string) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMultiBaseData(listid)
}

func (s *loggingService) DeleteMutliLabel(labelid []data.LabelIDModel) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "change_destination",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteMutliLabel(labelid)
}
