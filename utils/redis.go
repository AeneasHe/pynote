package utils

import (
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const DefaultTimeFormat = "2006-01-02 15:04:05"

type RedisConfig struct {
	Name     string `required:"true"`
	Password string `default:""`
	Port     uint   `default:"6379"`
	DB       int    `default:"15"`
}

var (
	RedisClient *redis.Client
	redisOnce   sync.Once
)

func GetRedisConn(conf *RedisConfig) *redis.Client {
	redisOnce.Do(func() {
		RedisClient = redis.NewClient(&redis.Options{
			Addr: conf.Name + ":" +
				strconv.FormatUint(uint64(conf.Port), 10),
			Password: conf.Password,
			DB:       conf.DB,
		})
	})
	return RedisClient
}

func RedisConnClose() {
	RedisClient.Close()
}

func RedisLimit(key string, conf *RedisConfig) bool {
	conn := GetRedisConn(conf)

	exist, err := conn.Exists(key).Result()
	if err != nil {
		return true
	}
	if exist == 1 {
		return true
	}

	_, err = conn.Set(key, true, time.Second).Result()
	if err != nil {
		return true
	}
	return false
}

var expireTime = time.Hour * 24

func RedisUserLimit(user string, conf *RedisConfig, limit int64) bool {
	redisConn := GetRedisConn(conf)

	var (
		limitKey    = "limit_user:" + user
		transferKey = "transfer:" + user
		counterKey  = "transfer_counter:" + user
	)

	exist, err := redisConn.Exists(limitKey).Result()
	if err != nil {
		logrus.Errorln(err)
		return true
	}
	if exist == 1 {
		return true
	}

	var count int64
	exist, err = redisConn.Exists(transferKey).Result()
	if err != nil {
		logrus.Errorln(err)
		return true
	}
	if exist != 1 {
		count, err = redisConn.Incr(transferKey).Result()
		if err != nil {
			logrus.Errorln(err)
			return true
		}
		redisConn.HMSet(counterKey, map[string]interface{}{
			"count":      1,
			"created_at": time.Now().Format(DefaultTimeFormat),
		})
		b := ExpireKeys(conf, expireTime, transferKey, counterKey)
		if b {
			return true
		}
	} else {
		count, err = redisConn.Incr(transferKey).Result()
		if err != nil {
			logrus.Errorln(err)
			return true
		}
		_, err = redisConn.HIncrBy(counterKey, "count", 1).Result()
		if err != nil {
			logrus.Errorln(err)
			return true
		}
	}

	// todo 需要改回来的
	if count >= limit {
		_, err = redisConn.Set(limitKey, true, expireTime).Result()
		if err != nil {
			logrus.Errorln(err)
			return true
		}
		b := ExpireKeys(conf, expireTime, counterKey)
		if b {
			return true
		}
		_, err = redisConn.Del(transferKey).Result()
		if err != nil {
			logrus.Errorln(err)
			return true
		}
	}
	return false
}

func ExpireKeys(conf *RedisConfig, expire time.Duration, keys ...string) bool {
	redisConn := GetRedisConn(conf)
	// defer redisConn.Close()

	for _, key := range keys {
		b, err := redisConn.Expire(key, expire).Result()
		if err != nil {
			logrus.Errorln(err)
			return true
		}
		if !b {
			return true
		}
	}

	return false
}

//func RedisQueuePush(queueName string, values string, conf *RedisConfig) error {
//	redisConn := GetRedisConn(conf)
//	defer redisConn.Close()
//
//	_, err := redisConn.LPush(queueName, values).Result()
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func RedisQueuePop(queueName string, timeout time.Duration, conf *RedisConfig) (data string, err error) {
//	redisConn := GetRedisConn(conf)
//	defer redisConn.Close()
//
//	datas, err := redisConn.BRPop(timeout, queueName).Result()
//	if err != nil {
//		if err == redis.Nil {
//			return "", errors.New("pop result is nil")
//		}
//		return "", err
//	}
//	return datas[1], nil
//}
