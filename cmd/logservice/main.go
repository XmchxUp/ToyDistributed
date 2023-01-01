package main

import (
	"context"
	"fmt"
	stdlog "log"
	"toy-distributed/log"
	"toy-distributed/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "6666"

	ctx, err := service.Start(context.Background(), host, port, "LogService", log.RegisterHandlers)
	if err != nil {
		stdlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log services.")
}
