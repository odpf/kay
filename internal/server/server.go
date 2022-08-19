package server

import (
	"fmt"

	"github.com/odpf/kay/config"
	"github.com/odpf/kay/internal/store/postgres"
)

func Start(cfg *config.Config) error {
	db, err := postgres.NewClient(cfg.DB)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println("closing db connection...")
		db.Close()
	}()
	return nil
}
