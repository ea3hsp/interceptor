package main

import (
	"log"
	"net/http"
	"os"

	"bitbucket.org/celsa/interceptor/internal/interceptorserver"
	"bitbucket.org/celsa/interceptor/rpc/interceptor"
	"github.com/joho/godotenv"
	"github.com/twitchtv/twirp/hooks/statsd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	hook := statsd.NewStatsdServerHooks(interceptorserver.LoggingStatter{os.Stderr})
	server := interceptor.NewInterceptorServer(&interceptorserver.RandomInterceptor{}, hook)
	log.Fatal(http.ListenAndServe(":8080", server))
}
