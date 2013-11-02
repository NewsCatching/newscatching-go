package main

import (
    "net/http"
)

func PingServer(w http.ResponseWriter, r *http.Request) {

    header := w.Header()
    // header.Set("Content-Type", "text/plain; charset=utf-8")
    header.Set("Content-Type", "application/javascript")
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

    w.WriteHeader(200)

    w.Write([]byte("PONG"))

    w.(http.Flusher).Flush()
}