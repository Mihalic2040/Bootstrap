package app

import (
	"context"
	"net/http"

	"github.com/Mihalic2040/Bootstrap/src/config"
	"github.com/Mihalic2040/Bootstrap/src/ui"
	"github.com/Mihalic2040/Hub/src/node"
	"github.com/Mihalic2040/Hub/src/server"
	"github.com/Mihalic2040/Hub/src/types"
)

var AppConfig types.Config
var App types.Host

func init() {
	AppConfig = config.Read()
}

func Serve() {
	ctx := context.Background()

	handlers := server.HandlerMap{}

	// Metric

	App = node.Server(ctx, handlers, AppConfig, false)

	http.HandleFunc("/", ui.Render)
	http.HandleFunc("/api/peers", ShowPeers)
	http.Handle("/static", ui.Static())

	http.ListenAndServe(":9999", nil)
}
