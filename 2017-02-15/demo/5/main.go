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
	"/QubitProducts_exporter_exporter_files/octocat-spinner-32.gif",
	"/QubitProducts_exporter_exporter_files/frameworks-b0d504a21e5761da5ec5eade7ed594f51a8583550c5666e515607039056c976d.js",
	"/QubitProducts_exporter_exporter_files/github-db2eee9d5e9a883f2ba4dc3de2cace58297fe2d1f66b91292ecb6a385c12d799.js",
	"/QubitProducts_exporter_exporter_files/github-0e373bf06af78ffa67fcc21199552cb7946a49bf88cbc2e1084257963fc45d1f.css",
	"/QubitProducts_exporter_exporter_files/octocat-spinner-32.gif",
	"/QubitProducts_exporter_exporter_files/frameworks-70aff62372b4dd20e8b7e3141aa52f2b7fda1b9238a597db09f7dd5bbcff25f6.css",
	"/QubitProducts_exporter_exporter_files/78389",
	"/QubitProducts_exporter_exporter_files/octocat-spinner-128.gif",
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
			glog.Infof("not a pusher")
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
