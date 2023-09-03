package server

import (
	"context"
	"fmt"
	"log"
	"moonlogs/api/server/router"
	"moonlogs/api/server/session"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func Serve() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := createServer()
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	log.Printf("moonlogs is up on %v\n", server.Addr)

	<-done
	gracefulShutdown(server)
}

func createServer() *http.Server {
	r := mux.NewRouter()
	registerRouter(r)

	r.Use(loggingMiddleware)

	return &http.Server{
		Addr:         "0.0.0.0:4200",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
}

func registerRouter(r *mux.Router) {
	store := session.RegisterSessionStore()

	router.RegisterLogSchemaRouter(r, store)
	router.RegisterLogRecordRouter(r, store)
	router.RegisterUserRouter(r, store)
	router.RegisterSessionRouter(r)
	router.RegisterWebRouter(r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("%s %s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("Completed in %v\n", duration)
	})
}

func gracefulShutdown(server *http.Server) {
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Errorf("Balansir shutdown failed: %w", err))
	}
	log.Println("moonlogs stopped")
}
