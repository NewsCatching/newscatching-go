package main

import (
    "github.com/c9s/gatsby"
)

type News struct {
    Id              int64     `field:"id,primary,serial"`
    Title           string    `field:"title"`
    Body            string    `field:"body"`
    PublishTime     uint64    `field:"publish_time"`
    Raw             string    `field:"raw"`
    Url             string    `field:"url"`
    Guid            string    `field:"Guid"`
    OgImage         string    `field:"og_image"`
    PicPath         string    `field:"pic_path"`
    ThumbPath       string    `field:"thumb_path"`
    OgDesc          string    `field:"og_description"`
    OgTitle         string    `field:"og_title"`
    Referral        string    `field:"referral"`
    RssReferral     string    `field:"rss_referral"`
    CreateTime      uint64    `field:"create_time"`
    UpdateTime      uint64    `field:"update_time"`
    DeleteTime      uint64   `field:"delete_time"`
    IsSupport       uint16    `field:"is_support"`
    IsHeadline      uint16    `field:"is_headline"`
    gatsby.BaseRecord
}