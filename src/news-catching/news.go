package main

import (
    "time"
    "github.com/c9s/gatsby"
)

type News struct {
    Id              int64       `field:"id,primary,serial" json:"id,string"`
    Title           string      `field:"title" json:"title,omitempty"`
    Body            string      `field:"body" json:"body,omitempty"`
    PublishTime     *time.Time  `field:"publish_time" json:"publishTime"`
    Raw             string      `field:"raw" json:"-"`
    Url             string      `field:"url" json:"url,omitempty"`
    Guid            string      `field:"Guid" json:"guid,omitempty"`
    OgImage         string      `field:"og_image" json:"ogImage,omitempty"`
    PicPath         string      `field:"pic_path" json:"picUrl,omitempty"`
    ThumbPath       string      `field:"thumb_path" json:"thumbUrl,omitempty"`
    OgDesc          string      `field:"og_description" json:"-"`
    OgTitle         string      `field:"og_title" json:"-"`
    Referral        string      `field:"referral" json:"referral"`
    RssReferral     string      `field:"rss_referral" json:"-"`
    CreateTime      *time.Time  `field:"create_time" json:"createTime"`
    UpdateTime      *time.Time  `field:"update_time" json:"-"`
    DeleteTime      *time.Time  `field:"delete_time" json:"-"`
    IsSupport       int16       `field:"is_support" json:"isSupport,string"`
    IsHeadline      int16       `field:"is_headline" json:"isHeadline,string"`
    gatsby.BaseRecord
}