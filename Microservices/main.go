package main

import (
	"context"
	"log"
	"net/http"
	"github.com/java-samurai/Go-Learn/Microservices/handlers"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, " product-api ", log.LstdFlags)
	hh := handlers.NewHello(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh) //Serve mux can handle multiple routes
	s := &http.Server{
		Addr: ":9090",
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
		IdleTimeout: 10*time.Second,
		Handler: sm,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <- sigChannel
	l.Println("Received Interrupt from OS. Performing Graceful Shutdown | ", sig)
	tc, _ := context.WithTimeout(context.Background(), 40*time.Second)
	s.Shutdown(tc)

}
