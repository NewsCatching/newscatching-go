package main

import (
    "io/ioutil"
    "encoding/json"
)

func GetConfig(path string) (*JfConfig, error) {
    var config JfConfig
    jsonStr, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    if err := json.Unmarshal(jsonStr, &config); err != nil {
        return nil, err
    }
    return &config, nil
}

type JfConfig struct {
    Redis string `json:"redis"`
    Path map[string]string `json:"path"`
    Fonts []FontConfig `json:"fonts"`
}

type FontConfig struct {
    Name string `json:"name"`
    Path string `json:"path"`
    Free bool `json:"free"`
    Vip bool `json:"vip"`
}