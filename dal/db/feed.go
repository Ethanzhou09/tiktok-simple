package db

import (
	"context"
	"time"
	"gorm.io/plugin/dbresolver"
	"gorm.io/gorm"
)


// MGetVideos
//
//	@Description: 获取最近发布的视频
//	@Date 2023-01-21 16:39:00
//	@param ctx
//	@param limit 获取的视频条数
//	@param latestTime 最早的时间限制
//	@return []*Video 视频列表
//	@return error
func MGetVideos(ctx context.Context, limit int, latestTime *int64) ([]*Video, error) {
	videos := make([]*Video, 0)

	if latestTime == nil || *latestTime == 0 {
		curTime := time.Now().UnixMilli()
		latestTime = &curTime
	}
	conn := GetDB().Clauses(dbresolver.Read).WithContext(ctx)
	if err := conn.Limit(limit).Order("created_at desc").Find(&videos, "created_at < ?", time.UnixMilli(*latestTime)).Error; err != nil {
		return nil, err
	}
	return videos, nil
}


// GetVideoById
//
//	@Description: 根据视频id获取视频
//	@Date 2023-01-24 15:58:52
//	@param ctx 数据库操作上下文
//	@param videoID 视频id
//	@return *Video 视频数据
//	@return error
func GetVideoById(ctx context.Context, videoID int64) (*Video, error) {
	res := new(Video)
	if err := GetDB().Clauses(dbresolver.Read).WithContext(ctx).First(&res, videoID).Error; err == nil {
		return res, nil
	} else if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else {
		return nil, err
	}
}

// GetVideoListByIDs
//
//	@Description: 根据视频id列表获取视频列表
//	@Date 2023-01-24 16:00:12
//	@param ctx 数据库操作上下文
//	@param videoIDs 视频id列表
//	@return []*Video 视频数据列表
//	@return error
func GetVideoListByIDs(ctx context.Context, videoIDs []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if len(videoIDs) == 0 {
		return res, nil
	}

	if err := GetDB().Clauses(dbresolver.Read).WithContext(ctx).Where("video_id in ?", videoIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}