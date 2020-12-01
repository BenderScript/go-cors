package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitRouterA() {

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("This is Mux Router A"))
	}).Methods("GET", "POST")

	server := &http.Server{
		Handler:      handlers.CORS()(newRouter),
		Addr:         "0.0.0.0:12703",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("==== Starting Server at: %s =====\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v", server.Addr, err)
		}
	}()

	go func() {
		// A interrupt signal
		<-quit
		log.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v", err)
		}
		// We are done
		close(done)
	}()

	<-done
	log.Println("Server stopped")

}
