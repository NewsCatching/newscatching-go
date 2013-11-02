package main

import (
    "time"
    "github.com/c9s/gatsby"
)

type Devices struct {
    Id              int64           `field:"id,primary,serial"`
    DeviceId        string          `field:"device_id"`
    RegId           string          `field:"reg_id"`
    SauthSalt       string          `field:"sauth_salt"`
    CreatedAt       *time.Time      `field:"created_at"`
    CreateIp        string          `field:"create_ip"`
    gatsby.BaseRecord
}

func WriteDevice(deviceID string, regID string) string {
    salt := RandomString(19)

    current_time := time.Now()
    device := Devices{}
    device.Init()
    device.DeviceId = deviceID
    device.RegId = regID
    device.SauthSalt = salt
    device.CreatedAt = &current_time
    res := device.Create()
    if res.Error != nil {
        return ""
    }
    return salt
}

func (self *Devices) Init() {
        self.SetTarget(self)
}

// Implement the GetPrimaryKeyValue interface
func (self *Devices) GetPrimaryKeyValue() int64 {
        return self.Id
}

func (self *Devices) SetPrimaryKeyValue(id int64) {
        self.Id = id
}

func (self *Devices) GetTableName() string {
        return "devices"
}