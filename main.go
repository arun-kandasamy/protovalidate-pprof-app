package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"time"

	pb "protovalidate-pprof-app/gen"

	"github.com/bufbuild/protovalidate-go"
)

func main() {
	// Setup pprof endpoints
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

		log.Println("Starting pprof server at :6060")
		http.ListenAndServe(":6060", mux)
	}()

	// Validator
	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalf("failed to create validator: %v", err)
	}

	// Simulate some validation work
	for {
		msg := &pb.User{
			Name: "Al",
			Age:  -1,
		}

		if err := validator.Validate(msg); err != nil {
			// expect validation error
			log.Println("validation failed:", err)
		} else {
			log.Println("validation passed")
		}

		time.Sleep(100 * time.Millisecond)
	}
}
