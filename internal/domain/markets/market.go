package markets

import (
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type Market struct {
	id           uuid.UUID
	name         string
	enabled      bool
	allowedRoles allowedRoles
	createdAt    *time.Time
	updatedAt    *time.Time
	deletedAt    *time.Time
}

func NewMarket(id uuid.UUID, name string, enabled bool, roles allowedRoles) (*Market, error) {
	if id == uuid.Nil {
		return nil, ErrEmptyID
	}

	if strings.TrimSpace(name) == "" {
		return nil, ErrEmptyName
	}

	roleSet := make(map[Role]struct{})
	for _, r := range roles {
		if !r.isValid() {
			return nil, ErrInvalidRole
		}
		roleSet[r] = struct{}{}
	}

	if enabled && len(roleSet) == 0 {
		return nil, ErrEnabledNoRoles
	}

	time := time.Now()

	return &Market{
		id:           id,
		name:         name,
		enabled:      enabled,
		allowedRoles: roles,
		createdAt:    &time,
	}, nil
}
