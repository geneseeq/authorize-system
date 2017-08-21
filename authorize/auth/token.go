// Package data contains the heart of the domain model.
package auth

import (
	"errors"
	"time"
)

// // LabelIDModel is
// type LabelIDModel struct {
// 	ID      string   `json:"id"`
// 	LabelID []string `json:"label_id"`
// }

// TokenModel is user struct
type TokenModel struct {
	ID           string
	AccessToken  string
	Validity     string
	TokenType    string
	ExpiresIn    int
	RefreshToken string
	UpdateTime   time.Time
}

// NewToken is create instance
func NewToken(id string) *TokenModel {
	return &TokenModel{
		ID: id,
	}
}

// ErrUnknown is unkown user error
var (
	ErrUnknown = errors.New("unknown user")
)

// TimeUtcToCst is format time
func TimeUtcToCst(t time.Time) time.Time {
	return t.Add(time.Hour * time.Duration(8))
}

// TokenRepository is user interface
type TokenRepository interface {
	Store(data *TokenModel) error
	Find(id string) (*TokenModel, error)
	FindAllToken() []*TokenModel
	Remove(id string) error
	Update(id string, data *TokenModel) error
}
