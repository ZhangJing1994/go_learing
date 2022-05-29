package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("client connect to server"))
	})

	// 单个服务器退出
	serverStop := make(chan struct{})
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		serverStop <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8081",
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit...")
		case <-serverStop:
			log.Println("server will stop")
		}

		CtxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		log.Println("server is stoped")
		return server.Shutdown(CtxTimeout)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return errors.Errorf("get os signal: %v", sig)
		}
	})

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}
