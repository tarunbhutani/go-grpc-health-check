package pkg

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Server interface {
	Start(wg *sync.WaitGroup, stop_signal chan struct{})
	Stop()
}

type Deamon struct {
	server_list []Server
}

func NewDaemonServer() *Deamon {
	return &Deamon{}
}

func (d *Deamon) Add(server Server) {
	d.server_list = append(d.server_list, server)
}

func (d *Deamon) Start() {
	var wg sync.WaitGroup
	stop_signal := make(chan os.Signal, 1)
	stop := make(chan struct{}, 1)
	signal.Notify(stop_signal, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)

	for _, server := range d.server_list {
		wg.Add(1)
		go server.Start(&wg, stop)
	}

	<-stop_signal
	close(stop)
	fmt.Println("Shutdown...requested")
	waitAndTimeOut(&wg, 30*time.Second)
	fmt.Println("Shutdown...")
}

func waitAndTimeOut(wg *sync.WaitGroup, timeout time.Duration) {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		wg.Wait()
	}()
	select {
	case <-ch:
		return
	case <-time.After(timeout):
		return
	}
}
