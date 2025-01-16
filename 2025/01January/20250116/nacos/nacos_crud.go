package main

import (
	"fmt"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
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

	// 创建配置客户端
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc, // 客户端配置
			ServerConfigs: sc,  // 服务器配置切片
		},
	)
	if err != nil {
		panic(err) // 如果创建客户端失败，程序终止并打印错误信息
	}

	// 发布配置
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data",    // 配置的数据ID
		Group:   "test-group",   // 配置的分组
		Content: "hello world!", // 配置的内容
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err) // 如果发布配置失败，打印错误信息
	}

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data-2",  // 第二个配置的数据ID
		Group:   "test-group",   // 第二个配置的分组
		Content: "hello world!", // 第二个配置的内容
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err) // 如果发布配置失败，打印错误信息
	}

	// 等待1秒，确保配置发布完成
	time.Sleep(1 * time.Second)

	// 获取配置
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "test-data",  // 要获取的配置的数据ID
		Group:  "test-group", // 要获取的配置的分组
	})
	fmt.Println("GetConfig,config :" + content) // 打印获取到的配置内容

	// 监听配置变更
	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-data",  // 要监听的配置的数据ID
		Group:  "test-group", // 要监听的配置的分组
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data) // 配置变更时打印信息
		},
	})
	if err != nil {
		fmt.Printf("ListenConfig err:%+v \n", err) // 如果监听失败，打印错误信息
	}

	err = client.ListenConfig(vo.ConfigParam{
		DataId: "test-data-2", // 第二个要监听的配置的数据ID
		Group:  "test-group",  // 第二个要监听的配置的分组
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data) // 配置变更时打印信息
		},
	})
	if err != nil {
		fmt.Printf("ListenConfig err:%+v \n", err) // 如果监听失败，打印错误信息
	}

	// 等待1秒，确保监听配置生效
	time.Sleep(1 * time.Second)

	// 修改配置
	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data",   // 要修改的配置的数据ID
		Group:   "test-group",  // 要修改的配置的分组
		Content: "test-listen", // 修改后的配置内容
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err) // 如果修改配置失败，打印错误信息
	}

	_, err = client.PublishConfig(vo.ConfigParam{
		DataId:  "test-data-2", // 第二个要修改的配置的数据ID
		Group:   "test-group",  // 第二个要修改的配置的分组
		Content: "test-listen", // 第二个修改后的配置内容
	})
	if err != nil {
		fmt.Printf("PublishConfig err:%+v \n", err) // 如果修改配置失败，打印错误信息
	}

	// 等待2秒，确保配置修改被监听到
	time.Sleep(2 * time.Second)

	// 等待1秒
	time.Sleep(1 * time.Second)

	// 删除配置
	_, err = client.DeleteConfig(vo.ConfigParam{
		DataId: "test-data",  // 要删除的配置的数据ID
		Group:  "test-group", // 要删除的配置的分组
	})
	if err != nil {
		fmt.Printf("DeleteConfig err:%+v \n", err) // 如果删除配置失败，打印错误信息
	}

	// 等待1秒，确保配置删除操作完成
	time.Sleep(1 * time.Second)

	// 取消监听配置变更
	err = client.CancelListenConfig(vo.ConfigParam{
		DataId: "test-data",  // 要取消监听的配置的数据ID
		Group:  "test-group", // 要取消监听的配置的分组
	})
	if err != nil {
		fmt.Printf("CancelListenConfig err:%+v \n", err) // 如果取消监听失败，打印错误信息
	}

	// 搜索配置
	searchPage, _ := client.SearchConfig(vo.SearchConfigParam{
		Search:   "blur", // 搜索模式，这里是模糊搜索
		DataId:   "",     // 数据ID，为空表示不指定
		Group:    "",     // 分组，为空表示不指定
		PageNo:   1,      // 搜索的页码
		PageSize: 10,     // 每页的大小
	})
	fmt.Printf("Search config:%+v \n", searchPage) // 打印搜索结果
}
