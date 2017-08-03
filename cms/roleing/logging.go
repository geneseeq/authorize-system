package roleing

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

func (s *loggingService) GetRole(id string) (role Role, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "load",
			"tracking_id", id,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.GetRole(id)
}

// func (s *loggingService) GetAllGroup() ([]Group, error) {
// 	defer func(begin time.Time) {
// 		s.logger.Log(
// 			"method", "request_routes",
// 			"took", time.Since(begin),
// 		)
// 	}(time.Now())
// 	return s.Service.GetAllGroup()
// }

// func (s *loggingService) PostGroup(group []Group) (ids []string, err error) {
// 	defer func(begin time.Time) {
// 		s.logger.Log(
// 			"method", "group",
// 			"took", time.Since(begin),
// 			"err", err,
// 		)
// 	}(time.Now())
// 	return s.Service.PostGroup(group)
// }

// func (s *loggingService) DeleteGroup(id string) (err error) {
// 	defer func(begin time.Time) {
// 		s.logger.Log(
// 			"method", "change_destination",
// 			"took", time.Since(begin),
// 			"err", err,
// 		)
// 	}(time.Now())
// 	return s.Service.DeleteGroup(id)
// }

// func (s *loggingService) DeleteMultiGroup(listid []string) (ids []string, err error) {
// 	defer func(begin time.Time) {
// 		s.logger.Log(
// 			"method", "change_destination",
// 			"took", time.Since(begin),
// 			"err", err,
// 		)
// 	}(time.Now())
// 	return s.Service.DeleteMultiGroup(listid)
// }

// func (s *loggingService) PutGroup(id string, group Group) (err error) {
// 	defer func(begin time.Time) {
// 		s.logger.Log(
// 			"method", "assign_to_route",
// 			"tracking_id", id,
// 			"took", time.Since(begin),
// 			"err", err,
// 		)
// 	}(time.Now())
// 	return s.Service.PutGroup(id, group)
// }

// func (s *loggingService) PutMultiGroup(group []Group) (ids []string, err error) {
// 	defer func(begin time.Time) {
// 		s.logger.Log(
// 			"method", "assign_to_route",
// 			"took", time.Since(begin),
// 			"err", err,
// 		)
// 	}(time.Now())
// 	return s.Service.PutMultiGroup(group)
// }
