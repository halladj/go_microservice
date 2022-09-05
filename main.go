package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"practice/handlers"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.HandleFunc("/", hh.ServeHTTP)
	sm.HandleFunc("/goodbye", gh.ServeHTTP)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt)
	signal.Notify(sc, os.Kill)

	sig := <-sc
	l.Println("Recived terminate, gracefull shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
