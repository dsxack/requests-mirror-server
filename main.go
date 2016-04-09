package main

import (
	"net/http"
	"log"
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if (len(port) <= 0) {
		port = "8000"
	}

	fmt.Printf(
		"Start listen %s port for requests\n",
		port,
	)

	log.Fatal(http.ListenAndServe(":" + port, http.HandlerFunc(func (response http.ResponseWriter, request *http.Request) {

		fmt.Printf(
			"%s %s \n",
			request.Method,
			request.RequestURI,
		);

		response.Write([]byte(fmt.Sprintf(
			"%s %s %s\n",
			request.Method,
			request.RequestURI,
			request.Proto,
		)))
		response.Write([]byte("Host: " + request.Host + "\n"))

		for key := range request.Header {
			response.Write([]byte(fmt.Sprintf(
				"%s: %s\n",
				key,
				request.Header.Get(key),
			)))
		}

		buffer := make([]byte, 1024)
		body_started := false

		for {
			n, err := request.Body.Read(buffer)

			if (n > 0) {
				if (!body_started) {
					response.Write([]byte("\n\n"))
					body_started = true
				}

				response.Write(buffer)

				if f, ok := response.(http.Flusher); ok {
					f.Flush()
				}
			}

			if (err != nil) {
				break
			}
		}
	})))
}