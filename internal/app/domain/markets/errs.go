package markets

import "errors"

var (
	ErrEmptyName      = errors.New("market name is empty")
	ErrEmptyID        = errors.New("market id is empty")
	ErrInvalidRole    = errors.New("invalid role")
	ErrEnabledNoRoles = errors.New("enabled market must have roles")
)
