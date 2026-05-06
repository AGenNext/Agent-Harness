package copilot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Microsoft Copilot / Teams Integration
type Config struct {
	TenantID     string
	ClientID    string
	ClientSecret string
}

type Activity struct {
	Type string `json:"type"`
	From    `json:"from"`
	Text    string `json:"text"`
}

type ChannelAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Client struct {
	config *Config
	agent interface{ Execute(ctx context.Context, agent, input string) (string, error) }
}

func NewClient(cfg *Config, agent interface{ Execute(ctx context.Context, agent, input string) (string, error) }) *Client {
	return &Client{config: cfg, agent: agent}
}

func (c *Client) HandleMessage(w http.ResponseWriter, r *http.Request) error {
	var activity Activity
	json.NewDecoder(r.Body).Decode(&activity)

	resp, _ := c.agent.Execute(r.Context(), "code-assist", activity.Text)

	return json.NewEncoder(w).Encode(Activity{
		Type: "message",
		From: ChannelAccount{ID: "agent-harness", Name: "Agent Harness"},
		Text: resp,
	})
}

/*
# Microsoft Copilot Integration

# Teams: @agent fix, @agent review
# .env: COPILOT_TENANT_ID, COPILOT_CLIENT_ID
# Endpoint: POST /copilot/message
*/