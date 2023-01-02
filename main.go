package main

import (
	"fmt"
	"github.com/pocketbase/pocketbase"
	"log"
	"os"
)

func main() {
	port := os.Getenv("HTTP_PLATFORM_PORT")
	if port == "" {
		port = "8040"
	}
	os.Args = append(os.Args, "serve")
	os.Args = append(os.Args, fmt.Sprint(`--http=127.0.0.1:`, port))

	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
