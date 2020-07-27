//go:generate go run gen.go

package main

import (
	"log"
	"os"

	"github.com/zserge/lorca"
)

func main() {
	assetsPath := "../assets"
	_, err := os.Stat(assetsPath)
	if os.IsNotExist(err) {
		if err := os.Mkdir(assetsPath, 0666); err != nil {
			log.Fatalln(err)
		}
	}
	if err := lorca.Embed("assets", "../assets/assets.go", "../www"); err != nil {
		log.Fatalln(err)
	}
}
