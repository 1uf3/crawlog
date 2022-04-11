package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	var isFirstRequest bool = true

	firstRequest := &http.Response{}
	var fbody []byte

	for {
		resq, err := http.Get("http://192.168.11.8:3000")
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
