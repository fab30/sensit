package server

import "net/http"
import "github.com/joelvim/sensit/authtoken"

// AuthHandler is an interface to abstract different kind of authentication
// such as Basic or Bearer token
type AuthHandler interface {
	Handle(handler requestHandler) requestHandler
}

// BasicAuthHandler handle Basic authentication
type BasicAuthHandler struct {
	login    string
	password string
}

// Handle basic authentication
func (b BasicAuthHandler) Handle(handler requestHandler) requestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		if l, p, ok := r.BasicAuth(); ok && l == b.login && p == b.password {
			handler(w, r)
		} else {
			// Unauthorized, return a response with a header WWW-Authenticate: Basic realm="sensit receiver"
			w.Header().Add("WWW-Authenticate", "Basic realm=\"sensit receiver\"")
			http.Error(w, "Authentication required", http.StatusUnauthorized)
		}
	}
}

// TokenAuthHandler handle authentication by access_token
type TokenAuthHandler struct {
	login    string
	password string
	salt     string
}

// Handle authentication by access_token
func (t TokenAuthHandler) Handle(handler requestHandler) requestHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		token := authtoken.Token(t.login, t.password, t.salt)

		if r.URL.Query().Get("access_token") == token {
			handler(w, r)
		} else {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
		}
	}
}
