package server

import (
	"context"
	"fmt"
	"log"
	"moonlogs/internal/api/server/router"
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/config"
	"moonlogs/internal/usecases"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func ListenAndServe(cfg *config.Config, uc *usecases.UseCases) error {
	server := createServer(cfg, uc)

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

func createServer(cfg *config.Config, uc *usecases.UseCases) *http.Server {
	r := mux.NewRouter()

	r.Use(corsMiddleware)

	registerRouter(r, uc)

	return &http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.Port),
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}
}

func registerRouter(r *mux.Router, uc *usecases.UseCases) {
	session.RegisterSessionStore(uc.UserUseCase)
	mw := router.InitMiddlewares(uc)

	cfg := &router.SubRouterConfig{
		R:  r,
		MW: mw,
		UC: uc,
	}

	router.RegisterSetupRouter(cfg)
	router.RegisterSchemaRouter(cfg)
	router.RegisterRecordRouter(cfg)
	router.RegisterUserRouter(cfg)
	router.RegisterApiTokenRouter(cfg)
	router.RegisterTagRouter(cfg)
	router.RegisterActionRouter(cfg)
	router.RegisterSessionRouter(cfg)

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
