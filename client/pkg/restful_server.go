package pkg

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type RestfulServer struct {
	Port    int
	Handler http.Handler
}

func (server *RestfulServer) Start(wg *sync.WaitGroup, stop_signal chan struct{}) {
	go func() {
		defer wg.Done()

		http_server := &http.Server{Addr: fmt.Sprintf(":%d", server.Port), Handler: server.Handler}
		serverCtx, serverStopCtx := context.WithCancel(context.Background())

		go func() {
			<-stop_signal

			shutdownCtx, shutdownCtxCancel := context.WithTimeout(serverCtx, 30*time.Second)
			defer shutdownCtxCancel()

			go func() {
				<-shutdownCtx.Done()
				if shutdownCtx.Err() == context.DeadlineExceeded {
					log.Fatal("rest_server graceful shutdown timed out.. forcing exit.")
				}
			}()

			// Trigger graceful shutdown
			err := http_server.Shutdown(shutdownCtx)
			if err != nil {
				log.Fatal(err)
			}
			serverStopCtx()
		}()
		fmt.Println("Starting Server...")
		// Run the server
		err := http_server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}

		<-serverCtx.Done()
	}()
}

func (rs *RestfulServer) Stop() {

}
