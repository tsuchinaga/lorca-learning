package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"

	"github.com/zserge/lorca"
	"gitlab.com/tsuchinaga/lorca-learning/07-notification-and-dialog/assets"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

	ui, err := lorca.New(
		"data:text/html,"+url.PathEscape(`<html><head><title>起動中</title><meta http-equiv="Content-Type" content="text/html; charset=utf-8"/></head><body><p>起動中</p></body></html>`),
		"", 480, 320)
	if err != nil {
		log.Fatalln(err)
	}
	defer ui.Close()

	srv, err := newSrv()
	if err != nil {
		log.Fatalln(err)
	}
	defer srv.close()
	go func() {
		if err := srv.run(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	if err := ui.Load(fmt.Sprintf("http://%s", srv.addr())); err != nil {
		log.Fatalln(err)
	}
	<-ui.Done()
}

func newSrv() (*srv, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	return &srv{ln: ln, srv: &http.Server{Handler: http.FileServer(assets.FS)}}, nil
}

type srv struct {
	ln  net.Listener
	srv *http.Server
}

func (s *srv) addr() net.Addr {
	return s.ln.Addr()
}

func (s *srv) run() error {
	return s.srv.Serve(s.ln)
}

func (s *srv) close() []error {
	var ers []error
	if err := s.srv.Shutdown(context.Background()); err != nil {
		ers = append(ers, err)
	}

	if err := s.ln.Close(); err != nil {
		ers = append(ers, err)
	}

	if err := s.srv.Close(); err != nil {
		ers = append(ers, err)
	}

	return ers
}
