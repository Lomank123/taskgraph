package http

import (
	"fmt"
	"net/http"

	"taskgraph/config"
)

func NewHTTPServer(router http.Handler) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Cfg.App.Port),
		Handler: router,
	}
}
