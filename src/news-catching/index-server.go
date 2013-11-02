package main

import (
    "fmt"
    "net/http"
    "github.com/c9s/gatsby"
)

func NewsHotestsAction(w http.ResponseWriter, r *http.Request) {

    header := w.Header()
    // header.Set("Content-Type", "text/plain; charset=utf-8")
    header.Set("Content-Type", "application/javascript")
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

    w.WriteHeader(200)

    // news := gatsby.NewQuery("news")
    // news.Select("id", "title")
    // // news.WhereFromMap( gatsby.ArgMap{
    // //     "is_headline": 0,
    // // })
    // sql := news.String()
    // args := news.Args()

    // fmt.Println(sql)
    // fmt.Println(args)

    news := gatsby.NewRecord(&News{}).(*News)

    res := news.Load(10)   // load the record where primary key = 10

    if res.Error != nil {
        fmt.Println(res.Error)
    }
    if res.IsEmpty {
        fmt.Println("Empty result")
    }

    w.(http.Flusher).Flush()
}