package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"repartners/pkg/calculator"
	"repartners/pkg/storage/memory"
	api "repartners/pkg/transport/http"
)

func main() {
	appCtx := context.Background()

	port := flag.Uint("PORT", 8080, "port for http server")
	flag.Parse()

	if port == nil {
		log.Fatal("unable to determine port")
	}

	packsManager := memory.New[int]()
	calc := calculator.New(packsManager)
	a := api.New(packsManager, calc)

	fmt.Printf("starting http server on port %d\n", *port)
	if err := a.Run(appCtx, *port); err != nil {
		log.Fatalf("http server error: %v", err)
	}
}
