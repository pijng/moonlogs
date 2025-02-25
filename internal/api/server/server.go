package server

import (
	"context"
	"fmt"
	"log"
	"moonlogs/internal/api/server/router"
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/usecases"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func ListenAndServe(uc *usecases.UseCases, geminiToken string, opts ...SrvOpt) error {
	server := createServer(uc, geminiToken, opts...)

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

type SrvOpt func(*http.Server)

func WithPort(port int) SrvOpt {
	return func(s *http.Server) {
		s.Addr = fmt.Sprintf(":%v", port)
	}
}

func WithReadTimeout(readTimeout time.Duration) SrvOpt {
	return func(s *http.Server) {
		s.ReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) SrvOpt {
	return func(s *http.Server) {
		s.WriteTimeout = writeTimeout
	}
}

func createServer(uc *usecases.UseCases, geminiToken string, opts ...SrvOpt) *http.Server {
	r := mux.NewRouter()

	r.Use(corsMiddleware)

	registerRouter(r, uc, geminiToken)

	server := &http.Server{
		Handler: r,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func registerRouter(r *mux.Router, uc *usecases.UseCases, geminiToken string) {
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
	router.RegisterAlertingRuleRouter(cfg)
	router.RegisterIncidentsRouter(cfg)
	router.RegisterNotificationProfileRouter(cfg)
	router.RegisterSessionRouter(cfg, geminiToken)

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
