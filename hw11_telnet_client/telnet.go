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
	ctx         context.Context
	cancel      context.CancelFunc
	conn        net.Conn
	connScanner *bufio.Scanner
	address     string
	timeout     time.Duration
	inScanner   *bufio.Scanner
	outWriter   io.Writer
}

func (t *Telnet) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), t.timeout)

	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", t.address)
	if err != nil {
		cancel()
		return err
	}

	t.ctx = ctx
	t.cancel = cancel
	t.conn = conn
	t.connScanner = bufio.NewScanner(conn)
	return nil
}

func (t *Telnet) Close() error {
	t.cancel()
	if err := t.conn.Close(); err != nil {
		return err
	}
	return nil
}

func (t *Telnet) Send() error {
	select {
	case <-t.ctx.Done():
		return nil
	default:
		if !t.inScanner.Scan() {
			return nil
		}
		line := t.inScanner.Text()
		_, err := t.conn.Write([]byte(fmt.Sprintf("%s\n", line)))
		if err != nil {
			t.cancel()
		}
		return err
	}
}

func (t *Telnet) Receive() error {
	select {
	case <-t.ctx.Done():
		return nil
	default:
		if !t.connScanner.Scan() {
			return nil
		}
		line := t.connScanner.Text()
		_, err := t.outWriter.Write([]byte(fmt.Sprintf("%s\n", line)))
		if err != nil {
			t.cancel()
		}
		return err
	}
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Не могу создать соединение на этом этапе, так как физическое соединение и ошибка происходит в функции Connect()
	return &Telnet{
		address:   address,
		timeout:   timeout,
		inScanner: bufio.NewScanner(in),
		outWriter: out,
	}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
