package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"net/http"
	"os"
	"strings"
	"path"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func main() {

	url := "https://i.artfile.ru/5000x3220_703804_[www.ArtFile.ru].jpg"
	name := path.Base(url)
	err := downloadFile(name, url)
	if err != nil {
		panic(err)
	}

	fmt.Println("finished")
}

func downloadFile(name string, url string) error {

	img, err := os.Create(name)
	if err != nil {
		return err
	}
	defer img.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	counter := &WriteCounter{}
	_, err = io.Copy(img, io.TeeReader(response.Body, counter))
	if err != nil {
		return err
	}

	fmt.Print("\n")

	return nil
}