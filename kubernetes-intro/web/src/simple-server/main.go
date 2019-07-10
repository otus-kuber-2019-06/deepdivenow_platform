package main

import (
	"net/http"
	"flag"
	"strconv"
)

func main() {

	portPtr := flag.Int("port", 8000, "listen port number")
	fsdirPtr := flag.String("path", "/app", "server filesystem dir")
	prefixPtr := flag.String("prefix", "/", "URI prefix")
	flag.Parse()

	fs := http.FileServer(http.Dir(*fsdirPtr+"/"))
	http.Handle(*prefixPtr, http.StripPrefix(*prefixPtr, fs))

	http.ListenAndServe(":"+strconv.Itoa(*portPtr), nil)
}
