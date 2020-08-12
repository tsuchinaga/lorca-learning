package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"gitlab.com/tsuchinaga/lorca-learning/08-embed-resources/assets"

	"github.com/zserge/lorca"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

	css, _ := getAssets("/css/style.css")
	js, _ := getAssets("/js/main.js")
	img, _ := getAssetsBase64("/img/f_f_object_170_s512_f_object_170_2nbg.png")

	ui, err := lorca.New(
		"data:text/html,"+url.PathEscape(fmt.Sprintf(`<html>
	<head>
		<title>画像の埋め込み</title>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
		<style>%s</style>
		<script>%s</script>
	</head>
	<body>
		<img src="data:;base64,%s">
	</body>
</html>`, css, js, img)), "", 600, 600)
	if err != nil {
		log.Fatalln(err)
	}
	defer ui.Close()
	<-ui.Done()
}

func getAssets(name string) (string, error) {
	f, err := assets.FS.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func getAssetsBase64(name string) (string, error) {
	f, err := assets.FS.Open(name)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
