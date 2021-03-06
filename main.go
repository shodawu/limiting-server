package main

import (
	"flag"
	"fmt"
	"limiting/clients"
	"limiting/server"
	"log"
	"net/http"
	"time"
)

func NewLimiter(hitMax uint, resetPolling time.Duration) *server.Limiter {
	return &server.Limiter{
		HitCount: server.HitCount{
			HitLimit:      hitMax,
			ResetInterval: resetPolling,
		},
		CmdRev: make(chan server.ExeCmd),
	}
}

var side = flag.String("s", "s", "s to run server side, c to run client side.")
var hitMax = flag.Uint("m", 60, "set hit limit, default is 60 times.")
var resetPolling = flag.Duration("p", 60, "set how long to reset hit limit, default is 60 seconds.")
var cli = flag.Int("c", 10, "run parallel clients, default is 10 clients.")
var dur = flag.Duration("d", 10, "take a break after a client make a HTTP request, default is 10 milliseconds.")
var retryInterval = flag.Duration("i", 3, "retry later if there is a error to handshake with server, default is 3 seconds.")

func main() {
	flag.Parse()

	switch *side {
	case "s":
		fmt.Println("Run Server...")
		limiter := NewLimiter(*hitMax, *resetPolling)
		go limiter.ResetPolling()
		go limiter.CommandPolling()

		http.HandleFunc("/hit", limiter.HitHandler)
		log.Fatal(http.ListenAndServe(":1234", nil))
	case "c":
		fmt.Println("Run Clients...")
		clients.MultiRun(*cli, *dur, *retryInterval)
	}

}
