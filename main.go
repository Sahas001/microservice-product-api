package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Sahas001/go-micro/handlers"
)

func main() {
	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()

	sm.Handle("/", ph)

	server := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	go func() {
		l.Println("Starting server on port 9090")
		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error Starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until a signal is received

	sig := <-c
	log.Println("Got Signal:", sig)

	// gracefully shutdown the server after waiting max 30 seconds for current operation to complete

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
