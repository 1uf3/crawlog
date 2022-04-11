package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	url := flag.String("url", "http://localhost:3000", "checking url website")

	var isFirstRequest bool = true
	var fbody []byte
	firstRequest := &http.Response{}

	for {
		resq, err := http.Get(*url)
		if err != nil {
			log.Fatal("HTTP protocol error")
		}

		if resq.StatusCode != 200 {
			log.Println("Server is not found")
			resq.Body.Close()
			continue
		}

		if isFirstRequest {
			body, err := io.ReadAll(resq.Body)
			if err != nil {
				log.Fatal("Error: doesn't get body")
			}
			fbody = body
			resq.Body.Close()

			firstRequest = resq

			isFirstRequest = false
			continue
		}

		body, err := io.ReadAll(resq.Body)
		if err != nil {
			log.Fatal("Error: doesn't get body")
		}
		resq.Body.Close()

		if firstRequest.ContentLength != resq.ContentLength {
			log.Println("ContentLength Edited! Check the Webpage!")
		}
		if strings.Compare(string(fbody), string(body)) != 0 {
			log.Println("Website Edited! Check the Webpage!")

			log.Println("--------------------------------------------------")
			log.Println("------------------- diff -------------------------")
			log.Println("--------------------------------------------------")
			log.Println(string(fbody))
			log.Println("--------------------------------------------------")
			log.Println(string(body))
			log.Println("--------------------------------------------------")
			log.Println("--------------------------------------------------")
			log.Println("--------------------------------------------------")
		}

		time.Sleep(time.Minute)
	}
}
