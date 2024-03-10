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
	"github.com/newrelic/go-agent/v3/newrelic"
)

func ListenAndServe(cfg *config.Config, nrapp *newrelic.Application) error {
	server := createServer(cfg, nrapp)

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

func createServer(cfg *config.Config, nrapp *newrelic.Application) *http.Server {
	r := mux.NewRouter()

	r.Use(corsMiddleware)

	if cfg.NewrelicProfiling {
		r.Use(newrelicMiddleware(nrapp))
	}

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
func newrelicMiddleware(nrapp *newrelic.Application) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			txn := nrapp.StartTransaction(r.URL.Path)
			defer txn.End()

			// Add transaction to request context
			ctx := newrelic.NewContext(r.Context(), txn)

			// Call the next handler, which can be another middleware in the chain or the final handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

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
