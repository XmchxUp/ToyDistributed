package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, host, port, serviceName string,
	registerHandlersFunc func()) (context.Context, error) {

	registerHandlersFunc()
	ctx = startService(ctx, serviceName, host, port)

	return ctx, nil
}

func startService(ctx context.Context, serviceName, host, port string) context.Context {

	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Printf("%v starte. Press any key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s)
		cancel()
	}()

	return ctx
}
