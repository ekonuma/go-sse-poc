package main

import (
	"fmt"
	"net/http"

	"github.com/ekonuma/go-sse-poc/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	out := make(chan amqp.Delivery)
	rabbitmqChannel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	go rabbitmq.Consume("messages", rabbitmqChannel, out)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "res/index.html")
	})
	http.HandleFunc("/sse", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		for message := range out {
			fmt.Fprintf(w, "event: message\n")
			fmt.Fprintf(w, "Data: %s\n\n", message.Body)
			w.(http.Flusher).Flush()
		}
	})
	http.ListenAndServe(":8080", nil)
}
