package main

import (
	"flag"
	"log"

	"github.com/dnataraj/luna/clients/luna_go/api"
)

var (
	addr = flag.String("logger_addr", "localhost:5050", "The address for Luna Logger")
)

func main() {
	flag.Parse()
	c, err := api.New(*addr)
	defer c.Close()
	if err != nil {
		log.Fatal("unable to initialize luna client", err)
	}
	l := log.New(c, "", log.LstdFlags)
	for i := 0; i < 5; i++ {
		l.Printf("counting : %d\n", i)
	}
}
