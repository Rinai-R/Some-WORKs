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
	"sync/atomic"
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

var sum = 0

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.195.129:6379",
		Password: "~Cy710822",
		DB:       0,
	})
}

var l MyLock

func dis_Lock(Key string) {
	for {
		locked, _ := rdb.SetNX(context.Background(), Key, 1, time.Second).Result()
		if locked {
			break
		}
		continue
	}
}
func dis_Unlock(Key string) {
	rdb.Del(context.Background(), Key)
}

func (s *service) handleRequest_Group(ctx context.Context, request *Request) (*Response, error) {
	str, _ := json.Marshal(request)
	v, err, _ := s.requestGroup.Do(string(str), func() (interface{}, error) {
		sum += 1
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
func (s *service) handleRequest2(ctx context.Context, request *Request) (*Response, error) {
	cacheKey := fmt.Sprintf("user:%d", request.user_id)
	res, err := rdb.HGetAll(ctx, cacheKey).Result()
	if len(res) == 0 {
		dis_Lock(cacheKey)
		defer dis_Unlock(cacheKey)
		//l.Lock()
		//defer l.Unlock()
		//lock.Lock()
		//defer lock.Unlock()
		sum += 1
		return &Response{
			username: "sql查询",
			password: "sql查询",
		}, nil

	}
	if err != nil {
		return nil, err
	}
	return &Response{
		username: res["name"],
		password: res["password"],
	}, nil
}

func (s *service) handleRequest(ctx context.Context, request *Request) (*Response, error) {
	cacheKey := fmt.Sprintf("user:%d", request.user_id)
	res, err := rdb.HGetAll(ctx, cacheKey).Result()
	if len(res) == 0 {
		LockKey := "locked:" + cacheKey
		locked, _ := rdb.SetNX(ctx, LockKey, "locked", time.Second*10).Result()
		if locked {
			fmt.Println("redis setnx-----------------------------------------------")
			//todo
			//访问数据库，将数据存入缓存
			rdb.Del(ctx, LockKey)
			sum += 1
			return &Response{
				username: "sql查询",
				password: "sql查询",
			}, nil
		} else {
			//1000并发：不等待：13.679386s
			//采取随机数等待：4.0750594s
			//若细化等待时间粒度，在0~100000microsecond时，效率最高：826.1751ms
			//万级并发量有点玄学了，果然只能找其他方法了吗...
			t := rand.Int() % 100000
			time.Sleep(time.Microsecond * time.Duration(t))
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
	begin := time.Now()
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done() // 确保每个 goroutine 都调用 wg.Done()

			response, err := s.handleRequest2(ctx, &Request{
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
	defer fmt.Println(time.Since(begin))
	defer fmt.Println(sum)
}

var lock sync.Mutex

// μs单位，相似了,byd,setnx不如lock
func lockTest() {
	lo.Lock()
	sum += 1
	lo.Unlock()
}

var lo = &MyLock{num: 0}

type MyLock struct {
	num int64
}

func (l *MyLock) Lock() {
	for {
		if atomic.CompareAndSwapInt64(&(l.num), 0, 1) {
			return
		}
	}
}
func (l *MyLock) Unlock() {
	atomic.StoreInt64(&(l.num), 0)
}
