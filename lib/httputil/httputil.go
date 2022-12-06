package httputil

import (
	"flag"
	"github.com/jatis/oms/lib/log"
	"gopkg.in/tylerb/graceful.v1"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var listenPort string

func init() {
	flag.StringVar(&listenPort, "p", "", "listener port")
}

func HTTPServe(port string, handler http.Handler, gracefulTimeout, readTimeout, writeTimeout time.Duration) error {

	l, err := httpListen(port)
	if err != nil {
		log.Fatalln(err)
	}

	srv := &graceful.Server{
		Timeout: gracefulTimeout,
		Server: &http.Server{
			Handler:      handler,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	}

	log.Println("starting serve on ", port)
	return srv.Serve(l)
}

// This method can be used for any TCP Listener, e.g. non HTTP
func httpListen(hport string) (net.Listener, error) {
	var l net.Listener

	fd := os.Getenv("EINHORN_FDS")
	if fd != "" {
		sock, err := strconv.Atoi(fd)
		if err == nil {
			hport = "socketmaster:" + fd
			log.Println("detected socketmaster, listening on", fd)
			file := os.NewFile(uintptr(sock), "listener")
			fl, err := net.FileListener(file)
			if err == nil {
				l = fl
			}
		}
	}

	if listenPort != "" {
		hport = ":" + listenPort
	}

	if l == nil {
		var err error
		l, err = net.Listen("tcp4", hport)
		if err != nil {
			return nil, err
		}
	}

	return l, nil
}
