package Registry

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type Registry struct {
	client *clientv3.Client
}

var EtcdRegistry *Registry

func init() {
	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Panic("初始化注册中心失败 " + err.Error())
	}
	EtcdRegistry = &Registry{etcd}
}

func (e *Registry) ServiceRegister(name string, addr string) {
	_, err := e.client.Put(context.Background(), name, addr)
	if err != nil {
		log.Panic(err)
	}
	log.Println(name + " service ok")
}

func (e *Registry) ServiceUnRegister(name string) {
	_, err := e.client.Delete(context.Background(), name)
	if err != nil {
		log.Panic(err)
	}
	log.Println(name + " service ok")
}
