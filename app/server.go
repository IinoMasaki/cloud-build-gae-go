package app

import (
	"net/http"

	"github.com/IinoMasaki/cloud-build-gae-go/src/handler"
	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handler.Hello)
	appengine.Main()
}
