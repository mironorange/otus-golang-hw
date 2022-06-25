package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Telnet struct {
	address     string
	timeout     time.Duration
	conn        net.Conn
	out         io.Writer
	inScanner   *bufio.Scanner
	connScanner *bufio.Scanner
	cancelFunc  context.CancelFunc
}

func (t *Telnet) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err == nil {
		t.conn = conn
		t.connScanner = bufio.NewScanner(conn)
	}
	return err
}

func (t *Telnet) Close() error {
	if t.conn != nil {
		if err := t.conn.Close(); err != nil {
			t.cancelFunc()
			return err
		}
	}
	return nil
}

func (t *Telnet) Send() error {
	if !t.inScanner.Scan() {
		t.cancelFunc()
		return nil
	}

	line := t.inScanner.Bytes()
	if _, err := t.conn.Write([]byte(fmt.Sprintf("%s\n", line))); err != nil {
		return err
	}

	return nil
}

func (t *Telnet) Receive() error {
	if !t.connScanner.Scan() {
		t.cancelFunc()
		return nil
	}

	line := t.connScanner.Bytes()
	if _, err := t.out.Write([]byte(fmt.Sprintf("%s\n", line))); err != nil {
		return err
	}

	return nil
}

func NewTelnetClient(
	address string,
	timeout time.Duration,
	in io.ReadCloser,
	out io.Writer,
	cancelFunc context.CancelFunc,
) TelnetClient {
	return &Telnet{
		address:    address,
		timeout:    timeout,
		out:        out,
		inScanner:  bufio.NewScanner(in),
		cancelFunc: cancelFunc,
	}
}
