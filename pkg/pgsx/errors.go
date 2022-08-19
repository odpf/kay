package pgsx

import "errors"

var (
	ErrNilDBClient         = errors.New("db client is nil")
	ErrNilPostgresClient   = errors.New("postgres client is nil")
	ErrDuplicateKey        = errors.New("duplicate key")
	ErrCheckViolation      = errors.New("check constraint violation")
	ErrForeignKeyViolation = errors.New("foreign key violation")
)
