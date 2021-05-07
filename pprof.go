package pprof

import (
	"net/http"
	_ "net/http/pprof"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Handler{})
	httpcaddyfile.RegisterHandlerDirective("pprof", parseCaddyfile)
}

// Handler implements an HTTP handler that ...
type Handler struct{}

// CaddyModule returns the Caddy module information.
func (Handler) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.pprof",
		New: func() caddy.Module { return new(Handler) },
	}
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) (err error) {
	if strings.HasPrefix(r.URL.Path, "/debug/pprof/") {
		http.DefaultServeMux.ServeHTTP(w, r)
		return
	}
	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile unmarshals Caddyfile tokens into h.
func (h *Handler) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	return nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	return &Handler{}, nil
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*Handler)(nil)
	_ caddyfile.Unmarshaler       = (*Handler)(nil)
)
