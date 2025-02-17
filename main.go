package main

import (
    "log"
    "net/http"
)

func main() {
    const port = "8080"
    // use http.ServeMux to start a server and handle the root path /
    mux := http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir(".")))
    // create a new http.Server struct with mux as the handler and 8080 as the .Addr field
    server := &http.Server{
        Addr:    ":" + port,
        Handler: mux,
    }

    // start the server with ListenAndServe
    log.Printf("Starting server on port %s", port)
    log.Fatal(server.ListenAndServe())
}
