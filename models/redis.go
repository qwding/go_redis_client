package models

import (
	"github.com/garyburd/redigo/redis"
	// "strconv"
	"fmt"
	"time"
)

var (
	RedisHost     = ""
	RedisPassword = ""
	Redisdb       = 0
	RedisMaxIdle  = 50
	RedisMaxConn  = 50
)

var RedisClient *redis.Pool

func init() {
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     RedisMaxIdle,
		MaxActive:   RedisMaxConn,
		IdleTimeout: 180 * 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost)
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", Redisdb)
			c.Send("AUTH", RedisPassword)
			return c, nil
		},
	}
}

func JudgeSetSuccess(res interface{}) bool {
	method := "JudgeSetSuccess"
	if resInt, ok := res.(int64); ok {
		if resInt == int64(1) {
			return true
		}
	}
	if resStr, ok := res.(string); ok {
		if resStr == "OK" {
			return true
		}
	}
	if res == int64(1) {
		return true
	}

	fmt.Println(method, "set false, and res is ", res)
	return false
}

func RedisOnlyCmd(action string) (interface{}, error) {
	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action)
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action)
		}
	}
	return res, err
}

func RedisRun(action, key string, args ...interface{}) (interface{}, error) {
	/*	if len(args) != 2 {
		return nil, fmt.Errorf("wrong args.should 2,got %d.", len(args))
	}*/

	rc := RedisClient.Get()
	defer rc.Close()

	keyargs := make([]interface{}, len(args)+1)
	keyargs[0] = key
	copy(keyargs[1:], args)

	res, err := rc.Do(action, keyargs...)
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, keyargs...)
		}
	}
	return res, err
}
