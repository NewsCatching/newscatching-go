package main

import (
    "io/ioutil"
    "encoding/json"
)

func GetConfig(path string) (*Config, error) {
    var config Config
    jsonStr, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    if err := json.Unmarshal(jsonStr, &config); err != nil {
        return nil, err
    }
    return &config, nil
}

type Config struct {
    Server string `json:"server"`
    Redis string `json:"redis"`
    Mysql string `json:"mysql"`
    UrlDomain string `json:"urlDomain"`
    Path map[string]string `json:"path"`
}