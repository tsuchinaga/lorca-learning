package main

import (
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
<html>
	<head>
		<title>Hello</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	</head>
	<body><p>Hello, world!</p><p>こんにちわーるど</p></body>
</html>
`), "", 480, 320)

	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := ui.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-ui.Done()
}
