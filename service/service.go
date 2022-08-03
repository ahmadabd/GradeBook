package service

import (
	"context"
	"fmt"
	"gradebook/register"
	"log"
	"net/http"
)

func Start(ctx context.Context, reg register.Registration, host string, port string, RegisterHandler func()) (context.Context, error) {
	RegisterHandler()

	ctx = startServer(ctx, reg.ServiceName, host, port)

	err := register.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startServer(ctx context.Context, serviceName register.ServiceName, host string, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	
	var srv http.Server

	srv.Addr = ":" + port

	go func ()  {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func ()  {
		fmt.Printf("%v started. press any key to stop.\n", serviceName)
       	var s string
       	fmt.Scanln(&s)
		err := register.ShutdownService(fmt.Sprintf("http://%v:%v", host, port))
		if err != nil {
			log.Println(err)
		}
       	srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}