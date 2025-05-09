package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"time"

	pb "protovalidate-pprof-app/gen/example/v1"

	"github.com/bufbuild/protovalidate-go"
)

func main() {
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

	validator, err := protovalidate.New(protovalidate.WithFailFast())
	if err != nil {
		log.Fatalf("failed to create validator: %v", err)
	}

	for {
		msg := &pb.User{
			Name: "Al",
			Age:  1,
		}

		if err := validator.Validate(msg); err != nil {
			log.Println("validation failed:", err)
		} else {
			log.Println("validation passed")
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
