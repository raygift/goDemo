package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func echo(wr http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	fmt.Printf("HTTP request: %s\n", msg)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}
	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
}
func main() {
	http.HandleFunc("/", echo)
	fmt.Printf("HTTP listen and serve : 192.168.92.3:8080\n")
	err := http.ListenAndServe("192.168.92.3:8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
