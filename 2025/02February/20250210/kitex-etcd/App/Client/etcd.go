package Client

import (
	"Golang/2025/02February/20250210/kitex-etcd/Logger"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

type Registry struct {
	client   *clientv3.Client
	Services map[string]string
}

var ETCD *Registry

func init() {
	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.195.129:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Panicln("初始化注册中心失败 " + err.Error())
	}
	ETCD = &Registry{
		client: etcd,
	}
}

func (client *Registry) DiscoverService(serviceName string) error {
	resp, err := client.client.Get(context.Background(), serviceName, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, v := range resp.Kvs {
		msg := fmt.Sprintf("update: " + string(v.Key) + " " + string(v.Value))
		Logger.Logger.Info(msg)
		client.UpdateService(string(v.Key), string(v.Value))
	}

	// 监听服务的变化
	rch := client.client.Watch(context.Background(), serviceName)
	go func() {
		for watchResponse := range rch {
			for _, ev := range watchResponse.Events {
				if ev.Type == clientv3.EventTypePut {
					// 服务注册（新增或更新）
					msg := fmt.Sprintf("update: " + string(ev.Kv.Key) + " " + string(ev.Kv.Value))
					Logger.Logger.Info(msg)
				} else if ev.Type == clientv3.EventTypeDelete {
					// 服务注销
					msg := fmt.Sprintf("delete: " + string(ev.Kv.Key))
					Logger.Logger.Info(msg)
				}
			}
		}
	}()
	return nil
}

func (client *Registry) UpdateService(key string, value string) {
	client.Services[key] = value
}

func (client *Registry) DeleteService(serviceName string) {
	delete(client.Services, serviceName)
}

func (client *Registry) GetService(serviceName string) string {
	value, ok := client.Services[serviceName]
	if !ok {
		return ""
	}
	return value
}
