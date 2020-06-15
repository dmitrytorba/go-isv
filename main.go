package main

import (
    "fmt"
    "net/http"
    "log"
)

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Login here\n")
}

func createHandler() http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("/login", login)

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