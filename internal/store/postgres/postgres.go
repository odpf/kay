package postgres

import (
	"embed"

	"github.com/odpf/kay/pkg/pgsx"
)

//go:embed migrations/*.sql
var fs embed.FS

func NewClient(cfg pgsx.Config) (*pgsx.Client, error) {
	return pgsx.NewClient(cfg, &fs)
}
