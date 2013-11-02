package main

import (
    // "fmt"
    "net/http"
    "time"
    "crypto/sha256"
    "io"
    "github.com/c9s/gatsby"
)

func AuthAction(w http.ResponseWriter, r *http.Request) {

    header := w.Header()
    // header.Set("Content-Type", "text/plain; charset=utf-8")
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

    w.WriteHeader(200)

    r.ParseForm()
    var salt string
    deviceID := r.FormValue("deviceID")
    regID := r.FormValue("regID")
    output := ApiResponseJson{}
    device := Devices{}

    rows, err := gatsby.QuerySelectWith(DbConnect, &device, "device_id = \"?\"", deviceID)
    if err != nil {
        if salt = WriteDevice(deviceID, regID); salt != "" {
            GCMRegistration(regID)
        } else {
            output.Error(501, "Write device failed")
        }
    } else {
        if data, err := gatsby.CreateStructSliceFromRows(&device, rows); err != nil {
            output.Error(501, err.Error())
        } else {
            device = data.(Devices)
            if device.RegId == regID {
                salt = device.SauthSalt
            } else {
                output.Error(401, "")
            }
        }
    }

    if output.NoError() {
        sha256h := sha256.New()
        token := string(device.Id) + "-" + string(time.Now().Unix()) + "-" + salt
        io.WriteString(sha256h, token)
        output.Data = token + "-" + string(sha256h.Sum(nil))
    }

    writeResponseJson(w, output, r.FormValue("callback"))

    w.(http.Flusher).Flush()
}