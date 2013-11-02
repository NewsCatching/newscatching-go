package main

import (
    // "bottle"
    "fmt"
    "runtime"
    "time"
    "flag"
    "net/http"
    "github.com/garyburd/redigo/redis"
    _ "github.com/go-sql-driver/mysql"
    "github.com/NewsCatching/gatsby"
    "database/sql"
)

var RedisPool *redis.Pool

func main() {

    var (
        configPath string
    )
    flag.StringVar(&configPath, "config", "", "Config path for news-catching.")
    flag.Parse()

    if configPath == "" {
        configPath = "config.json"
    }

    config, err := GetConfig(configPath)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%#v\n", config)

    db, err := sql.Open("mysql", (*config).Mysql)
    if err != nil {
        fmt.Println(err)
        return
    }
    gatsby.SetupConnection(db, gatsby.DriverMysql)

    RedisPool = &redis.Pool{
            MaxIdle: 3,
            MaxActive: 0,
            IdleTimeout: 240 * time.Second,
            Dial: func() (redis.Conn, error) {
                c, err := redis.Dial("tcp", (*config).Redis)
                if err != nil {
                    return nil, err
                }
                return c, err
            },
            TestOnBorrow: func(c redis.Conn, t time.Time) error {
                _, err := c.Do("PING")
                return err
            },
        }

    cpu_num := runtime.NumCPU()
    fmt.Println("Cpu num: ", cpu_num)
    runtime.GOMAXPROCS(cpu_num)

    http.HandleFunc("/api/v1/ping", PingAction)
    http.HandleFunc("/api/v1/index", IndexAction)

    http.ListenAndServe((*config).Server, nil)

    return
}