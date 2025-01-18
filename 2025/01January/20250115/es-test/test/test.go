package main

import (
	"Golang/2025/01January/20250115/es-test/model"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

func main() {
	var key string
	_, err := fmt.Scan(&key)
	if err != nil {
		return
	}
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.195.129:9200"),
		elastic.SetSniff(false),       // 禁用嗅探
		elastic.SetHealthcheck(false), // 禁用健康检查
	)

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 检查连接是否成功
	_, _, err = client.Ping("http://192.168.195.129:9200/").Do(context.Background())
	if err != nil {
		log.Fatalf("Error pinging: %s", err)
	}
	fmt.Println("Elasticsearch连接成功")
	// 定义查询
	matchQuery := elastic.NewMatchQuery("all", key)

	// 执行查询
	searchResult, err := client.Search().
		Index("goods"). // 查询的索引
		Query(matchQuery). // 使用的查询
		TrackTotalHits(true). // 跟踪总命中数
		Size(100).
		Do(context.Background()) // 执行查询
	if err != nil {
		log.Fatalf("Error executing search: %s", err)
	}

	// 处理响应
	totalHits := searchResult.Hits.TotalHits.Value // 获取总命中数
	fmt.Printf("Total hits: %d\n", totalHits)

	var ans []model.Goods

	// 遍历搜索结果的 Hits
	for _, hit := range searchResult.Hits.Hits {
		var goods model.Goods

		// 解码 hit.Source 直接到结构体
		err = json.Unmarshal(hit.Source, &goods)
		if err != nil {
			log.Fatalf("Error binding JSON to struct: %s", err)
		}
		ans = append(ans, goods)
	}

	// 打印或返回产品信息
	for _, f := range ans {
		fmt.Println(f)
	}
}
