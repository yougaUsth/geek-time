package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	server *http.Server
)


func NewServer(port string, handle func(writer http.ResponseWriter, request *http.Request)) (*http.Server, error) {
	server := &http.Server{
		Addr: port,
	}
	http.HandleFunc("/index", handle)
	err := http.ListenAndServe("127.0.0.1" + port, nil)

	return server, err

}

func handleService(writer http.ResponseWriter, request *http.Request){
	// Do something ...
	time.Sleep(5 * time.Second)
	fmt.Printf("Service start ...")
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		ser, err := NewServer(":8080", handleService)
		if err != nil {
			return err
		}
		server = ser
		return nil
	})


	g.Go(func() error {
		// 监听Kill signal
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGQUIT, syscall.SIGKILL)
		select {
		case <- sig:
			return server.Shutdown(ctx)
		case <- ctx.Done():
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("Group meet error %v", err)
	}else {
		fmt.Println("Done!")
	}

}
