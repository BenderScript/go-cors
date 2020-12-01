package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func InitRouterB() {

	newRouter := mux.NewRouter()

	newRouter.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("This is Mux Router B"))
	}).Methods("GET", "POST")

	server := &http.Server{
		Handler:      handlers.CORS()(newRouter),
		Addr:         "0.0.0.0:12704",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("==== Starting Server at: %s =====\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v", server.Addr, err)
		}
	}()
}
