package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
)

var pushPaths = []string{
	"./kn8_files/main.js",
	"./kn8_files/commons.js",
	"./kn8_files/css",
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		glog.Infof("pushing")
		if p, ok := w.(http.Pusher); ok {
			for _, pth := range pushPaths {
				err := p.Push(pth, nil)
				if err != nil {
					glog.Errorf("could not push: %v", err)
					break
				}
			}
		} else {
			glog.Infof("%T is not a pusher", w)
		}
		time.Sleep(330 * time.Millisecond)
	}

	http.ServeFile(w, r, "./static"+r.URL.Path)
}

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	err := http.ListenAndServeTLS(":4005", "cert.pem", "key.pem",
		prometheus.InstrumentHandler("demo", http.HandlerFunc(handler)))
	if err != nil {
		glog.Errorf("failed to listen and serve: %v", err)
		os.Exit(1)
	}
}
