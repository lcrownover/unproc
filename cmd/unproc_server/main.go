package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lcrownover/unproc/internal/server"
)

func main() {
	host := flag.String("h", "0.0.0.0", "host to listen on")
	port := flag.String("port", "6661", "port to listen on")

	flag.Parse()

	lis := *server.NewTCPListener(*host, *port)
	srv := *(server.NewGRPCServer())

	fmt.Printf("listening on %s:%s\n", *host, *port)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
