package main

import (
    "reflect"
	"io/ioutil"
	"flag"
	"html/template"
	"log"
	"net/http"
	"net/http/httputil"
    "encoding/json"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {

	var i uint
	arr := make([]byte, 32)
	for i = 0; i < ((uint)(len(arr))); i++ {
		arr[i] = (byte)(i)
		log.Println(arr[i])
	}

	flag.Parse()
    http.Handle("/", http.FileServer(http.Dir("public/")))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

