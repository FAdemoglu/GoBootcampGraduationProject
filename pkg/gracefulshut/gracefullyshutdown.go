package gracefulshut

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func ShutdownGin(instance *http.Server, timeout time.Duration) {

	close := make(chan os.Signal)
	signal.Notify(close)
	<-close

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := instance.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds...")
	}
	log.Println("Server exiting")
}
