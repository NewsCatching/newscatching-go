package main

import (
    // "bottle"
    "fmt"
    "runtime"
    "time"
    "flag"
    "net/http"
    "github.com/garyburd/redigo/redis"
)

var RedisPool *redis.Pool

func main() {

    var configPath string
    flag.StringVar(&configPath, "Config Page", "", "Config path for news-catching.")
    flag.Parse()

    if configPath == "" {
        configPath = "config.json"
    }

    config, err := GetConfig(configPath)
    if err != nil {
        fmt.Println(err)
        return
    }

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

    http.HandleFunc("/ping", PingServer)
    http.ListenAndServe(":1234", nil)

    return
}