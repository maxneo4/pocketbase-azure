package main

import (
	"fmt"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
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

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
