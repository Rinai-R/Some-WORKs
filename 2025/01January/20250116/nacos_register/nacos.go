package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/util"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	// 创建 ServerConfig，配置 Nacos 服务器地址、端口及上下文路径
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("192.168.195.129", 8848, constant.WithContextPath("/nacos")),
	}

	// 创建 ClientConfig，配置客户端的基本参数，如命名空间、超时时间、日志路径等
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("ace1b5fe-80c3-4fab-b89a-625f9ff41093"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
		constant.WithUsername("nacos"),     // 用户名
		constant.WithPassword("nacos"), // 密码
	)

	// 创建命名服务客户端
	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err) // 如果客户端创建失败，则抛出异常
	}

	// 注册服务实例到 Nacos
	registerServiceInstance(client, vo.RegisterInstanceParam{
		Ip:          "10.0.0.10",                          // 服务实例的 IP 地址
		Port:        8848,                                 // 服务实例的端口号
		ServiceName: "test.go",                            // 服务名称
		GroupName:   "group-a",                            // 分组名称
		ClusterName: "cluster-a",                          // 集群名称
		Weight:      10,                                   // 权重
		Enable:      true,                                 // 是否启用
		Healthy:     true,                                 // 是否健康
		Ephemeral:   true,                                 // 是否为临时实例
		Metadata:    map[string]string{"idc": "shanghai"}, // 元数据信息
	})

	//从 Nacos 取消注册服务实例
	deRegisterServiceInstance(client, vo.DeregisterInstanceParam{
		Ip:          "10.0.0.10", // 服务实例的 IP 地址
		Port:        8848,        // 服务实例的端口号
		ServiceName: "demo.go",   // 服务名称
		GroupName:   "group-a",   // 分组名称
		Cluster:     "cluster-a", // 集群名称
		Ephemeral:   true,        // 必须为临时实例
	})

	time.Sleep(1 * time.Second) // 等待 1 秒

	// 批量注册多个服务实例
	batchRegisterServiceInstance(client, vo.BatchRegisterInstanceParam{
		ServiceName: "demo.go", // 服务名称
		GroupName:   "group-a", // 分组名称
		Instances: []vo.RegisterInstanceParam{{
			Ip:          "10.0.0.10",                          // 第一个服务实例的 IP 地址
			Port:        8848,                                 // 第一个服务实例的端口号
			Weight:      10,                                   // 权重
			Enable:      true,                                 // 是否启用
			Healthy:     true,                                 // 是否健康
			Ephemeral:   true,                                 // 是否为临时实例
			ClusterName: "cluster-a",                          // 集群名称
			Metadata:    map[string]string{"idc": "shanghai"}, // 元数据信息
		}, {
			Ip:          "10.0.0.12",                          // 第二个服务实例的 IP 地址
			Port:        8848,                                 // 第二个服务实例的端口号
			Weight:      7,                                    // 权重
			Enable:      true,                                 // 是否启用
			Healthy:     true,                                 // 是否健康
			Ephemeral:   true,                                 // 是否为临时实例
			ClusterName: "cluster-a",                          // 集群名称
			Metadata:    map[string]string{"idc": "shanghai"}, // 元数据信息
		}},
	})

	time.Sleep(1 * time.Second) // 等待 1 秒

	// 根据服务名称、分组名称和集群名称获取服务信息
	getService(client, vo.GetServiceParam{
		ServiceName: "demo.go",             // 服务名称
		GroupName:   "group-a",             // 分组名称
		Clusters:    []string{"cluster-a"}, // 集群名称列表
	})

	// 获取指定服务的所有实例
	selectAllInstances(client, vo.SelectAllInstancesParam{
		ServiceName: "demo.go",             // 服务名称
		GroupName:   "group-a",             // 分组名称
		Clusters:    []string{"cluster-a"}, // 集群名称列表
	})

	// 获取指定服务的健康实例
	selectInstances(client, vo.SelectInstancesParam{
		ServiceName: "demo.go",             // 服务名称
		GroupName:   "group-a",             // 分组名称
		Clusters:    []string{"cluster-a"}, // 集群名称列表
		HealthyOnly: true,                  // 仅获取健康实例
	})

	// 根据加权随机算法获取一个健康实例
	selectOneHealthyInstance(client, vo.SelectOneHealthInstanceParam{
		ServiceName: "demo.go",             // 服务名称
		GroupName:   "group-a",             // 分组名称
		Clusters:    []string{"cluster-a"}, // 集群名称列表
	})

	// 订阅服务变更，当服务信息发生变化时，会触发回调函数
	subscribeParam := &vo.SubscribeParam{
		ServiceName: "demo.go", // 服务名称
		GroupName:   "group-a", // 分组名称
		SubscribeCallback: func(services []model.Instance, err error) {
			fmt.Printf("callback return services:%s \n\n", util.ToJsonString(services)) // 变更时打印服务实例信息
		},
	}
	subscribe(client, subscribeParam)

	// 等待 3 秒，让客户端从服务端拉取变更
	time.Sleep(3 * time.Second)

	// 更新服务实例信息
	updateServiceInstance(client, vo.UpdateInstanceParam{
		Ip:          "10.0.0.11",                          // 更新后的 IP 地址
		Port:        8848,                                 // 服务实例的端口号
		ServiceName: "demo.go",                            // 服务名称
		GroupName:   "group-a",                            // 分组名称
		ClusterName: "cluster-a",                          // 集群名称
		Weight:      10,                                   // 权重
		Enable:      true,                                 // 是否启用
		Healthy:     true,                                 // 是否健康
		Ephemeral:   true,                                 // 是否为临时实例
		Metadata:    map[string]string{"idc": "beijing1"}, // 更新后的元数据信息
	})

	// 等待 3 秒，让客户端从服务端拉取变更
	time.Sleep(3 * time.Second)

	// 取消订阅服务变更
	unSubscribe(client, subscribeParam)

	// 获取指定分组下的所有服务名称列表
	getAllService(client, vo.GetAllServiceInfoParam{
		GroupName: "group-a", // 分组名称
		PageNo:    1,         // 分页页码
		PageSize:  10,        // 每页大小
	})
}

//==========================================以下为函数实现===================================================

// registerServiceInstance 向 Nacos 注册一个服务实例
func registerServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	// 调用 RegisterInstance 方法注册服务实例
	success, err := client.RegisterInstance(param)
	if !success || err != nil {
		// 如果注册失败，抛出 panic 并打印错误信息
		panic("RegisterServiceInstance failed!" + err.Error())
	}
	// 打印注册参数和结果
	fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

// batchRegisterServiceInstance 向 Nacos 批量注册多个服务实例
func batchRegisterServiceInstance(client naming_client.INamingClient, param vo.BatchRegisterInstanceParam) {
	// 调用 BatchRegisterInstance 方法批量注册服务实例
	success, err := client.BatchRegisterInstance(param)
	if !success || err != nil {
		// 如果批量注册失败，抛出 panic 并打印错误信息
		panic("BatchRegisterServiceInstance failed!" + err.Error())
	}
	// 打印批量注册参数和结果
	fmt.Printf("BatchRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

// deRegisterServiceInstance 从 Nacos 取消注册一个服务实例
func deRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	// 调用 DeregisterInstance 方法取消注册服务实例
	success, err := client.DeregisterInstance(param)
	if !success || err != nil {
		// 如果取消注册失败，抛出 panic 并打印错误信息
		panic("DeRegisterServiceInstance failed!" + err.Error())
	}
	// 打印取消注册参数和结果
	fmt.Printf("DeRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

// updateServiceInstance 更新 Nacos 中已注册的服务实例信息
func updateServiceInstance(client naming_client.INamingClient, param vo.UpdateInstanceParam) {
	// 调用 UpdateInstance 方法更新服务实例信息
	success, err := client.UpdateInstance(param)
	if !success || err != nil {
		// 如果更新失败，抛出 panic 并打印错误信息
		panic("UpdateInstance failed!" + err.Error())
	}
	// 打印更新参数和结果
	fmt.Printf("UpdateServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

// getService 从 Nacos 获取指定服务的信息
func getService(client naming_client.INamingClient, param vo.GetServiceParam) {
	// 调用 GetService 方法获取服务信息
	service, err := client.GetService(param)
	if err != nil {
		// 如果获取服务信息失败，抛出 panic 并打印错误信息
		panic("GetService failed!" + err.Error())
	}
	// 打印获取服务信息的参数和服务信息结果
	fmt.Printf("GetService,param:%+v, result:%+v \n\n", param, service)
}

// selectAllInstances 从 Nacos 获取指定服务的所有实例
func selectAllInstances(client naming_client.INamingClient, param vo.SelectAllInstancesParam) {
	// 调用 SelectAllInstances 方法获取所有实例
	instances, err := client.SelectAllInstances(param)
	if err != nil {
		// 如果获取所有实例失败，抛出 panic 并打印错误信息
		panic("SelectAllInstances failed!" + err.Error())
	}
	// 打印获取所有实例的参数和实例信息结果
	fmt.Printf("SelectAllInstance,param:%+v, result:%+v \n\n", param, instances)
}

// selectInstances 从 Nacos 获取指定服务的健康实例
func selectInstances(client naming_client.INamingClient, param vo.SelectInstancesParam) {
	// 调用 SelectInstances 方法获取健康实例
	instances, err := client.SelectInstances(param)
	if err != nil {
		// 如果获取健康实例失败，抛出 panic 并打印错误信息
		panic("SelectInstances failed!" + err.Error())
	}
	// 打印获取健康实例的参数和实例信息结果
	fmt.Printf("SelectInstances,param:%+v, result:%+v \n\n", param, instances)
}

// selectOneHealthyInstance 从 Nacos 获取一个健康实例（使用加权随机算法）
func selectOneHealthyInstance(client naming_client.INamingClient, param vo.SelectOneHealthInstanceParam) {
	// 调用 SelectOneHealthyInstance 方法获取一个健康实例
	instances, err := client.SelectOneHealthyInstance(param)
	if err != nil {
		// 如果获取健康实例失败，抛出 panic 并打印错误信息
		panic("SelectOneHealthyInstance failed!")
	}
	// 打印获取健康实例的参数和实例信息结果
	fmt.Printf("SelectOneHealthyInstance,param:%+v, result:%+v \n\n", param, instances)
}

// subscribe 订阅 Nacos 中指定服务的变化
func subscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	// 调用 Subscribe 方法订阅服务变化
	client.Subscribe(param)
}

// unSubscribe 取消订阅 Nacos 中指定服务的变化
func unSubscribe(client naming_client.INamingClient, param *vo.SubscribeParam) {
	// 调用 Unsubscribe 方法取消订阅服务变化
	client.Unsubscribe(param)
}

// getAllService 从 Nacos 获取指定分组下的所有服务名称列表
func getAllService(client naming_client.INamingClient, param vo.GetAllServiceInfoParam) {
	// 调用 GetAllServicesInfo 方法获取所有服务名称列表
	service, err := client.GetAllServicesInfo(param)
	if err != nil {
		// 如果获取服务名称列表失败，抛出 panic 并打印错误信息
		panic("GetAllService failed!")
	}
	// 打印获取服务名称列表的参数和服务名称列表结果
	fmt.Printf("GetAllService,param:%+v, result:%+v \n\n", param, service)
}
