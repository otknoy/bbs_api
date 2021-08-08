package main

import (
	"bbs_api/infra"
	"bbs_api/infra/bbsclient"
	"bbs_api/interfaces"
	"bbs_api/openapi"
	"bbs_api/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	handler := func() http.Handler {
		ub := bbsclient.NewUrlBuilder("5ch.net")

		defaultApiController := openapi.NewDefaultApiController(
			interfaces.NewBbsController(
				service.NewBbsService(
					infra.NewBoardListRepository(ub),
					infra.NewThreadListRepository(ub),
					infra.NewThreadRepository(ub),
				),
			),
		)

		r := openapi.NewRouter(defaultApiController)
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")

				next.ServeHTTP(w, req)
			})
		})

		return r
	}()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: handler,
	}

	log.Println("server start")
	defer log.Println("server stop")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	done := make(chan struct{})
	go func() {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("shutdown error: %v\n", err)
		}

		close(done)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	<-done
}
