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

func RedisZeroArgs(action, key string) (interface{}, error) {
	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key)
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key)
		}
	}
	return res, err
}

func RedisRunArgs(action, key string, args []string) (interface{}, error) {
	length := len(args)
	if length%2 == 1 {
		return nil, fmt.Errorf("please give key with value.")
	}

	switch length {
	case 0:
		return RedisZeroArgs(action, key)
	case 2:
		return RedisTwoArgs(action, key, args)
	case 4:
		return RedisFourArgs(action, key, args)
	case 6:
		return RedisSixArgs(action, key, args)
	case 8:
		return RedisEightArgs(action, key, args)
	case 10:
		return RedisTenArgs(action, key, args)
	case 12:
		return RedisTwelveArgs(action, key, args)
	case 14:
		return RedisFourteenArgs(action, key, args)
	case 16:
		return RedisSixteenArgs(action, key, args)
	case 18:
		return RedisEighteenArgs(action, key, args)
	case 20:
		return RedisTwentyArgs(action, key, args)
	case 22:
		return RedisTwentytwoArgs(action, key, args)
	default:
		return nil, fmt.Errorf("too many args. you can write function by your self.")
	}
}

func RedisTwoArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("wrong args.should 2,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1])
		}
	}
	return res, err
}

func RedisFourArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 4 {
		return nil, fmt.Errorf("wrong args.should 4,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3])
		}
	}
	return res, err
}

func RedisSixArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 6 {
		return nil, fmt.Errorf("wrong args.should 6,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5])
		}
	}
	return res, err
}

func RedisEightArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 8 {
		return nil, fmt.Errorf("wrong args.should 8,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
		}
	}
	return res, err
}

func RedisTenArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 10 {
		return nil, fmt.Errorf("wrong args.should 10,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9])
		}
	}
	return res, err
}

func RedisTwelveArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 12 {
		return nil, fmt.Errorf("wrong args.should 12,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11])
		}
	}
	return res, err
}

func RedisFourteenArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 14 {
		return nil, fmt.Errorf("wrong args.should 14,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13])
		}
	}
	return res, err
}
func RedisSixteenArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 16 {
		return nil, fmt.Errorf("wrong args.should 16,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15])
		}
	}
	return res, err
}

func RedisEighteenArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 18 {
		return nil, fmt.Errorf("wrong args.should 18,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17])
		}
	}
	return res, err
}
func RedisTwentyArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 20 {
		return nil, fmt.Errorf("wrong args.should 20,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17], args[18], args[19])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17], args[18], args[19])
		}
	}
	return res, err
}

func RedisTwentytwoArgs(action, key string, args []string) (interface{}, error) {
	if len(args) != 22 {
		return nil, fmt.Errorf("wrong args.should 22,got %d.", len(args))
	}

	rc := RedisClient.Get()
	defer rc.Close()

	res, err := rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17], args[18], args[19], args[20], args[21])
	if err != nil {
		for err != nil && err.Error() == "EOF" {
			rc = RedisClient.Get()
			defer rc.Close()
			res, err = rc.Do(action, key, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17], args[18], args[19], args[20], args[21])
		}
	}
	return res, err
}
