package msa

import "time"

const headerAuthorization = "Authorization"

var DefaultClientConfiguration = ClientConfiguration{
	Timeout:        20 * time.Second,
	ConnectionPool: 10,
	Debug:          false,
}

type ClientConfiguration struct {
	Timeout        time.Duration
	ConnectionPool int
	Headers        map[string]string
	Debug          bool
}

func (c *ClientConfiguration) SetToken(token string) {
	if c.Headers == nil {
		c.Headers = map[string]string{headerAuthorization: token}
		return
	}
	c.Headers[headerAuthorization] = token
}
