package pprof

import (
	"net/http"
	"net/http/pprof"
	"strings"

	_ "github.com/google/pprof/driver"

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
type Handler struct{
	ServeMux *http.ServeMux
}

// CaddyModule returns the Caddy module information.
func (Handler) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.pprof",
		New: func() caddy.Module { return new(Handler) },
	}
}

// Provision implements caddy.Provisioner.
func (m *Handler) Provision(ctx caddy.Context) error {
	m.ServeMux = http.NewServeMux()
	m.ServeMux.HandleFunc("/debug/pprof/", pprof.Index)
	m.ServeMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	m.ServeMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	m.ServeMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	m.ServeMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) (err error) {
	if strings.HasPrefix(r.URL.Path, "/debug/pprof/") {
		m.ServeMux.ServeHTTP(w, r)
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
