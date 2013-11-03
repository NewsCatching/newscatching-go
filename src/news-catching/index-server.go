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
    freg := r.URL.Path[18:]

    if freg == "" {
        output.Error(404, "No news id.")
        writeResponseJson(w, output, r.FormValue("callback"))
        w.(http.Flusher).Flush()
        return
    }

    if newsId, err = strconv.ParseInt(freg, 10, 64); err != nil {
        output.Error(502, err.Error())
        writeResponseJson(w, output, r.FormValue("callback"))
        w.(http.Flusher).Flush()
        return
    }

    header := w.Header()
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
        data, err := GetNewsMeta(news.Id)
        if err == nil {
            (*data)["news"] = News{
                Id: news.Id,
                Title: news.Title,
                Body: news.Body,
                PublishTime: news.PublishTime,
                Url: news.Url,
                Guid: news.Guid,
                OgImage: news.OgImage,
                PicPath: UrlDomain + news.PicPath[2:],
                ThumbPath: UrlDomain + news.ThumbPath[2:],
                Referral: news.Referral,
                CreateTime: news.CreateTime,
                IsSupport: news.IsSupport,
                IsHeadline: news.IsHeadline,
            }
            output.Data = *data
        } else {
            output.Error(501, err.Error())
        }
    }
    writeResponseJson(w, output, r.FormValue("callback"))
    w.(http.Flusher).Flush()
}

func NewsReportAction(w http.ResponseWriter, r *http.Request) {
    header := w.Header()
    header.Set("Content-Type", "application/javascript; charset=utf-8")
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

    // fmt.Printf("%#v\n", news)
    output := ApiResponseJson{}
    var params []interface{}
    pi := 0
    offset := r.FormValue("offset")
    length := r.FormValue("rows")
    qsearch := r.FormValue("q")
    if offset == "" {
        offset = "0"
    }
    if length == "" {
        length = "20"
    }
    sql := "RIGHT JOIN `comments` ON `comments`.`news_id` = `news`.`id` WHERE 1 AND delete_time IS NULL ORDER BY `comments`.`created_at` DESC LIMIT ?, ? "
    if qsearch != "" {
        sql = "RIGHT JOIN `comments` ON `comments`.`news_id` = `news`.`id` WHERE 1 AND delete_time IS NULL AND title LIKE ? ORDER BY `comments`.`created_at` DESC LIMIT ?, ? "
        params = make([]interface{},3)
        params[pi] = "%" + qsearch + "%"
        pi++
    } else {
        params = make([]interface{},2)
    }
    params[pi] = offset
    pi++
    params[pi] = length
    pi++
    rows, err := gatsby.QuerySelectWith(DbConnect, &News{}, sql, params...)
    if err == nil {
        news := News{}
        data, err := gatsby.CreateStructSliceFromRows(&news, rows)
        if err != nil {
            fmt.Println(err)
            output.Error(501, err.Error())
        } else {
            newsList := data.([]News)
            for k, _ := range newsList {
                newsList[k].Raw = ""
                newsList[k].Body = ""
            }
            output.Data = newsList
        }
    } else {
        output.Error(404, err.Error())
    }
    writeResponseJson(w, output, r.FormValue("callback"))

    w.(http.Flusher).Flush()
}

func NewsListAction(w http.ResponseWriter, r *http.Request) {
    header := w.Header()
    header.Set("Content-Type", "application/javascript; charset=utf-8")
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

    // fmt.Printf("%#v\n", news)
    output := ApiResponseJson{}
    var params []interface{}
    pi := 0
    offset := r.FormValue("offset")
    length := r.FormValue("rows")
    qsearch := r.FormValue("q")
    if offset == "" {
        offset = "0"
    }
    if length == "" {
        length = "20"
    }
    sql := "WHERE 1 AND delete_time IS NULL ORDER BY `create_time` LIMIT ?, ? "
    if qsearch != "" {
        sql = "WHERE 1 AND delete_time IS NULL AND title LIKE ? ORDER BY `create_time` DESC LIMIT ?, ? "
        params = make([]interface{},3)
        params[pi] = "%" + qsearch + "%"
        pi++
    } else {
        params = make([]interface{},2)
    }
    params[pi] = offset
    pi++
    params[pi] = length
    pi++
    rows, err := gatsby.QuerySelectWith(DbConnect, &News{}, sql, params...)
    if err == nil {
        news := News{}
        data, err := gatsby.CreateStructSliceFromRows(&news, rows)
        if err != nil {
            fmt.Println(err)
            output.Error(501, err.Error())
        } else {
            newsList := data.([]News)
            for k, _ := range newsList {
                newsList[k].Raw = ""
                newsList[k].Body = ""
            }
            output.Data = newsList
        }
    } else {
        output.Error(404, err.Error())
    }
    writeResponseJson(w, output, r.FormValue("callback"))

    w.(http.Flusher).Flush()
}

func NewsHotAction(w http.ResponseWriter, r *http.Request) {
    header := w.Header()
    header.Set("Content-Type", "application/javascript; charset=utf-8")
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

    // fmt.Printf("%#v\n", news)
    output := ApiResponseJson{}
    // randomSource := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
    var params []interface{}
    pi := 0
    offset := r.FormValue("offset")
    length := r.FormValue("rows")
    qsearch := r.FormValue("q")
    if offset == "" {
        offset = "0"
    }
    if length == "" {
        length = "20"
    }
    sql := "WHERE 1 AND delete_time IS NULL ORDER BY guid LIMIT ?, ? "
    if qsearch != "" {
        sql = "WHERE 1 AND delete_time IS NULL AND title LIKE ? ORDER BY guid LIMIT ?, ? "
        params = make([]interface{},3)
        params[pi] = "%" + qsearch + "%"
        pi++
    } else {
        params = make([]interface{},2)
    }
    // params[pi] = randomSource.Int31n(29)
    // pi++
    params[pi] = offset
    pi++
    params[pi] = length
    pi++
    rows, err := gatsby.QuerySelectWith(DbConnect, &News{}, sql, params...)
    if err == nil {
        news := News{}
        data, err := gatsby.CreateStructSliceFromRows(&news, rows)
        if err != nil {
            fmt.Println(err)
            output.Error(501, err.Error())
        } else {
            newsList := data.([]News)
            for k, _ := range newsList {
                newsList[k].Raw = ""
                newsList[k].Body = ""
            }
            output.Data = newsList
        }
    } else {
        output.Error(404, err.Error())
    }
    writeResponseJson(w, output, r.FormValue("callback"))

    w.(http.Flusher).Flush()
}

func NewsHotestsAction(w http.ResponseWriter, r *http.Request) {

    header := w.Header()
    header.Set("Content-Type", "application/javascript; charset=utf-8")
    header.Set("X-Content-Type-Options", "nosniff")
    header.Set("Cache-Control", "no-cache, no-store, must-revalidate")
    header.Set("Pragma", "no-cache")
    header.Set("Expires", "Thu, 01 Dec 1994 16:00:00 GMT")

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
        rows, err := gatsby.QuerySelectWith(DbConnect, &News{}, "WHERE id IN (?,?,?,?,?,?) AND create_time > ? AND thumb_path <> '' AND delete_time IS NULL ", params...)
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