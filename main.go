package crawlog

import (
	"log"
	"net/http"
)

func main() {

	var isFirstRequest bool = true

	firstRequest := &http.Response{}

	for {
		resq, err := http.Get("https://lufe.jp")
		if err != nil {
			log.Fatal("HTTP protocol error")
		}

		if resq.StatusCode != 200 {
			log.Println("Server is not found")
			continue
		}

		if isFirstRequest {
			firstRequest = resq

			isFirstRequest = false
			continue
		}

		if firstRequest.ContentLength == resq.ContentLength {
			log.Println("ContentLength Edited! Check the Webpage!")
		}
		if firstRequest.Body == resq.Body {
			log.Println("Website Edited! Check the Webpage!")
		}
	}
}
