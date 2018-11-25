package arvi

import (
	"net/http"

	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

type arviHandler struct {
	Next             httpserver.Handler
	AdditionalHeader string
}

func (h arviHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	// Adding header values as test
	w.Header().Set("Arvi-Ack", h.AdditionalHeader)
	return h.Next.ServeHTTP(w, r)
}

func setup(c *caddy.Controller) error {
	var value string
	for c.Next() {
		if !c.NextArg() {
			return c.ArgErr()
		}
		value = c.Val()
	}
	cfg := httpserver.GetConfig(c)
	mid := func(next httpserver.Handler) httpserver.Handler {
		return arviHandler{Next: next, AdditionalHeader: value}
	}
	cfg.AddMiddleware(mid)
	return nil
}

func init() {
	caddy.RegisterPlugin("arvi", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}
