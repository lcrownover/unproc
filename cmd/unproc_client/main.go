package main

import (
	"context"
	"flag"
	"fmt"
	//"github.com/lcrownover/unproc/internal/client"
	"log"
	"os"
	"time"

	pb "github.com/lcrownover/unproc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	host := flag.String("host", "", "host to connect to [required]")
	port := flag.String("port", "6661", "port to connect to")
	username := flag.String("username", "", "username to query [required]")

	flag.Parse()

	if *host == "" || *username == "" {
		flag.Usage()
		os.Exit(1)
	}

	addr := fmt.Sprintf("%s:%s", *host, *port)

	log.Printf("querying user %s at %s\n", *username, addr)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUnprocClient(conn)

	user, err := c.Retrieve(ctx, &pb.User{Name: *username})
	if err != nil {
		log.Fatalf("failed to query user: %v", err)
	}

	log.Println(user.ProcessCount)

}
