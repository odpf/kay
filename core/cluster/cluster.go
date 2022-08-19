package cluster

import "time"

// Cluster is a structure that holds kafka cluster object.
type Cluster struct {
	Urn       string    `json:"urn"`
	Name      string    `json:"name"`
	Servers   string    `json:"servers"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
