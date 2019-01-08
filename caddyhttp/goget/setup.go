package goget

import (
	"github.com/mholt/caddy"
	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("goget", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

//  golang.org/x/nett
//             $1 $2
// goget abc.com github.com/$1/$2
// goget abc.com github.com/twdp/$2
func setup(c *caddy.Controller) error {
	cfg := httpserver.GetConfig(c)

	var rule string
	for c.Next() {
		args := c.RemainingArgs()
		rule = args[0]
	}
	cfg.AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		return Goget{
			Next: next,
			Rule: rule,
		}
	})

	return nil
}