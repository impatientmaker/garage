package main

import (
	"fmt"
	"html"
	"log"
	"os/exec"
	"time"
	"net/http"
)

const (
   CMD = "/usr/bin/usbrelay"
   RELAY="3V6F0_1"
   RELAY_ON=RELAY+"=1"
   RELAY_OFF=RELAY+"=0"
)

func test() string {
	return "tested."
}

func garage() string {
	cmdOn := exec.Command(CMD, RELAY_ON)
	err := cmdOn.Run()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep( 500 * time.Millisecond)
	cmdOff := exec.Command(CMD, RELAY_OFF)
	err = cmdOff.Run()
	if err != nil {
		log.Fatal(err)
	}
	return "Garage toggled."
}

func main() {

    http.HandleFunc("/garage2.0", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, garage())
    })

    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, html.EscapeString( test()))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}

