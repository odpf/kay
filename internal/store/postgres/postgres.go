package postgres

import (
	"context"
	"embed"

	"github.com/jmoiron/sqlx"
	"github.com/odpf/kay/pkg/pgsx"
)

// schema represents the storage schema.
// Note: Update the constants above if the table name is changed.
//
//go:embed schema.sql
var schema string

//go:embed migrations/*.sql
var fs embed.FS

func NewClient(cfg pgsx.Config) (*pgsx.Client, error) {
	return pgsx.NewClient(cfg, &fs)
}

type Store struct {
	db *sqlx.DB
}

func (st *Store) Bootstrap(ctx context.Context) error {
	_, err := st.db.ExecContext(ctx, schema)
	return err
}

func NewStore(cfg pgsx.Config) (*Store, error) {
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Store{db: client.DB}, nil
}
