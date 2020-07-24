package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"gitlab.com/tsuchinaga/lorca-learning/03-file-server-html/assets"

	"github.com/zserge/lorca"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatalln(err)
	}
	defer closer(ui)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer closer(ln)
	go func() {
		if err := http.Serve(ln, http.FileServer(assets.FS)); err != nil {
			log.Fatalln(err)
		}
	}()
	if err := ui.Load(fmt.Sprintf("http://%s", ln.Addr())); err != nil {
		log.Fatalln(err)
	}

	<-ui.Done()
}

func closer(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatalln(err)
	}
}
