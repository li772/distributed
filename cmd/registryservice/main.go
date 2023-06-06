package main

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &registry.RegistryService{})

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	var srv http.Server
	srv.Addr = registry.ServicePort

	go func() {
		log.Println(srv.ListenAndServe())
		cancle()
	}()

	go func() {
		fmt.Println("Registry service started. press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancle()
	}()

	<-ctx.Done()
	fmt.Println("shutting down registry service")
}
