package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 480, 320)
	if err != nil {
		log.Fatalln(err)
	}
	if err := load(ui, "./www/index.html"); err != nil {
		log.Fatalln(err)
	}
	defer closer(ui)
	<-ui.Done()
}

func load(ui lorca.UI, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer closer(file)

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return ui.Load("data:text/html," + url.PathEscape(string(b)))
}

func closer(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatalln(err)
	}
}
