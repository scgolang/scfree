package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/scgolang/sc"
)

func main() {
	var (
		port int
		host string
	)

	flag.IntVar(&port, "p", 57120, "scsynth port")
	flag.StringVar(&host, "h", "127.0.0.1", "scsynth host")

	flag.Parse()

	serverAddr := fmt.Sprintf("%s:%d", host, port)

	client, err := sc.NewClient("udp", "127.0.0.1:57110", serverAddr, 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: free other groups as well?
	if err := client.FreeAll(sc.DefaultGroupID); err != nil {
		log.Fatal(err)
	}
}
