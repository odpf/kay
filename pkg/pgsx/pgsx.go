package pgsx

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	DB *sqlx.DB
	fs *embed.FS
}

func (c *Client) Close() error {
	return c.DB.Close()
}

// Execs is used for executing list of db query
func (c *Client) Execs(ctx context.Context, queries []string) error {
	for _, query := range queries {
		_, err := c.DB.ExecContext(ctx, query)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Migrate(cfg Config) (err error) {
	iofsDriver, err := iofs.New(c.fs, "migrations")
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", iofsDriver, cfg.ConnURL().String())
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}

func (c *Client) RunWithinTx(ctx context.Context, f func(tx *sqlx.Tx) error) error {
	tx, err := c.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}

	if err := f(tx); err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return fmt.Errorf("rollback transaction error: %v (original error: %w)", txErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

// NewClient initializes database connection
func NewClient(cfg Config, fs *embed.FS) (*Client, error) {
	db, err := sqlx.Connect("postgres", cfg.ConnURL().String())
	if err != nil {
		return nil, fmt.Errorf("error creating and connecting DB: %w", err)
	}
	if db == nil {
		return nil, ErrNilDBClient
	}
	return &Client{db, fs}, nil
}
