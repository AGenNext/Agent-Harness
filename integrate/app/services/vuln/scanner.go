// Vulnerability Scanner
package vuln

import "regexp"

type Vulnerability struct {
	ID string `json:"id"`
	Severity string `json:"severity"`
	Title string `json:"title"`
	CWE string `json:"cwe"`
}

var DefaultVulns = []Vulnerability{
	{ID: "CWE-78", Severity: "high", Title: "Command Injection", CWE: "CWE-78"},
	{ID: "CWE-89", Severity: "high", Title: "SQL Injection", CWE: "CWE-89"},
	{ID: "CWE-79", Severity: "medium", Title: "XSS", CWE: "CWE-79"},
	{ID: "CWE-798", Severity: "high", Title: "Hardcoded Credentials", CWE: "CWE-798"},
	{ID: "CWE-22", Severity: "medium", Title: "Path Traversal", CWE: "CWE-22"},
}

type Scanner struct {
	patterns []Vulnerability
}

func New() *Scanner { return &Scanner{patterns: DefaultVulns} }

func (s *Scanner) Scan(code string) []Vulnerability {
	var found []Vulnerability
	patterns := map[string]Vulnerability{
		`exec\(|system\(|popen\(`:      DefaultVulns[0],
		`(?i)(select|insert).*from`:  DefaultVulns[1],
		`innerHTML\s*=|document\.write`: DefaultVulns[2],
		`(?i)password\s*=\s*['"]`:    DefaultVulns[3],
	}

	for pattern, vuln := range patterns {
		if regexp.MustCompile(pattern).MatchString(code) {
			found = append(found, vuln)
		}
	}
	return found
}