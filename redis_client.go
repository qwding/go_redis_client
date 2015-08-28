package main

import (
	"flag"
	"fmt"
	// "github.com/garyburd/redigo/redis"
	"redis_client/models"
	"strings"
	// "encoding/json"
)

var (
	host     *string = flag.String("host", "127.0.0.1:6379", "use your own redis host.default is localhost:6379")
	db       *int    = flag.Int("db", 0, "select the db you want to act.default is 0.")
	password *string = flag.String("password", "password", "give the redis password.default is password")
	maxIdle  *int    = flag.Int("maxIdle", 50, "set max idle. default 50.")
	maxConn  *int    = flag.Int("maxConn", 50, "set max conn. default 50.")
)

func init() {
	flag.Parse()
	models.RedisHost = *host
	models.RedisPassword = *password
	models.Redisdb = *db
	models.RedisMaxIdle = 50
	models.RedisMaxConn = 50
}

func ReadLine() string {
	var in byte
	var err error
	var bytes []byte

	for err == nil {
		_, err = fmt.Scanf("%c", &in)
		if in != '\n' {
			bytes = append(bytes, in)
		} else {
			break
		}
	}
	return string(bytes)
}

func showSelfCmd() {
	fmt.Printf("1. delkeys. delete multiply keys, use like 'delkeys 2015_sample_*'.\n")
}

func main() {
	for {
		fmt.Println("please give the command. type help to see ourself command.")

		cmd := ReadLine()

		if strings.TrimSpace(cmd) == "help" {
			showSelfCmd()
			continue
		}

		arr := strings.Split(cmd, " ")
		for i, _ := range arr {
			arr[i] = strings.TrimSpace(arr[i])
		}

		var res interface{}
		var err error

		length := len(arr)
		if length <= 0 {
			continue
		} else if length == 1 {
			res, err = models.RedisOnlyCmd(arr[0])
		} else if length == 2 {
			res, err = models.RedisZeroArgs(arr[0], arr[1])
		} else {
			switch arr[0] {
			case "delkeys":
				fmt.Println("delete keys")
			default:
				res, err = models.RedisRunArgs(arr[0], arr[1], arr[2:])
			}
		}
		if err != nil {
			fmt.Println("request error", err)
		}

		if resArr, ok := res.([]interface{}); ok {
			fmt.Println("len of res is ", len(resArr))
			for j, val := range resArr {
				if str, ok := val.([]byte); ok {
					fmt.Println("res to string is :", j, string(str))
				}
				if strArr, ok := val.([]interface{}); ok {
					for i, v := range strArr {
						fmt.Println("res to array ", j, i, string(v.([]byte)))
					}

				}
			}
		} else {
			fmt.Println("res is ", res)
		}
	}
}

func delkeys(key string) {
	method := "delkeys"
	res, err := models.RedisZeroArgs("keys", key)
	if err != nil {
		fmt.Println(method)
	}

	if resArr, ok := res.([]interface{}); ok {
		fmt.Println("len of keys is ", len(resArr))
		for i, val := range resArr {
			if str, ok := val.([]byte); ok {

				delkey := string(str)
				res, err := models.RedisZeroArgs("del", delkey)
				if err != nil {
					fmt.Println("error ", err)
				}
				fmt.Printf("delete id %-5d key %-40s. res is  %-30d\n", i, string(str), res)
			}
		}
	} else {
		fmt.Println("res is ", res)
	}

}
