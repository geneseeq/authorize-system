package usering

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

// func (s *instrumentingService) PostUser(user User) error {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "book").Add(1)
// 		s.requestLatency.With("method", "book").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.Store(user)
// }

func (s *instrumentingService) GetUser(id string) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "load").Add(1)
		s.requestLatency.With("method", "load").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetUser(id)
}

// func (s *instrumentingService) GetAllUser() []User {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "request_routes").Add(1)
// 		s.requestLatency.With("method", "request_routes").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.GetAllUser()
// }

// func (s *instrumentingService) PutUser(id string, user User) (err error) {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "assign_to_route").Add(1)
// 		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.PutUser(id, user)
// }

// func (s *instrumentingService) DeleteUser(id string) (err error) {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "change_destination").Add(1)
// 		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.DeleteUser(id)
// }
