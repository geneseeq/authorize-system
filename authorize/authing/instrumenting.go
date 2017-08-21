package authing

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

func (s *instrumentingService) PostToken(token []Token) ([]string, []string, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "data").Add(1)
		s.requestLatency.With("method", "data").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PostToken(token)
}

func (s *instrumentingService) GetToken(id string) (token Token, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "load").Add(1)
		s.requestLatency.With("method", "load").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetToken(id)
}

func (s *instrumentingService) GetAllToken() ([]Token, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "request_routes").Add(1)
		s.requestLatency.With("method", "request_routes").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetAllToken()
}

func (s *instrumentingService) PutMultiToken(token []Token) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "assign_to_route").Add(1)
		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.PutMultiToken(token)
}

func (s *instrumentingService) DeleteMultiToken(listid []string) (ids []string, failed []string, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "change_destination").Add(1)
		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.DeleteMultiToken(listid)
}
