package dataing

import (
	"time"

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

func (s *instrumentingService) PostDataSet(set []DataSet) ([]string, []string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "set").Add(1)
		s.requestLatency.With("method", "set").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PostDataSet(set)
}

func (s *instrumentingService) GetDataSet(id string) (set DataSet, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "load").Add(1)
		s.requestLatency.With("method", "load").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetDataSet(id)
}

func (s *instrumentingService) GetAllDataSet() ([]DataSet, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "request_routes").Add(1)
		s.requestLatency.With("method", "request_routes").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetAllDataSet()
}

func (s *instrumentingService) PutMultiDataSet(set []DataSet) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "assign_to_route").Add(1)
		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PutMultiDataSet(set)
}

func (s *instrumentingService) DeleteMultiDataSet(listid []string) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "change_destination").Add(1)
		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteMultiDataSet(listid)
}
