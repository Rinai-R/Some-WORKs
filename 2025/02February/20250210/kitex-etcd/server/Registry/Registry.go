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

var err error

func init() {
	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.195.129:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Panicln("初始化注册中心失败 " + err.Error())
	}
	EtcdRegistry = &Registry{etcd}
}

func (e *Registry) ServiceRegister(name string, add string) {
	_, err := e.client.Put(context.Background(), name, add)
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
