package redis

import "log"

type ClientConfig struct {
	Host string
	Port int
}

type Client struct {
	Host string
	Port int
}

func NewClient(c ClientConfig) *Client {
	return &Client{
		Host: c.Host,
		Port: c.Port,
	}
}

func (c *Client) Ping() error {
	log.Printf("Fake pinging redis at %s:%d\n", c.Host, c.Port)
	return nil
}
