package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// =============================================
// Webhook Listener
// =============================================

type WebhookEvent struct {
	Source  string      `json:"source"`
	Type    string      `json:"type"`
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

type HandlerFunc func(ctx context.Context, event *WebhookEvent) error

type WebhookHandler struct {
	handlers map[string]HandlerFunc
}

func NewWebhookHandler() *WebhookHandler {
	h := &WebhookHandler{handlers: make(map[string]HandlerFunc)}
	h.registerDefaults()
	return h
}

func (h *WebhookHandler) registerDefaults() {
	h.Register("github.issue", h.handleGitHubIssue)
	h.Register("github.pr", h.handleGitHubPR)
	h.Register("github.push", h.handleGitHubPush)
	h.Register("linear.issue", h.handleLinearIssue)
	h.Register("slack.command", h.handleSlackCommand)
}

func (h *WebhookHandler) Register(eventType string, fn HandlerFunc) {
	h.handlers[eventType] = fn
}

func (h *WebhookHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	source := r.Header.Get("X-Webhook-Source")
	eventType := r.Header.Get("X-Webhook-Type")

	var payload interface{}
	json.NewDecoder(r.Body).Decode(&payload)

	event := &WebhookEvent{Source: source, Type: eventType, Payload: payload}

	key := fmt.Sprintf("%s.%s", source, eventType)
	if handler, ok := h.handlers[key]; ok {
		return handler(context.Background(), event)
	}
	return nil
}

// GitHub handlers
func (h *WebhookHandler) handleGitHubIssue(ctx context.Context, e *WebhookEvent) error { return nil }
func (h *WebhookHandler) handleGitHubPR(ctx context.Context, e *WebhookEvent) error { return nil }
func (h *WebhookHandler) handleGitHubPush(ctx context.Context, e *WebhookEvent) error { return nil }

// External handlers
func (h *WebhookHandler) handleLinearIssue(ctx context.Context, e *WebhookEvent) error { return nil }
func (h *WebhookHandler) handleSlackCommand(ctx context.Context, e *WebhookEvent) error { return nil }

/*
# Endpoints:
POST /webhooks/github
POST /webhooks/linear  
POST /webhooks/jira
POST /webhooks/slack

# GitHub payload URL:
https://server.com/webhooks/github
*/