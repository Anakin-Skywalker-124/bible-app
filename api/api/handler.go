package api

import (
	"net/http"
	"Anakin-Skywalker-124/bible-app/pkg/server"
)

func Handler(w http.ResponseWriter,r *http.Request) {
	
	server.ApiRequestHandler(w, r)
}