package pgsx

import (
	"net"
	"net/url"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}

// ConnURL returns the connection URL for the Postgres database.
func (c *Config) ConnURL() *url.URL {
	pgURL := &url.URL{
		Scheme: "postgres",
		Host:   net.JoinHostPort(c.Host, strconv.Itoa(c.Port)),
		User:   url.UserPassword(c.User, c.Password),
		Path:   c.Name,
	}
	q := pgURL.Query()
	q.Add("sslmode", "disable")
	pgURL.RawQuery = q.Encode()

	return pgURL
}
