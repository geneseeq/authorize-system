// Package labeling provides the use-case of booking a cargo. Used by views
// facing an administrator.
package labeling

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
	GetLabel(id string) (Label, error)
	PostLabel(label []Label) ([]string, []string, error)
	GetAllLabel() ([]Label, error)
	DeleteMultiLabel(listid []string) ([]string, []string, error)
	PutMultiLabel(label []Label) ([]string, []string, error)
}

// Label is a basedata base info
type Label struct {
	LabelID      string    `json:"label_id"`
	LabelType    string    `json:"label_type"`
	SubLableID   []string  `json:"sub_label_id"`
	SampleID     []string  `json:"sample_id"`
	OrderID      []string  `json:"order_id"`
	Action       []string  `json:"action"`
	UpdateTime   time.Time `json:"update_time"`
	CreateTime   time.Time `json:"create_time"`
	MedicalID    []string  `json:"medical_id"`
	CreateUserID string    `json:"create_user_id"` //创建人ID
	UpdateUserID string    `json:"update_user_id"`
	Buildin      bool      `json:"buildin"`
	Validity     bool      `json:"validity"`
}

type service struct {
	labels data.LableRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(labels data.LableRepository) Service {
	return &service{
		labels: labels,
	}
}

func (s *service) PostLabel(d []Label) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(d) < LimitMaxSum {
		for _, content := range d {
			curTime := data.TimeUtcToCst(time.Now())
			content.CreateTime = curTime
			content.UpdateTime = curTime
			err := s.labels.Store(labelToLabelModel(content))
			if err != nil {
				failed = append(failed, content.LabelID)
				continue
			}
			sucessed = append(sucessed, content.LabelID)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func (s *service) GetLabel(id string) (Label, error) {
	if id == "" {
		return Label{}, ErrInvalidArgument
	}
	d, error := s.labels.Find(id)
	if error != nil {
		return Label{}, ErrNotFound
	}
	return labelModellToLabel(d), nil
}

func (s *service) GetAllLabel() ([]Label, error) {
	var result []Label
	for _, d := range s.labels.FindLabelAll() {
		result = append(result, labelModellToLabel(d))
	}
	return result, nil
}

func (s *service) PutMultiLabel(d []Label) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(d) < LimitMaxSum {
		for _, data := range d {
			if len(data.LabelID) == 0 {
				failed = append(failed, data.LabelID)
				continue
			}
			_, err := s.GetLabel(data.LabelID)
			if err != nil {
				failed = append(failed, data.LabelID)
				continue
			}
			err = s.labels.Update(data.LabelID, labelToLabelModel(data))
			if err != nil {
				failed = append(failed, data.LabelID)
				continue
			}
			sucessed = append(sucessed, data.LabelID)
		}
		return sucessed, failed, nil
	}
	return sucessed, failed, ErrExceededMount
}

func (s *service) DeleteMultiLabel(listid []string) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(listid) == 0 {
		return sucessed, failed, ErrInvalidArgument
	}
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			error := s.labels.Remove(id)
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

func labelToLabelModel(d Label) *data.LabelModel {

	return &data.LabelModel{
		UnionID:      d.LabelID,
		LabelType:    d.LabelType,
		SubLableID:   d.SubLableID,
		Action:       d.Action,
		CreateUserID: d.CreateUserID,
		UpdateUserID: d.UpdateUserID,
		Buildin:      d.Buildin,
		Validity:     d.Validity,
		SampleID:     d.SampleID,
		OrderID:      d.OrderID,
		LabelID:      d.LabelID,
		MedicalID:    d.MedicalID,
		UpdateTime:   d.UpdateTime,
		CreateTime:   d.CreateTime,
	}
}

func labelModellToLabel(d *data.LabelModel) Label {
	return Label{
		SampleID:     d.SampleID,
		OrderID:      d.OrderID,
		LabelID:      d.LabelID,
		UpdateTime:   d.UpdateTime,
		CreateTime:   d.CreateTime,
		MedicalID:    d.MedicalID,
		LabelType:    d.LabelType,
		SubLableID:   d.SubLableID,
		Action:       d.Action,
		CreateUserID: d.CreateUserID,
		UpdateUserID: d.UpdateUserID,
		Buildin:      d.Buildin,
		Validity:     d.Validity,
	}
}
