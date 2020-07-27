package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/zserge/lorca"
	"gitlab.com/tsuchinaga/lorca-learning/04-quit-button/assets"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatalln(err)
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	srv := &http.Server{Handler: http.FileServer(assets.FS)}
	defer srv.Close()
	go func() {
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	messageHidden := true
	_ = ui.Bind("quit", func() {
		go quit(srv, ui)
		messageHidden = false
		ui.Eval("rerender()")
	})
	_ = ui.Bind("messageHidden", messageHidden)
	if err := ui.Load(fmt.Sprintf("http://%s", ln.Addr())); err != nil {
		log.Fatalln(err)
	}

	<-ui.Done()
}

func quit(srv *http.Server, ui lorca.UI) {
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
	ui.Close()
}
