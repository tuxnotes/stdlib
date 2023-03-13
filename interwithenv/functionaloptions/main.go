package main

import (
	"fmt"
	"os"
)

type Client struct {
	ConsulIP   string
	ConnString string
}

func (c *Client) String() string {
	return fmt.Sprintf("consulIP: %s, connString: %s", c.ConsulIP, c.ConnString)
}

type Option func(*Client)

var defaultClient = Client{
	ConsulIP:   "localhost",
	ConnString: "postgresql://localhost:5432",
}

func FromEnv() Option {
	return func(c *Client) {
		connstring, exist := os.LookupEnv("DB_CONN")
		if exist {
			c.ConnString = connstring
		}
	}
}

func NewClient(opts ...Option) *Client {
	var client = defaultClient
	for _, opt := range opts {
		opt(&client)
	}
	return &client
}

func main() {
	client := NewClient()
	fmt.Println(client.ConsulIP)
	fmt.Println(client.ConnString)
}
