package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	svc "go-kit-service-template/pkg/service"

	kitlog "github.com/go-kit/kit/log"
	"github.com/oklog/oklog/pkg/group"
)

func main() {

	httpAddr := ":8080"
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stdout))

	var (
		service     = svc.NewServer()
		endpoints   = svc.MakeEndpoints(service)
		httpHandler = svc.NewHTTPHandler(endpoints)
	)

	var g group.Group
	httpListener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		logger.Log("message", "could not set up HTTP listner", "error", err)
		return
	}
	g.Add(func() error {
		return http.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})

	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})

	logger.Log("HTTP", "listening", "addr", httpAddr)
	g.Run()
}
