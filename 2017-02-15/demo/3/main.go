package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		glog.Infof("Sleeping for 330ms")
		time.Sleep(330 * time.Millisecond)
	}

	http.ServeFile(w, r, "./static"+r.URL.Path)
}

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	err := http.ListenAndServe(":4003", http.HandlerFunc(handler))
	if err != nil {
		glog.Errorf("failed to listen and serve: %v", err)
		os.Exit(1)
	}
}
