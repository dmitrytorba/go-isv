package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/dmitrytorba/go-isv/handlers"
)

func createHandler() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/login", handlers.Login)

    return mux
}

func main() {
    server := &http.Server{
        Addr: fmt.Sprintf(":8000"),
        Handler: createHandler(),
    }

    log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
    if err := server.ListenAndServe(); err != http.ErrServerClosed {
        log.Printf("%v", err)
    } else {
        log.Println("Server closed!")
    }
}