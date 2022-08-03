package main

import (
	"context"
	"fmt"
	"gradebook/log"
	"gradebook/register"
	"gradebook/service"
	stlog "log"
)

func main() {
	log.Run("./app.log")

	host := "localhost"
	port := "4000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r register.Registration
	r.ServiceName = register.LogService
	r.ServiceURL = serviceAddress

	r.RequiredServices = make([]register.ServiceName, 0)
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(context.Background(), r, host, port, log.RegisterHandler)

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down the log service ...")
}