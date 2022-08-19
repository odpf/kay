package cluster

import (
	"context"
	"time"
)

// Cluster is a structure that holds kafka cluster object.
type Cluster struct {
	Urn       string    `json:"urn"`
	Name      string    `json:"name"`
	Servers   string    `json:"servers"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Repository interface {
	Create(ctx context.Context, cluster Cluster) error
	Update(ctx context.Context, cluster Cluster) error
	Get(ctx context.Context, id string) (Cluster, error)
	List(ctx context.Context) ([]Cluster, error)
}
