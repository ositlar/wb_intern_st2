package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	t    string
	host string
	port string
)

func init() {
	flag.StringVar(&t, "t", "10s", "timeout for connecting to the server")
	flag.StringVar(&host, "host", "", "host")
	flag.StringVar(&port, "port", "", "port")
}

func main() {
	flag.Parse()
	// if len(flag.Args()) != 3 {
	// 	fmt.Println("use : go-telnet --timeout=10s <host> <port>")
	// 	return
	// }
	timeout, err := time.ParseDuration(t)
	if err != nil {
		log.Fatal(err)
	}
	if host == "" || port == "" {
		log.Fatal("Неверно указаны аргументы")
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	connChan := make(chan string)
	inChan := make(chan string)
	var wg sync.WaitGroup
	ipAddr := fmt.Sprintf("%s:%s", host, port)
	conn, err := net.Dial("tcp", ipAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go ScanInput(inChan, cancel)
	go ScanConn(connChan, conn, cancel)

	wg.Add(1)
	go func() {
		defer wg.Done()
		ManageMsgs(ctx, conn, connChan, inChan)
	}()
	wg.Wait()
}

func ScanInput(inputChan chan<- string, cancel context.CancelFunc) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		msg = fmt.Sprintf("%s\n", msg)
		inputChan <- msg
	}
	cancel()
}

func ScanConn(connChan chan<- string, conn net.Conn, cancel context.CancelFunc) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		connChan <- msg
	}
	cancel()
}

func ManageMsgs(ctx context.Context, conn net.Conn, connChan <-chan string, inChan <-chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-connChan:
			fmt.Println(msg)
		case msg := <-inChan:
			conn.Write([]byte(msg))
		}
	}
}
