package main

import (
	"crypto/tls"
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

	srv := &http.Server{
		Handler:      http.HandlerFunc(handler),
		Addr:         ":4003",
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){},
	}
	err := srv.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil {
		glog.Errorf("failed to listen and serve: %v", err)
		os.Exit(1)
	}
}
