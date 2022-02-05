package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "Mark Git Get Info : ", log.LstdFlags)

	port := ":9000"
	server := newMux(logger, port)

	done := make(chan bool, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go gracefullShutdown(server, logger, interrupt, done)

	// server.SetKeepAlivesEnabled(false)
	logger.Printf("The service is ready to listen and serve at :%s.", port)
	go func() {
		logger.Fatal(server.ListenAndServe())
	}()

	<-done
	logger.Println("Server stopped")
}

func gracefullShutdown(server *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- bool) {
	killSignal := <-quit
	switch killSignal {
	case os.Interrupt:
		logger.Print("Got SIGINT...")
	case syscall.SIGTERM:
		logger.Print("Got SIGTERM...")
	}
	logger.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}

// http://localhost:9000/api/v1/user/name=mark
// /user-lang/name=mark&lang=go
func newMux(logger *log.Logger, port string) *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Println("You're in root path.")
		w.WriteHeader(http.StatusOK)
	})

	api1 := r.PathPrefix("/api/v1").Subrouter()

	api1.HandleFunc("/user", getUserHandler).Queries("name", "{name}").Methods(http.MethodGet)
	api1.HandleFunc("/repo", getUserReposHandler).Queries("name", "{name}").Methods(http.MethodGet)

	return &http.Server{
		Addr:    port,
		Handler: r,
	}
}

// เรียก User
// เรียก ภาษา
// เรียก Follower
// เรียก Repo
