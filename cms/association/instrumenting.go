package association

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

func (s *instrumentingService) GetRole(id string) (role Role, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "load").Add(1)
		s.requestLatency.With("method", "load").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetRole(id)
}

// func (s *instrumentingService) GetAllRole() ([]Role, error) {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "request_routes").Add(1)
// 		s.requestLatency.With("method", "request_routes").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.GetAllRole()
// }

// func (s *instrumentingService) PostRole(role []Role) ([]string, []string, error) {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "book").Add(1)
// 		s.requestLatency.With("method", "book").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.PostRole(role)
// }

// func (s *instrumentingService) DeleteRole(id string) (err error) {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "change_destination").Add(1)
// 		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.DeleteRole(id)
// }

// func (s *instrumentingService) DeleteMultiRole(listid []string) (sucessedIds []string, failedIds []string, err error) {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "change_destination").Add(1)
// 		s.requestLatency.With("method", "change_destination").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.DeleteMultiRole(listid)
// }

// func (s *instrumentingService) PutRole(id string, role Role) (err error) {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "assign_to_route").Add(1)
// 		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.PutRole(id, role)
// }

// func (s *instrumentingService) PutMultiRole(role []Role) (sucessedIds []string, failedIds []string, err error) {
// 	defer func(begin time.Time) {
// 		s.requestCount.With("method", "assign_to_route").Add(1)
// 		s.requestLatency.With("method", "assign_to_route").Observe(time.Since(begin).Seconds())
// 	}(time.Now())

// 	return s.Service.PutMultiRole(role)
// }
