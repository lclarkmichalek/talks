package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/golang/glog"

	"github.com/pressly/chi"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok {
		err := p.Push("/img/foo.png", nil)
		if err != nil {
			glog.Errorf("could not push /img/foo.jpg: %v", err)
		}
	} else {
		glog.Infof("not a pusher")
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<html><img src="/img/foo.png"/></html>`))
}

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	mux := chi.NewRouter()
	mux.Get("/", handler)
	mux.Get("/img/foo.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./img/foo.png")
	})

	err := http.ListenAndServeTLS(":4002", "cert.pem", "key.pem", mux)
	if err != nil {
		glog.Errorf("failed to listen and serve: %v", err)
		os.Exit(1)
	}
}
