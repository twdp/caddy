package goget

import (
	"fmt"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"net/http"
	"strings"
)

type Goget struct {
	Next     httpserver.Handler
	Rule string
}

func (get Goget) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	query := r.URL.Query()
	// 如果是go-get=1
	//
	if query["go-get"][0] == "1" {
		path := r.URL.Path
		domain := r.URL.Hostname()
		fmt.Println(path)
		fmt.Println(r.Host)
		fmt.Println(domain)
		paths := strings.Split(path, "/")
		git := get.Rule
		if strings.Contains(git, "$1") && len(paths) > 1 {
			git = strings.Replace(get.Rule, "$1", paths[1], 1)
		}
		if strings.Contains(git, "$2") && len(paths) > 2 {
			git = strings.Replace(get.Rule, "$2", paths[2], 1)
		}
		body := fmt.Sprintf("<!DOCTYPE html><html><head><meta content='%s%s git %s.git' name='go-import'></head></html>", domain, path, git)
		w.Write([]byte(body))
		return http.StatusOK, nil
	} else {
		return get.Next.ServeHTTP(w, r)
	}
}