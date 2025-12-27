package markets

import (
	"strings"
	"time"
)

func (m *Market) Enable() error {
	if len(m.allowedRoles) == 0 {
		return ErrEnabledNoRoles
	}
	m.enabled = true
	now := time.Now()
	m.updatedAt = &now
	return nil
}

func (m *Market) Disable() {
	m.enabled = false
	now := time.Now()
	m.updatedAt = &now
}

func (m *Market) UpdateName(newName string) error {
	if strings.TrimSpace(newName) == "" {
		return ErrEmptyName
	}
	m.name = newName
	now := time.Now()
	m.updatedAt = &now
	return nil
}

func (m *Market) AddRole(r Role) error {
	if !r.isValid() {
		return ErrInvalidRole
	}
	m.allowedRoles = append(m.allowedRoles, r)
	now := time.Now()
	m.updatedAt = &now
	return nil
}

func (m *Market) RemoveRole(r Role) {
	newRoles := allowedRoles{}
	for _, role := range m.allowedRoles {
		if role != r {
			newRoles = append(newRoles, role)
		}
	}
	m.allowedRoles = newRoles
	now := time.Now()
	m.updatedAt = &now
}

func (m *Market) MarkDeleted() {
	now := time.Now()
	m.deletedAt = &now
}

func (m *Market) MarkUpdated() {
	now := time.Now()
	m.updatedAt = &now
}
