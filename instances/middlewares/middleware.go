package infrastructures

import (
	"fmt"
	"net/http"
	"time"
)

// Middleware doc
type Middleware struct{}

// Apply doc
func (m *Middleware) Apply(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	m.log(r)
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

func (m *Middleware) log(r *http.Request) {
	fmt.Printf("%s %s %s%s\n", time.Now().UTC().Format(time.RFC3339), r.Method, r.Host, r.RequestURI)
}
