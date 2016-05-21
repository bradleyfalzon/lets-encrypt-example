package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var webroot = flag.String("webroot", "", "web root directory used by certbot")
	var domain = flag.String("domain", "", "domain name used by certbot")
	flag.Parse()

	http.HandleFunc("/", handler) // example handler
	http.Handle("/.well-known/", http.StripPrefix("/.well-known/", http.FileServer(http.Dir(*webroot+".well-known"))))

	go log.Println(http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/"+*domain+"/fullchain.pem", "/etc/letsencrypt/live/"+*domain+"/privkey.pem", nil))

	log.Fatal(http.ListenAndServe(":80", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi")
}
