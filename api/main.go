package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	//"github.com/efedyakov/project2/internal/logger"
)

const appName = "bannerrotation"

func main() {
	ctx := context.Background()
	//defer logger.Logger().Sync()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()

	// Use the http.RedirectHandler() function to create a handler which 307
	// redirects all requests it receives to http://example.org.
	rh := http.RedirectHandler("http://example.org", 307)

	// Next we use the mux.Handle() function to register this with our new
	// servemux, so it acts as the handler for all incoming requests with the URL
	// path /foo.
	mux.Handle("/foo", rh)

	log.Print("Listening...")

	// Then we create a new server and start listening for incoming requests
	// with the http.ListenAndServe() function, passing in our servemux for it to
	// match requests against as the second argument.
	http.ListenAndServe(":3000", mux)

}
