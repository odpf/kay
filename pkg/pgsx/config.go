package pgsx

import (
	"net"
	"net/url"
	"strconv"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
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
