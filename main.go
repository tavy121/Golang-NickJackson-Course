package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	productHandler := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", productHandler)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminat, gracefull shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeoutContext)
}
