package main

import (
    // "fmt"
    "net/http"
    // "github.com/c9s/gatsby"
)

func CreateReportAction(w http.ResponseWriter, r *http.Request) {

    header := w.Header()
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

    r.ParseForm()
    accessToken := r.FormValue("access_token")
    // releationNews := r.FormValue("url")
    // nickname := r.FormValue("nickname")
    // reportText := r.FormValue("text")
    output := ApiResponseJson{}

    var device *Devices

    if d, err := CheckAccessToken(accessToken); err != nil {
        output.Error(403, err.Error())
        writeResponseJson(w, output, r.FormValue("callback"))
        w.(http.Flusher).Flush()
        return
    } else {
        device = d
    }

    output.Data = device

    writeResponseJson(w, output, r.FormValue("callback"))

    w.(http.Flusher).Flush()
}

func CreateTalkAction(w http.ResponseWriter, r *http.Request) {

}