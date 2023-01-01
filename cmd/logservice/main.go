package main

import (
	"context"
	"fmt"
	stdlog "log"
	"toy-distributed/log"
	"toy-distributed/registry"
	"toy-distributed/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "6666"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  serviceAddress,
	}

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		stdlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log services.")
}
