package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/liyue201/goqr"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

type result struct {
	filename string
	path     string
	url      string
	status   int
}

func main() {

	var input string
	var status bool
	channel := make(chan result)
	statusCheckChannel := make(chan int)

	flag.StringVar(&input, "input", "", "image file path. Can also glob pattern like /home/user/*.jpeg")
	flag.BoolVar(&status, "status", false, "Perform an HTTP GET request to the decoded url")
	flag.Parse()

	if input == "" {
		log.Fatal("input path is required")
	}

	matches, err := filepath.Glob(input)
	if err != nil {
		log.Fatal("Path not found")
	}

	for _, path := range matches {
		go decodeQRCode(path, channel)
	}

	for i := 0; i < len(matches); i++ {
		chMessage := <-channel
		message := fmt.Sprintf("\n%s\n%s", chMessage.filename, chMessage.url)

		if status {
			go checkURL(chMessage.url, chMessage.filename, statusCheckChannel)
			chMessage.status = <-statusCheckChannel
			message = fmt.Sprintf("\n%s\n%s\nstatus code:%d", chMessage.filename, chMessage.url, chMessage.status)
		}

		fmt.Println(message)

	}
}

func decodeQRCode(path string, ch chan result) {
	imgData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("error in path: ", err)
	}
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Fatal("error in decoding from image: ", err, path)
	}
	qrCode, err := goqr.Recognize(img)
	if err != nil {
		log.Fatal("qrcode decoding error: ", err)
	}
	ch <- result{
		filename: getFileName(path),
		path:     path,
		url:      string(qrCode[0].Payload),
	}
}

func checkURL(qrURL string, filename string, ch chan int) {
	resp, err := http.Get(qrURL)
	if err != nil {
		fmt.Printf("\nFailed to fetch %s\nFile name %s\n\n", qrURL, filename)
		log.Fatal("error : ", err)
	}
	ch <- resp.StatusCode
}

func getFileName(path string) string {
	var newpath string
	switch os := runtime.GOOS; os {
	case "windows":
		newpath = filepath.FromSlash(path)
	default:
		newpath = path
	}
	_, file := filepath.Split(newpath)
	return file
}
