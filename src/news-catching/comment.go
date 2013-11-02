package main

import (
    "time"
    "github.com/c9s/gatsby"
)

type Comments struct {
    Id              int64       `field:"id,primary,serial" json:"id,string"`
    DeviceID        int64       `field:"device_id" json:"deviceId,string"`
    NewsID          int64       `field:"news_id" json:"newsId,string"`
    News            *News       `field:"-" json:"news,omitempty"`
    Type            string      `field:"type" json:"type,string"`
    Nickname        string      `field:"nickname" json:"nickname"`
    Ip              string      `field:"ip" json:"-"`
    Text            string      `field:"text" json:"text"`
    CreatedAt        *time.Time  `field:"created_at" json:"createTime"`
    DeletedAt        *time.Time  `field:"deleted_at" json:"-"`
    gatsby.BaseRecord
}

func (self *Comments) Init() {
        self.SetTarget(self)
}

// Implement the GetPrimaryKeyValue interface
func (self *Comments) GetPrimaryKeyValue() int64 {
        return self.Id
}

func (self *Comments) SetPrimaryKeyValue(id int64) {
        self.Id = id
}

func (self *Comments) GetTableName() string {
        return "comments"
}