package client

import (
	"encoding/json"
	"net/http"
	"os"
)

type Client struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: "https://canary.discord.com/api/v10",
		Token:   os.Getenv("DISCORD_BOT_TOKEN"),
		Client:  &http.Client{},
	}
}

func (c *Client) Get(path string, result interface{}) error {
	req, err := http.NewRequest("GET", c.BaseURL+path, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bot "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}
