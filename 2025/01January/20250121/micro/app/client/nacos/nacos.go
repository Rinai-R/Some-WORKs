package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var Client naming_client.INamingClient

func init() {
	// 创建ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			"192.168.195.129",                  // Nacos服务器的IP地址
			8848,                               // Nacos服务器的端口号
			constant.WithContextPath("/nacos"), // 上下文路径
		),
	}

	// 创建ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("ace1b5fe-80c3-4fab-b89a-625f9ff41093"), // 命名空间ID
		constant.WithTimeoutMs(5000),                                     // 超时时间（毫秒）
		constant.WithNotLoadCacheAtStart(true),                           // 启动时不加载缓存
		constant.WithLogDir("/tmp/nacos/log"),                            // 日志目录
		constant.WithCacheDir("/tmp/nacos/cache"),                        // 缓存目录
		constant.WithLogLevel("debug"),                                   // 日志级别
		constant.WithUsername("nacos"),                                   // 用户名
		constant.WithPassword("nacos"),                               // 密码
	)
	var err error
	// 创建配置客户端
	Client, err = clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err) // 如果客户端创建失败，则抛出异常
	}

}

func RegisterServiceInstance(client naming_client.INamingClient, param vo.RegisterInstanceParam) {
	// 调用 RegisterInstance 方法注册服务实例
	success, err := client.RegisterInstance(param)
	if !success || err != nil {
		// 如果注册失败，抛出 panic 并打印错误信息
		panic("RegisterServiceInstance failed!" + err.Error())
	}
	// 打印注册参数和结果
	fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}

func DeRegisterServiceInstance(client naming_client.INamingClient, param vo.DeregisterInstanceParam) {
	// 调用 DeregisterInstance 方法取消注册服务实例
	success, err := client.DeregisterInstance(param)
	if !success || err != nil {
		// 如果取消注册失败，抛出 panic 并打印错误信息
		panic("DeRegisterServiceInstance failed!" + err.Error())
	}
	// 打印取消注册参数和结果
	fmt.Printf("DeRegisterServiceInstance,param:%+v,result:%+v \n\n", param, success)
}
