package redis

import (
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const ExpireTime = 6 * time.Second

var (
	redisOnce     sync.Once
	redisHelper   *RedisHelper
	FavoriteMutex *redsync.Mutex
	RelationMutex *redsync.Mutex
)

type RedisHelper struct {
	*redis.Client
}

func GetRedisHelper() *RedisHelper {
	return redisHelper
}

// LockByMutex Obtain a lock for our given mutex. After this is successful, no one else can obtain the same lock (the same mutex name) until we unlock it.
func LockByMutex(ctx context.Context, mutex *redsync.Mutex) error {
	if err := mutex.LockContext(ctx); err != nil {
		return err
	}
	return nil
}

// UnlockByMutex Release the lock so other processes or threads can obtain a lock.
func UnlockByMutex(ctx context.Context, mutex *redsync.Mutex) error {
	if _, err := mutex.UnlockContext(ctx); err != nil {
		return err
	}
	return nil
}

func NewRedisHelper() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", "127.0.0.1","6379"),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		//MaxConnAge:   1 * time.Minute,	go-redis v9 已删去
		PoolSize:    10,
		PoolTimeout: 30 * time.Second,
	})

	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client = rdb
		redisHelper = rdh
	})
	return rdb
}

func init() {
	ctx := context.Background()
	rdb := NewRedisHelper()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return
	}

	// 开启定时同步至数据库
	GoCronFavorite()
	GoCronRelation()

	// Redis锁
	// 创建Redis连接池
	pool := goredis.NewPool(rdb)
	// Create an instance of redisync to be used to obtain a mutual exclusion lock.
	rs := redsync.New(pool)
	// Obtain a new mutex by using the same name for all instances wanting the same lock.
	FavoriteMutex = rs.NewMutex("mutex-favorite")
	RelationMutex = rs.NewMutex("mutex-relation")
}
