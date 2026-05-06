// Guardrail - Content Safety
package guardrail

import "regexp"

type RuleType string

const (
	RuleBlock RuleType = "block"
	RuleFilter RuleType = "filter"
)

type Rule struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type RuleType `json:"type"`
	Pattern string `json:"pattern"`
}

var DefaultRules = []Rule{
	{ID: "sudo", Type: RuleBlock, Pattern: `sudo|su|root`},
	{ID: "exec", Type: RuleBlock, Pattern: `exec\(|system\(|eval\(`},
	{ID: "sql", Type: RuleBlock, Pattern: `(?i)(union|select).*from.*`},
	{ID: "xss", Type: RuleBlock, Pattern: `<script|javascript:`},
	{ID: "pwd", Type: RuleFilter, Pattern: `(?i)password`},
}

type Guardrail struct {
	rules []Rule
}

func New() *Guardrail { return &Guardrail{rules: DefaultRules} }

func (g *Guardrail) Check(input string) []string {
	var violations []string
	for _, r := range g.rules {
		if regexp.MustCompile(r.Pattern).MatchString(input) {
			violations = append(violations, r.Name)
		}
	}
	return violations
}

func (g *Guardrail) Allowed(input string) bool {
	return len(g.Check(input)) == 0
}