package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type service struct {
	requestGroup singleflight.Group
}

type Request struct {
	user_id int64
}

type Response struct {
	username string
	password string
}

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})
}

func (s *service) handleRequest_Group(ctx context.Context, request *Request) (*Response, error) {
	str, _ := json.Marshal(request)
	v, err, _ := s.requestGroup.Do(string(str), func() (interface{}, error) {
		result, err := rdb.HGetAll(ctx, "user:"+strconv.FormatInt(request.user_id, 10)).Result()
		if err != nil {
			return nil, err
		}
		return &Response{
			username: result["name"],
			password: result["password"],
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return v.(*Response), nil
}

func (s *service) handleRequest(ctx context.Context, request *Request) (*Response, error) {
	cacheKey := fmt.Sprintf("user:%d", request.user_id)
	res, err := rdb.HGetAll(ctx, cacheKey).Result()
	if len(res) == 0 {
		LockKey := "locked:" + cacheKey
		locked, _ := rdb.SetNX(ctx, LockKey, "locked", time.Second*10).Result()
		fmt.Println("locked:", locked)
		if locked {
			fmt.Println("redis setnx-----------------------------------------------")
			//todo
			//访问数据库，将数据存入缓存
			rdb.Del(ctx, LockKey)
			return &Response{
				username: "sql查询",
				password: "sql查询",
			}, nil
		} else {
			t := rand.Int() % 100
			time.Sleep(time.Millisecond * time.Duration(t))
			return s.handleRequest(ctx, request)
		}
	}
	if err != nil {
		return nil, err
	}
	return &Response{
		username: res["name"],
		password: res["password"],
	}, nil
}
func main() {
	var ctx = context.Background()
	wg := &sync.WaitGroup{}
	n := 1000
	wg.Add(n)
	s := &service{requestGroup: singleflight.Group{}}
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done() // 确保每个 goroutine 都调用 wg.Done()

			response, err := s.handleRequest(ctx, &Request{
				user_id: 2,
			})
			if err != nil {
				// 打印错误信息，并返回
				fmt.Println("Error:", err)
				return
			}
			// 如果没有错误，打印 response
			fmt.Println(*response)
		}()
	}
	wg.Wait() // 等待所有 goroutine 完成
}
