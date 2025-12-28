package dto

import (
	"errors"
	"strings"
	"time"
)

type ViewMarketsByRolesRequest struct {
	UserRoles []string
}

type MarketDTO struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Enabled      bool       `json:"enabled"`
	AllowedRoles []string   `json:"allowed_roles"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type ViewMarketsByRolesResponse struct {
	Markets []MarketDTO `json:"markets"`
}

var (
	ErrNoRolesProvided = errors.New("at least one role must be provided")
	ErrEmptyRole       = errors.New("role cannot be empty or whitespace")
)

func (v *ViewMarketsByRolesRequest) Validate() error {
	if len(v.UserRoles) == 0 {
		return ErrNoRolesProvided
	}

	for _, role := range v.UserRoles {
		if strings.TrimSpace(role) == "" {
			return ErrEmptyRole
		}
	}

	return nil
}
