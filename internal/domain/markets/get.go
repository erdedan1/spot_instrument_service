package markets

import (
	"time"

	"github.com/gofrs/uuid"
)

func (m *Market) ID() uuid.UUID {
	return m.id
}

func (m *Market) Name() string {
	return m.name
}

func (m *Market) Enabled() bool {
	return m.enabled
}

func (m *Market) AllowedRoles() allowedRoles {
	return m.allowedRoles
}

func (m *Market) CreatedAt() *time.Time {
	return m.createdAt
}

func (m *Market) UpdatedAt() *time.Time {
	return m.updatedAt
}

func (m *Market) DeletedAt() *time.Time {
	return m.deletedAt
}
