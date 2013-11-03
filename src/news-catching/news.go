package main

import (
    "fmt"
    "time"
    "strings"
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

func GetNewsMeta(newsId int64) (*map[string]interface{}, error) {
    comment := Comments{}
    rows, err := gatsby.QuerySelectWith(DbConnect, &comment, "WHERE news_id = ? AND deleted_at IS NULL ORDER BY type", newsId)
    if err != nil {
        return nil, err
    } else {
        if data, err := gatsby.CreateStructSliceFromRows(&comment, rows); err != nil {
            return nil, err
        } else {
            output := make(map[string]interface{})
            commentsList := data.([]Comments)
            ti, ri, ni, length := 0, 0, 0, len(commentsList)
            reports := make([]Comments, length)
            talks := make([]Comments, length)
            newsIds := make([]string, length)
            newsIdFlags := make(map[int64]bool)
            for _, comment := range commentsList {
                if comment.DeviceID == 0 {
                    if _, ok := newsIdFlags[comment.NewsID]; ! ok {
                        newsIds[ni] = fmt.Sprintf("%d", comment.NewsID)
                        ni ++
                    }
                    reports[ri] = comment
                    ri ++
                } else {
                    talks[ti] = comment
                    ti ++
                }
            }
            if ni > 0 {
                cnews := make(map[int64]*News, ni)
                nrows, err := gatsby.QuerySelectWith(DbConnect, &News{}, "WHERE id IN (" + strings.Join(newsIds, ",") + ") AND delete_time IS NULL ")
                if err == nil {
                    news := News{}
                    data, err := gatsby.CreateStructSliceFromRows(&news, nrows)
                    if err != nil {
                        return nil, err
                    } else {
                        newsList := data.([]News)
                        for _, news := range newsList {
                            news.Body = ""
                            news.PicPath = UrlDomain + news.PicPath[2:]
                            news.ThumbPath = UrlDomain + news.ThumbPath[2:]
                            cnews[news.Id] = &news
                        }
                        for k, v := range reports {
                            reports[k].News = cnews[v.NewsID]
                        }
                    }
                } else {
                    return nil, err
                }
            }
            output["reports"] = reports
            output["talks"] = talks
            return &output, nil
        }
    }
}