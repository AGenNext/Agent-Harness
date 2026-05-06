// PII Detection & Removal - LangSmith compatible
package privacy

import (
	"context"
	"regexp"
	"strings"
)

// =============================================
// LangSmith PII Patterns
// https://github.com/langchain-ai/langsmith-pii-removal
// =============================================

type PIIType string

const (
	PIIEmail     PIIType = "email"
	PIIPhone    PIIType = "phone"
	PIIKey      PIIType = "api_key"
	PIIJWT      PIIType = "jwt"
	PIIIP       PIIType = "ip"
	PIIUUID     PIIType = "uuid"
	PIIIPv6     PIIType = "ipv6"
	PIIBank     PIIType = "bank"
	PIIZip      PIIType = "zip"
	PIICredit  PIIType = "credit_card"
	PIISSN     PIIType = "ssn"
)

// LangSmith compatible patterns
var LangSmithPatterns = []PIIPattern{
	// Email
	{Type: PIIEmail, Regex: regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`), Replace: "<|email|>"},
	
	// Phone (US)
	{Type: PIIPhone, Regex: regexp.MustCompile(`\b(\+1|1)?[-.\s]?\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}\b`), Replace: "<|phone|>"},
	
	// API Key (generic)
	{Type: PIIKey, Regex: regexp.MustCompile(`(?i)(api[_-]?|secret[_-]?|access[_-]?)?(key|token|secret)\s*[:=]\s*['"]?([a-zA-Z0-9_-]{16,})['"]?`), Replace: "<|api_key|>"},
	
	// JWT
	{Type: PIIJWT, Regex: regexp.MustCompile(`eyJ[a-zA-Z0-9_-]*\.eyJ[a-zA-Z0-9_-]*\.[a-zA-Z0-9_-]*`), Replace: "<|jwt|>"},
	
	// IPv4
	{Type: PIIIP, Regex: regexp.MustCompile(`\b(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\b`), Replace: "<|ip|>"},
	
	// IPv6
	{Type: PIIIPv6, Regex: regexp.MustCompile(`\b(?:[0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}\b`), Replace: "<|ipv6|>"},
	
	// UUID
	{Type: PIIUUID, Regex: regexp.MustCompile(`\b[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}\b`), Replace: "<|uuid|>"},
	
	// Credit Card
	{Type: PIICredit, Regex: regexp.MustCompile(`\b(?:\d{4}[- ]?){3}\d{4}\b`), Replace: "<|credit_card|>"},
	
	// SSN
	{Type: PIISSN, Regex: regexp.MustCompile(`\b\d{3}[-]?\d{2}[-]?\d{4}\b`), Replace: "<|ssn|>"},
	
	// Bank Account
	{Type: PIIBank, Regex: regexp.MustCompile(`(?i)(account|routing)[- ]?(?:number|#)?\s*[:=]?\s*\b\d{6,17}\b`), Replace: "<|bank|>"},
	
	// Zip Code (US)
	{Type: PIIZip, Regex: regexp.MustCompile(`\b\d{5}(?:-\d{4})?\b`), Replace: "<|zip|>"},
}

type PIIPattern struct {
	Type    PIIType
	Regex  *regexp.Regexp
	Replace string
}

type Scrubber struct {
	patterns []PIIPattern
}

func NewScrubber() *Scrubber {
	return &Scrubber{patterns: LangSmithPatterns}
}

func (s *Scrubber) AddPattern(p PIIPattern) {
	s.patterns = append(s.patterns, p)
}

func (s *Scrubber) Redact(text string) string {
	result := text
	for _, p := range s.patterns {
		result = p.Regex.ReplaceAllString(result, p.Replace)
	}
	return result
}

func Detect(text string) []struct{ Type PIIType; Text string } {
	var findings []struct{ Type PIIType; Text string }
	for _, p := range LangSmithPatterns {
		matches := p.Regex.FindAllString(text, -1)
		for _, m := range matches {
			findings = append(findings, struct{ Type PIIType; Text string }{Type: p.Type, Text: m})
		}
	}
	return findings
}

func AnonymizeUser(id string) string {
	return "user_" + strings.ReplaceAll(id, "a", "x")[:8]
}