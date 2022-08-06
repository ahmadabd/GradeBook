package main

import (
	"context"
	"fmt"
	"gradebook/log"
	"gradebook/register"
	"gradebook/service"
	"gradebook/teacherportal"
	stlog "log"
)

func main() {
	err := teacherportal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}

	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r register.Registration
	r.ServiceName = register.TeacherPortal
	r.ServiceURL = serviceAddress
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.RequiredServices = []register.ServiceName{
		register.GradingService,
		register.LogService,
	}

	ctx, err := service.Start(context.Background(), r, host, port, teacherportal.RegisterHandler)

	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := register.GetProvider(register.LogService); err == nil {
		fmt.Printf("Logging service found at %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()

	fmt.Println("Suhtting down teacherPortal service ...")
}