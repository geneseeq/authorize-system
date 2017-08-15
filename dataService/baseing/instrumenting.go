package baseing

import (
	"time"

	"github.com/geneseeq/authorize-system/dataService/data"
	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService) PostBaseData(data []BaseData) ([]string, []string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "data").Add(1)
		s.requestLatency.With("method", "data").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PostBaseData(data)
}

func (s *instrumentingService) GetBaseData(id string) (data BaseData, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "load").Add(1)
		s.requestLatency.With("method", "load").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetBaseData(id)
}

func (s *instrumentingService) GetAllBaseData() ([]BaseData, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "request_routes").Add(1)
		s.requestLatency.With("method", "request_routes").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetAllBaseData()
}

func (s *instrumentingService) PutMultiBaseData(data []BaseData) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "assign_to_route").Add(1)
		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PutMultiBaseData(data)
}

func (s *instrumentingService) DeleteMultiBaseData(listid []string) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "change_destination").Add(1)
		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteMultiBaseData(listid)
}

func (s *instrumentingService) DeleteMutliLabel(labelid []data.LabelIDModel) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "change_destination").Add(1)
		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteMutliLabel(labelid)
}
