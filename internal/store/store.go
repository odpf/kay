package store

import "github.com/odpf/kay/core/cluster"

type ClusterRepository interface {
	Create(*cluster.Cluster) error
	Update(*cluster.Cluster) error
	Get(id string) (*cluster.Cluster, error)
	List() ([]*cluster.Cluster, error)
}
