package handler

import (
	"fmt"
	"net/http"
	"server/constant"
	"server/util/log"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprint(w, "pong")
		log.Log(constant.LOG_INFO, "received a GET request from: (%s)", r.Host)
	default:
		http.NotFound(w, r)
	}
}
