package main

import (
	"fmt"
	"log"
	"net/http"
	"server/constant"
	util "server/util/log"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			writer.WriteHeader(http.StatusOK)
			fmt.Fprint(writer, "pong")
			util.Log(constant.LOG_INFO, "received a GET request from: (%s)", request.Host)
		default:
			http.NotFound(writer, request)
		}
	})
	s := &http.Server{
		Addr:    ":9777",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}
