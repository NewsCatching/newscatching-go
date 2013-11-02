package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "math/rand"
    "time"
)

func (self ApiResponseJson) NoError() bool {
    return self.Status == 0
}

func (self *ApiResponseJson) Error(status int32, message string) {
    self.Status = status
    self.Message = message
}

func RandomString(length int) string {
    randSource := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
    var base = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=,./<>?;:[]{}|`~"
    o, bl := "", int32(len(base))
    for i := 0; i<length; i++ {
        o = o + string(base[int(randSource.Int31n(bl-1))])
    }
    return o
}

func writeResponseJson(w http.ResponseWriter, output ApiResponseJson, callback string) {
    if outputBytes, err := json.Marshal(output); err == nil {
        w.WriteHeader(200)
        if callback != "" {
            w.Write([]byte(callback))
            w.Write([]byte("("))
            w.Write(outputBytes)
            w.Write([]byte(");"))
        } else {
            w.Write(outputBytes)
        }
    } else {
        fmt.Printf("%#v\n", output)
        fmt.Println("Json faild.")
        w.WriteHeader(500)
        return
    }
}