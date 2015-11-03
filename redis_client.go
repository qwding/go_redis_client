package main

/**
 * @author carlding
 * @date : 2015-08
 */
import (
	"flag"
	"fmt"
	// "github.com/garyburd/redigo/redis"
	"go_redis_client/models"
	"strings"
	// "encoding/json"
	"reflect"
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
	fmt.Printf("2. Setkeyskv. hmset multiply hash keys with key and value, use like 'setkeykv 2015_sample_* mykey myvalue'.\n")
	fmt.Printf("3. getjson. if the key's value is json,use this command will print it string.\n")
}

func main() {
	fmt.Println("input redis command or self command,input help see detail.")
	for {
		fmt.Printf(">> ")

		cmd := ReadLine()
		switch strings.TrimSpace(cmd) {
		case "":
			continue
		case "help":
			showSelfCmd()
			continue
		case "exit":
		case "quit":
			return
		}

		arr := strings.Split(cmd, " ")
		for i, _ := range arr {
			arr[i] = strings.TrimSpace(arr[i])
		}

		var res interface{}
		var err error

		length := len(arr)
		m := MyMethod{}
		/*		if length <= 0 {
					continue
				} else if length == 1 {
					res, err = models.RedisOnlyCmd(arr[0])
				} else if length == 2 {
					res, err = models.RedisRun(arr[0], arr[1])
				} else {
					switch arr[0] {
					case "delkeys":
						fmt.Println("delete keys")
					default:
						res, err = models.RedisRunArgs(arr[0], arr[1], arr[2:])
					}
				}*/
		if length <= 0 {
			continue
		} else if m.IsFunc(arr[0]) {
			f, ok := m.GetFunc(arr[0])
			if !ok {
				fmt.Println("Error! Can't find the function.")
			}
			if length == 1 {
				fmt.Println("have not have one args function.")
			} else {
				//即使arr的长度只有2，那么这里是不会报错的。而且正常调用
				res, err = f(arr[1], arr[2:]...)
			}
		} else {
			if length == 1 {
				res, err = models.RedisOnlyCmd(arr[0])
			} else if length == 2 {
				res, err = models.RedisRun(arr[0], arr[1])
			} else if length == 3 {
				res, err = models.RedisRun(arr[0], arr[1], arr[2])
			} else {
				res, err = models.RedisRun(arr[0], arr[1], arr[2:])
			}
		}
		if err != nil {
			fmt.Println("request error", err)
		}

		if resArr, ok := res.([]interface{}); ok {
			fmt.Println("len of res is ", len(resArr))
			for j, val := range resArr {
				if str, ok := val.([]byte); ok {
					fmt.Println("res interface to string is :", j, string(str))
				}
				if strArr, ok := val.([]interface{}); ok {
					for i, v := range strArr {
						fmt.Println("res to array ", j, i, string(v.([]byte)))
					}

				}
			}
		} else {
			if strByte, ok := res.([]byte); ok {
				fmt.Println("res byte to string is :", string(strByte))
			} else {
				fmt.Println("res:")
				fmt.Println(res)
			}
		}
	}
}

type MyMethod struct {
}

func (this *MyMethod) IsFunc(str string) bool {
	str = strings.ToUpper(str[0:1]) + strings.ToLower(str[1:])
	method := reflect.ValueOf(this).MethodByName(str)

	return method.IsValid()
}

func (this *MyMethod) GetFunc(str string) (func(string, ...string) (interface{}, error), bool) {
	str = strings.ToUpper(str[0:1]) + strings.ToLower(str[1:])
	method := reflect.ValueOf(this).MethodByName(str)
	if !method.IsValid() {
		return nil, false
	}
	return method.Interface().(func(string, ...string) (interface{}, error)), true
}

func (this *MyMethod) Delkeys(key string, _ ...string) (interface{}, error) {
	// method := "delkeys"
	res, err := models.RedisRun("keys", key)
	if err != nil {
		return "delkeys error", err
	}

	if resArr, ok := res.([]interface{}); ok {
		fmt.Println("len of keys is ", len(resArr))
		for i, val := range resArr {
			if str, ok := val.([]byte); ok {

				delkey := string(str)
				res, err := models.RedisRun("del", delkey)
				if err != nil {
					return "delkeys errors", err
				}
				fmt.Printf("delete id %-5d key %-60s. res is  %-30d\n", i, string(str), res)
			}
		}
	} else {
		fmt.Println("result: ", res)
	}
	return "delkeys success", nil
}

func (this *MyMethod) Setkeyskv(key string, arr ...string) (interface{}, error) {
	method := "SetkeysKV"
	res, err := models.RedisRun("keys", key)
	if err != nil {
		return "Setkeyskv error", err
	}

	if resArr, ok := res.([]interface{}); ok {
		fmt.Println("len of keys is ", len(resArr))
		for i, val := range resArr {
			if str, ok := val.([]byte); ok {

				unitkey := string(str)
				res, err := models.RedisRun("hmset", unitkey, arr)
				if err != nil {
					return method + "errors", err
				}
				fmt.Printf(method+"id %-5d key %-60s. res is  %-30s\n", i, string(str), res)
			}
		}
	} else {
		return method + " failed", nil
	}
	return method + "success", nil
}

func (this *MyMethod) Getjson(key string, args ...string) (interface{}, error) {
	method := "SetkeysKV"
	if len(args) > 0 {
		return method + " error", fmt.Errorf("this command has none args.")
	}
	res, err := models.RedisRun("get", key)
	if err != nil {
		return "GetJson error", err
	}

	if resArr, ok := res.([]byte); ok {
		return string(resArr), nil
	} else {
		return method + " failed", nil
	}
}
