package main

import (
	"flag"
	"fmt"
	"os"
)

func runServer(host string, port int) {
	fmt.Printf("listening at %s:%v\n", host, port)
	for {
	}
}

func runClient(host string, port int, user string) {
	fmt.Printf("connecting to %s:%v\n", host, port)
	fmt.Printf("querying user %s\n", user)
}

func main() {
	serverMode := flag.Bool("s", false, "run in server mode")
	clientMode := flag.Bool("c", false, "run in client mode")
	host := flag.String("h", "0.0.0.0", "host to listen on/connect to")
	port := flag.Int("port", 6661, "port to listen on/connect to")
	user := flag.String("user", "", "client mode [required]: user to query")

	flag.Parse()

	if !*serverMode && !*clientMode {
		fmt.Println("You must provide a mode of operation (-c or -s)")
		flag.Usage()
		os.Exit(1)
	}

	if *serverMode && *clientMode {
		fmt.Println("You must only provide a single mode of operation (-c or -s)")
		flag.Usage()
		os.Exit(1)
	}

	if *serverMode {
		runServer(*host, *port)
	}

	if *clientMode {
		if *user == "" {
			fmt.Println("If running in client mode, you must provide a user to query")
			flag.Usage()
			os.Exit(1)
		}
		runClient(*host, *port, *user)
	}

}
