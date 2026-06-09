package proxmox

import (
	"context"
	"net/http"
)

func (c *Client) GetBaseURL() string {
	return c.baseURL
}

func (c *Client) GetHTTPClient() *http.Client {
	return c.httpClient
}

func (c *Client) GetToken() string {
	return c.token
}

func (c *Client) HasSession() bool {
	return c.Session() != nil
}

func (c *Client) RefreshSession(ctx context.Context) error {
	if _, err := c.Version(ctx); err == nil {
		return nil
	}

	if err := c.RefreshTicket(ctx); err == nil {
		return nil
	}

	c.sessionMux.Lock()
	c.session = nil
	c.sessionMux.Unlock()

	return c.CreateSession(ctx)
}

func NewNode(client *Client, name string) *Node {
	return &Node{
		Name:   name,
		client: client,
	}
}
