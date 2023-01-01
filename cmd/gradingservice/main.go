package main

import (
	"context"
	"fmt"
	"toy-distributed/grades"
	"toy-distributed/log"
	"toy-distributed/registry"
	"toy-distributed/service"

	stdlog "log"
)

func main() {
	host, port := "localhost", "6668"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddress + "/services",
	}

	ctx, err := service.Start(context.Background(), host, port,
		r, grades.RegisterHandlers)
	if err != nil {
		stdlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
