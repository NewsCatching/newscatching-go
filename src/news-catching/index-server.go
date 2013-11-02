package main

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"
    "github.com/c9s/gatsby"
    "strconv"
    // "database/sql"
)

func NewsReadAction(w http.ResponseWriter, r *http.Request) {

    var newsId int64
    var err error
    output := ApiResponseJson{}

    if newsId, err = strconv.ParseInt(r.URL.Path[18:], 10, 64); err != nil {
        output.Error(501, err.Error())
        writeResponseJson(w, output, r.FormValue("callback"))
        w.(http.Flusher).Flush()
        return
    }

    header := w.Header()
    // header.Set("Content-Type", "text/plain; charset=utf-8")
    header.Set("Content-Type", "application/javascript; charset=utf-8")
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

    news := gatsby.NewRecord(&News{}).(*News)

    res := news.Load(newsId)
    if res.Error != nil {
        output.Error(501, res.Error.Error())
    } else {
        if res.IsEmpty {
            output.Error(404, "empty result")
        }
    }

    if output.NoError() {
        data := make(map[string]interface{})
        data["news"] = News{
            Id: news.Id,
            Title: news.Title,
            Body: news.Body,
            PublishTime: news.PublishTime,
            Url: news.Url,
            Guid: news.Guid,
            OgImage: news.OgImage,
            PicPath: news.PicPath,
            ThumbPath: news.ThumbPath,
            Referral: news.Referral,
            CreateTime: news.CreateTime,
            IsSupport: news.IsSupport,
            IsHeadline: news.IsHeadline,
        }

        output.Data = data
        // fmt.Printf("%#v\n", news)
    }
    writeResponseJson(w, output, r.FormValue("callback"))
    w.(http.Flusher).Flush()
}

func NewsHotestsAction(w http.ResponseWriter, r *http.Request) {

    header := w.Header()
    // header.Set("Content-Type", "text/plain; charset=utf-8")
    header.Set("Content-Type", "application/javascript; charset=utf-8")
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

    // news := gatsby.NewQuery("news")
    // news.Select("id", "title")
    // // news.WhereFromMap( gatsby.ArgMap{
    // //     "is_headline": 0,
    // // })
    // sql := news.String()
    // args := news.Args()

    // fmt.Println(sql)
    // fmt.Println(args)

    // news := gatsby.NewRecord(&News{}).(*News)

    // res := news.Load(10)   // load the record where primary key = 10

    // if res.Error != nil {
    //     fmt.Println(res.Error)
    // }
    // if res.IsEmpty {
    //     fmt.Println("Empty result")
    // }

    // fmt.Printf("%#v\n", news)
    output := ApiResponseJson{}
    randSource := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))
    var news_max int64
    err := DbConnect.QueryRow("SELECT MAX(id) FROM news").Scan(&news_max)
    switch {
    case err != nil:
        fmt.Println(err)
        output.Error(501, err.Error())
    default:
        params := make([]interface{}, 7)
        for i := 0; i< 6; i++ {
            params[i] = randSource.Int63n(news_max)
        }
        params[6] = time.Now().AddDate(0,0,-2).Unix()
        fmt.Println(params)
        rows, err := gatsby.QuerySelectWith(DbConnect, &News{}, "WHERE id IN (?,?,?,?,?,?) AND create_time > ? AND thumb_path <> '' ", params...)
        if err == nil {
            news := News{}
            data, err := gatsby.CreateStructSliceFromRows(&news, rows)
            if err != nil {
                fmt.Println(err)
                output.Error(501, err.Error())
            } else {
                output.Data = data.([]News)
            }
        } else {
            output.Error(404, err.Error())
        }
    }
    writeResponseJson(w, output, r.FormValue("callback"))

    w.(http.Flusher).Flush()
}