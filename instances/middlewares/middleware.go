package infrastructures

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Middleware doc
type Middleware struct {
	ctx context.Context
}

// Apply doc
func (m *Middleware) Apply(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	m.log(r)
	m.sanitize()
	m.authenticate()
	m.authorize()
	m.inject()
	next(w, r)
}

func (m *Middleware) sanitize() {
	fmt.Println("sanitizing")
}

func (m *Middleware) authenticate() {
	fmt.Println("authenticating")
}

func (m *Middleware) authorize() {
	fmt.Println("authorizing")
}

func (m *Middleware) inject() {
	fmt.Println("injecting some of common dependency")
}

func (m *Middleware) log(r *http.Request) {
	fmt.Printf("%s %s %s%s\n", time.Now().UTC().Format(time.RFC3339), r.Method, r.Host, r.RequestURI)
}
