package db

import (
	"context"
	"time"
	"gorm.io/plugin/dbresolver"
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

