package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

var address = "http://localhost:8080"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Starting resquest for '%s'\n", address)
	req, err := http.NewRequestWithContext(ctx, "GET", address, nil)
	if err != nil {
		log.Fatalf("Error creating the request for '%s': %v\n", address, err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error invoking %v\n", err)
	}

	bBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v\n", err)
	}
	log.Printf("Reading response body: <%s>\n", string(bBody))

	defer res.Body.Close()
	defer log.Printf("Finished invoking request for '%s'\n", address)
}
