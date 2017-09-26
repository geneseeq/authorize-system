package grouping

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

func (s *instrumentingService) GetGroup(id string) (group Group, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "load").Add(1)
		s.requestLatency.With("method", "load").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetGroup(id)
}

func (s *instrumentingService) GetAllGroup() ([]Group, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "request_routes").Add(1)
		s.requestLatency.With("method", "request_routes").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetAllGroup()
}

func (s *instrumentingService) PostGroup(group []Group) ([]string, []string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "book").Add(1)
		s.requestLatency.With("method", "book").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PostGroup(group)
}

func (s *instrumentingService) DeleteGroup(id string) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "change_destination").Add(1)
		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteGroup(id)
}

func (s *instrumentingService) DeleteMultiGroup(listid []string) (sucessed []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "change_destination").Add(1)
		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteMultiGroup(listid)
}

func (s *instrumentingService) PutGroup(id string, group Group) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "assign_to_route").Add(1)
		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PutGroup(id, group)
}

func (s *instrumentingService) PutMultiGroup(group []Group) (sucessed []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "assign_to_route").Add(1)
		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PutMultiGroup(group)
}
