package session

import "github.com/labstack/echo/v4"

type MiddlewareOptions struct {
	Path        string
	Domain      string
	MaxAge      int
	Secure      bool
	HttpOnly    bool
	SessionName string
	SessionKey  string
}

// NewMiddleware new session middleware
func NewMiddleware(o MiddlewareOptions) echo.MiddlewareFunc {
	store := NewCookieStore([]byte(o.SessionKey))
	store.Options(Options{
		Path:     o.Path,
		Domain:   o.Domain,
		MaxAge:   o.MaxAge,
		HttpOnly: o.HttpOnly,
		Secure:   o.Secure,
	})

	return Sessions(o.SessionName, store)
}
