package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

var Client *elastic.Client
var err error

type Goods struct {
	Id         string  `json:"id"`
	Avatar     string  `json:"avatar"`
	Goods_name string  `json:"goods_name"`
	Shop_id    string  `json:"shop_id"`
	Content    string  `json:"content"`
	Type       string  `json:"Type"`
	Number     int     `json:"number"`
	Price      float64 `json:"price"`
	Star       int     `json:"star"`
	Score      float64
}

func init() {
	Client, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("ES ok")
}
func main() {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}
	query := elastic.NewMatchQuery("all", s)
	searchResult, err := Client.Search().Index("goods").Query(query).TrackTotalHits(true).Size(100).Do(context.Background())
	if err != nil {
		panic(err)
	}

	// 处理响应
	totalHits := searchResult.Hits.TotalHits.Value // 获取总命中数
	fmt.Printf("Total hits: %d\n", totalHits)

	var ans []Goods

	// 遍历搜索结果的 Hits
	for _, hit := range searchResult.Hits.Hits {

		var goods Goods

		// 解码 hit.Source 直接到结构体
		err = json.Unmarshal(hit.Source, &goods)
		if err != nil {
			log.Fatalf("Error binding JSON to struct: %s", err)
		}
		goods.Score = *hit.Score
		ans = append(ans, goods)
	}

	// 打印或返回产品信息
	for _, f := range ans {
		fmt.Println(f)
	}
}
