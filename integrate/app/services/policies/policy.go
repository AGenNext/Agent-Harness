// Policy Engine - Access, Rate Limiting, Security
package policies

import (
	"context"
	"net/http"
)

// Policy types
type PolicyType string

const (
	PolicyAccess PolicyType = "access"
	PolicyRate PolicyType = "rate"
)

type Policy struct {
	ID     string `json:"id"`
	Name  string `json:"name"`
	Type  PolicyType `json:"type"`
	Effect string `json:"effect"`
}

type Store struct {
	policies map[string]*Policy
}

func NewStore() *Store {
	return &Store{policies: make(map[string]*Policy)}
}

func (s *Store) Add(p *Policy) {
	s.policies[p.ID] = p
}

func (s *Store) List() []*Policy {
	var list []*Policy
	for _, p := range s.policies {
		list = append(list, p)
	}
	return list
}

type Evaluator struct {
	store *Store
}

func NewEvaluator(store *Store) *Evaluator {
	return &Evaluator{store: store}
}

func (e *Evaluator) Evaluate(actor, action string) (bool, string) {
	return true, "allowed"
}