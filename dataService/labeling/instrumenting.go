package labeling

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

func (s *instrumentingService) PostLabel(label []Label) ([]string, []string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "data").Add(1)
		s.requestLatency.With("method", "data").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PostLabel(label)
}

func (s *instrumentingService) GetLabel(id string) (label Label, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "load").Add(1)
		s.requestLatency.With("method", "load").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetLabel(id)
}

func (s *instrumentingService) GetAllLabel() ([]Label, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "request_routes").Add(1)
		s.requestLatency.With("method", "request_routes").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetAllLabel()
}

func (s *instrumentingService) PutMultiLabel(label []Label) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "assign_to_route").Add(1)
		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PutMultiLabel(label)
}

func (s *instrumentingService) DeleteMultiLabel(listid []string) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "change_destination").Add(1)
		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteMultiLabel(listid)
}
