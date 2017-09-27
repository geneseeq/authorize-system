// Package authing provides the use-case of booking a cargo. Used by views
// facing an administrator.
package authing

import (
	"errors"
	"time"

	"github.com/geneseeq/authorize-system/authorize/auth"
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
	GetToken(id string) (Token, error)
	PostToken(token []Token) ([]string, []string, error)
	GetAllToken() ([]Token, error)
	DeleteMultiToken(listid []string) ([]string, []string, error)
	PutMultiToken(token []Token) ([]string, []string, error)
}

// Token is a basedata base info
type Token struct {
	ID           string    `json:"id"`
	AccessToken  string    `json:"access_token"`
	Validity     bool      `json:"validity"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
	UpdateTime   time.Time `json:"update_time"`
	CreateTime   time.Time `json:"create_time"`
}

type service struct {
	tokens auth.TokenRepository
}

// NewService creates a booking service with necessary dependencies.
func NewService(tokens auth.TokenRepository) Service {
	return &service{
		tokens: tokens,
	}
}

func (s *service) PostToken(d []Token) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(d) < LimitMaxSum {
		for _, content := range d {
			currentTime := auth.TimeUtcToCst(time.Now())
			content.UpdateTime = currentTime
			content.CreateTime = currentTime
			err := s.tokens.Store(tokenToTokenModel(content))
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

func (s *service) GetToken(id string) (Token, error) {
	if id == "" {
		return Token{}, ErrInvalidArgument
	}
	d, error := s.tokens.Find(id)
	if error != nil {
		return Token{}, ErrNotFound
	}
	return tokenModellToToken(d), nil
}

func (s *service) GetAllToken() ([]Token, error) {
	var result []Token
	for _, d := range s.tokens.FindAllToken() {
		result = append(result, tokenModellToToken(d))
	}
	return result, nil
}

func (s *service) PutMultiToken(d []Token) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(d) < LimitMaxSum {
		for _, data := range d {
			if len(data.ID) == 0 {
				return nil, nil, ErrInvalidArgument
			}
			_, err := s.GetToken(data.ID)
			if err != nil {
				failed = append(failed, data.ID)
				continue
			}
			err = s.tokens.Update(data.ID, tokenToTokenModel(data))
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

func (s *service) DeleteMultiToken(listid []string) ([]string, []string, error) {
	var sucessed []string
	var failed []string
	if len(listid) < LimitMaxSum {
		for _, id := range listid {
			if id == "" {
				return sucessed, failed, ErrInvalidArgument
			}
			error := s.tokens.Remove(id)
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

func tokenToTokenModel(d Token) *auth.TokenModel {

	return &auth.TokenModel{
		UnionID:      d.ID,
		ID:           d.ID,
		AccessToken:  d.AccessToken,
		Validity:     d.Validity,
		TokenType:    d.TokenType,
		ExpiresIn:    d.ExpiresIn,
		RefreshToken: d.RefreshToken,
		UpdateTime:   d.UpdateTime,
		CreateTime:   d.CreateTime,
	}
}

func tokenModellToToken(d *auth.TokenModel) Token {
	return Token{
		ID:           d.ID,
		AccessToken:  d.AccessToken,
		Validity:     d.Validity,
		TokenType:    d.TokenType,
		ExpiresIn:    d.ExpiresIn,
		RefreshToken: d.RefreshToken,
		UpdateTime:   d.UpdateTime,
		CreateTime:   d.CreateTime,
	}
}
