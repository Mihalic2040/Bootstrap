package app

import (
	"net/http"

	"github.com/Mihalic2040/Bootstrap/src/config"
	"github.com/Mihalic2040/Bootstrap/src/ui"
	hub "github.com/Mihalic2040/Hub"
	"github.com/Mihalic2040/Hub/src/types"
)

var AppConfig types.Config
var App hub.App

func init() {
	AppConfig = config.Read()
}

func Serve() {

	// Metric
	App.Settings(AppConfig)

	App.Start(false)

	http.HandleFunc("/", ui.Render)
	http.HandleFunc("/api/peers", ShowPeers)
	http.HandleFunc("/api/addrs", ShowAddrs)
	http.Handle("/static", ui.Static())

	http.ListenAndServe(":9999", nil)
}
