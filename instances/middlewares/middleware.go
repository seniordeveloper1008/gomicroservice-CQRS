package infrastructures

import (
	"fmt"
	"net/http"
)

// Middleware doc
type Middleware struct{}

// Apply doc
func (m *Middleware) Apply(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	m.sanitize()
	m.authenticate()
	m.authorize()
	next(w, r)
}

func (m *Middleware) sanitize() {
	fmt.Println("sanitize")
}
func (m *Middleware) authenticate() {
	fmt.Println("authenticate")
}
func (m *Middleware) authorize() {
	fmt.Println("authorize")
}
