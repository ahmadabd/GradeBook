package main

import (
	"context"
	"fmt"
	"gradebook/register"
	"log"
	"net/http"
)

func main() {
	http.Handle("/services", &register.RegisteryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = register.ServicePort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("Registery service started. press any key to stop.\n")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down the registery service ...")
}
