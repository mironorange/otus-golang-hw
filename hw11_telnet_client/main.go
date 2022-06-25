package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var connectTimeout time.Duration

func init() {
	flag.DurationVar(&connectTimeout, "timeout", 10*time.Second, "connection timeout")
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatalln("host or port not set")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx, cancelFunc := context.WithCancel(context.Background())
	client := NewTelnetClient(net.JoinHostPort(host, port), connectTimeout, os.Stdin, os.Stdout, cancelFunc)

	log.Println("trying to connect to the server")
	if err := client.Connect(); err == nil {
		log.Println("exchange data")
		go func() {
			for {
				if err := client.Send(); err != nil {
					log.Printf("error sending data, close the connection: %s\n", err)
					cancelFunc()
				}
			}
		}()

		go func() {
			for {
				if err := client.Receive(); err != nil {
					log.Printf("error receiving data, close the connection: %s\n", err)
					cancelFunc()
				}
			}
		}()
	} else {
		cancelFunc()
		client.Close()
		close(c)
		log.Fatalln(err)
	}

	select {
	case <-ctx.Done():
		log.Println("goodbye because done")
		cancelFunc()
		client.Close()
		close(c)
	case <-c:
		log.Println("goodbye because signal")
		cancelFunc()
		client.Close()
	}
}
