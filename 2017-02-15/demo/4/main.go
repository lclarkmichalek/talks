package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
)

var pushPaths = []string{
	"/kn8_files/css",
	"/kn8_files/common.js",
	"/kn8_files/main.js",
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

	err := http.ListenAndServeTLS(":4004", "cert.pem", "key.pem", http.HandlerFunc(handler))
	if err != nil {
		glog.Errorf("failed to listen and serve: %v", err)
		os.Exit(1)
	}
}
