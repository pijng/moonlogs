package server

import (
	"context"
	"fmt"
	"log"
	"moonlogs/internal/api/server/router"
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func ListenAndServe(cfg *config.Config) error {
	server := createServer(cfg)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	errCh := make(chan error, 1)

	go func() {
		errCh <- server.ListenAndServe()
	}()

	log.Printf("moonlogs is up on %v\n", server.Addr)

	select {
	case <-done:
		return gracefulShutdown(server)
	case err := <-errCh:
		return err
	}
}

func createServer(cfg *config.Config) *http.Server {
	r := mux.NewRouter()
	// r.Use(loggingMiddleware)
	r.Use(corsMiddleware)
	registerRouter(r)

	return &http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.Port),
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
}

func registerRouter(r *mux.Router) {
	session.RegisterSessionStore()

	router.RegisterSetupRouter(r)
	router.RegisterSchemaRouter(r)
	router.RegisterRecordRouter(r)
	router.RegisterUserRouter(r)
	router.RegisterApiTokenRouter(r)
	router.RegisterTagRouter(r)
	router.RegisterSessionRouter(r)
	router.RegisterWebRouter(r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		log.Printf("%s %s\n", r.Method, r.URL.Path)

// 		next.ServeHTTP(w, r)

// 		duration := time.Since(start)
// 		log.Printf("Completed in %v\n", duration)
// 	})
// }

func gracefulShutdown(server *http.Server) error {
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed shutting down: %w", err)
	}

	log.Println("moonlogs stopped")

	return nil
}
