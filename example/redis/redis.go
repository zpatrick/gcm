package redis

type ClientConfig struct {
	Host string
	Port int
}

type Client struct{}

func NewClient(c ClientConfig) *Client {
	return &Client{}
}

func (c *Client) Ping() error {
	return nil
}
