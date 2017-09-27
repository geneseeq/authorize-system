// Package fielding provides the use-case of booking a cargo. Used by views
// facing an administrator.
package fielding

import (
	"errors"
	"time"

	"github.com/geneseeq/authorize-system/dataService/data"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
	ErrExceededMount   = errors.New("exceeded max mount")
	LimitMaxSum        = 50
)

// Service is the interface that provides booking methods.
type Service interface {
	GetField(id string) (Field, error)
	PostField(field []Field) ([]string, []string, error)
	GetAllField() ([]Field, error)
	DeleteMultiField(listid []string) ([]string, []string, error)
	PutMultiField(field []Field) ([]string, []string, error)
}

// Field is a basedata base info
type Field struct {
	ID         string    `json:"id"`
	Field      string    `json:"field"`
	Type       string    `json:"type"`
	Comment    string    `json:"comment"`
	UpdateTime time.Time `json:"update_time"`
	CreateTime time.Time `json:"create_time"`
}

type service struct {
	fields data.FieldRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(fields data.FieldRepository) Service {
	return &service{
		fields: fields,
	}
}

func (s *service) PostField(d []Field) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(d) < LimitMaxSum {
		for _, content := range d {
			curTime := data.TimeUtcToCst(time.Now())
			content.CreateTime = curTime
			content.UpdateTime = curTime
			err := s.fields.Store(fieldToFieldModel(content))
			if err != nil {
				failed = append(failed, content.ID)
				continue
			}
			sucessed = append(sucessed, content.ID)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func (s *service) GetField(id string) (Field, error) {
	if id == "" {
		return Field{}, ErrInvalidArgument
	}
	d, error := s.fields.Find(id)
	if error != nil {
		return Field{}, ErrNotFound
	}
	return fieldModelToField(d), nil
}

func (s *service) GetAllField() ([]Field, error) {
	var result []Field
	for _, d := range s.fields.FindFieldAll() {
		result = append(result, fieldModelToField(d))
	}
	return result, nil
}

func (s *service) PutMultiField(d []Field) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(d) < LimitMaxSum {
		for _, data := range d {
			if len(data.ID) == 0 {
				failed = append(failed, data.ID)
				continue
			}
			_, err := s.GetField(data.ID)
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			err = s.fields.Update(data.ID, fieldToFieldModel(data))
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			sucessed = append(sucessed, data.ID)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func (s *service) DeleteMultiField(listid []string) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(listid) == 0 {
		return sucessed, failed, ErrInvalidArgument
	}
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			error := s.fields.Remove(id)
			if error != nil {
				failed = append(failed, id)
				continue
			}
			sucessed = append(sucessed, id)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func fieldToFieldModel(d Field) *data.FieldModel {

	return &data.FieldModel{
		UnionID:    d.ID,
		ID:         d.ID,
		Field:      d.Field,
		Type:       d.Type,
		Comment:    d.Comment,
		UpdateTime: d.UpdateTime,
		CreateTime: d.CreateTime,
	}
}

func fieldModelToField(d *data.FieldModel) Field {
	return Field{
		ID:         d.ID,
		Field:      d.Field,
		Type:       d.Type,
		Comment:    d.Comment,
		UpdateTime: d.UpdateTime,
		CreateTime: d.CreateTime,
	}
}
