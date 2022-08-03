package main

import (
	"context"
	"fmt"
	"gradebook/grades"
	"gradebook/log"
	"gradebook/register"
	"gradebook/service"
	stlog "log"
)

func main() {
	host, port := "localhost", "6000"	

	serviceName := fmt.Sprintf("http://%v:%v", host, port)

	var r register.Registration
	r.ServiceName = register.GradingService
	r.ServiceURL = serviceName
	r.RequiredServices = []register.ServiceName{register.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(context.Background(), r, host, port, grades.RegisterHandler)

	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := register.GetProvider(register.LogService); err == nil {
		fmt.Printf("Logging service found at %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service ...")
}