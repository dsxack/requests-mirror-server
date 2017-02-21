package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if len(port) <= 0 {
		port = "8000"
	}

	fmt.Printf(
		"Start listen %s port for requests\n",
		port,
	)

	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		writer := io.MultiWriter(response, os.Stdout)

		fmt.Fprintf(
			writer,
			"%s %s %s\n",
			request.Method,
			request.RequestURI,
			request.Proto,
		)

		writer.Write([]byte("Host: " + request.Host + "\n"))

		for key := range request.Header {
			fmt.Fprintf(
				writer,
				"%s: %s\n",
				key,
				request.Header.Get(key),
			)
		}

		writer.Write([]byte("\n\n"))
		io.Copy(writer, request.Body)
	})))
}
