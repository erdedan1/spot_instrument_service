package markets

import (
	"context"
	"spot_instrument_service/internal/domain/markets"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
	return &PostgresRepo{
		db: db,
	}
}

func (r *PostgresRepo) ViewMarketsByRoles(ctx context.Context, userRoles []markets.Role) ([]*markets.Market, error) {
	return nil, nil
}

func (r *PostgresRepo) Delete(id uuid.UUID) error {
	return nil
}

func (r *PostgresRepo) Save(m *markets.Market) (*markets.Market, error) {
	return m, nil
}
