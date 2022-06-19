package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("It`s work!")

	//// Читать из входящего потока ввода и выводить на экран
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	line := scanner.Text()
	//	fmt.Println(line)
	//}

	//// Соединиться по tcp с сервером
	//dialer := &net.Dialer{}
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//defer cancel()
	//
	//conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:9000")
	//if err != nil {
	//	log.Fatalf("Cannot connect: %v", err)
	//}
	//
	//<-ctx.Done()
	//time.Sleep(5*time.Second)
	//conn.Close()

	//// Соединиться по tcp с сервером и записать в него данные из потока ввода
	//dialer := &net.Dialer{}
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//
	//conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:9000")
	//if err != nil {
	//	log.Fatalf("Cannot connect: %v", err)
	//}
	//
	//go readFromConnection(ctx, conn, cancel)
	//go writeToConnection(ctx, conn, cancel)
	//
	//<-ctx.Done()
	//time.Sleep(5*time.Second)
	//conn.Close()

	//// Ctrl-D sends EOF which doesn't make sense when you're not getting input.
	//// Выйти из приложения при условии, что получен из вне сигнал о выходе
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//<-c
	//fmt.Println("Bye!")

	client := NewTelnetClient("127.0.0.1:9000", 5*time.Second, os.Stdin, os.Stdout)
	defer client.Close()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err := client.Connect(); err == nil {
		go func() {
			for {
				if err := client.Send(); err != nil {
					return
				}
			}
		}()

		go func() {
			for {
				if err := client.Receive(); err != nil {
					return
				}
			}
		}()
	}

	<-ctx.Done()
}

func readFromConnection(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	scanner := bufio.NewScanner(conn)
	func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !scanner.Scan() {
					return
				}
				text := scanner.Text()
				log.Printf("From server: %s", text)
			}
		}
	}()
	cancel()
}

func writeToConnection(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	scanner := bufio.NewScanner(os.Stdin)
	func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !scanner.Scan() {
					return
				}
				line := scanner.Text()
				log.Printf("To server %v\n", line)

				conn.Write([]byte(fmt.Sprintf("%s\n", line)))
			}
		}
	}()
	cancel()
}
