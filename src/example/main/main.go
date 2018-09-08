package main

import (
	"flag"
	"log"
	"net/http"
    "mime"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {

	flag.Parse()
    mime.AddExtensionType(".apk", "application/vnd.android.package-archive")
    mime.AddExtensionType(".plist", "application/xml")
    mime.AddExtensionType(".ipa", "application/iphone")
    http.Handle("/", http.FileServer(http.Dir("public/")))
	err := http.ListenAndServeTLS(*addr, "cert.pem", "key.pem",nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

