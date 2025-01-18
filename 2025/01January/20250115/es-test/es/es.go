package main

import (
	"Golang/2025/01January/20250115/es-test/model"
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"

	_ "github.com/go-sql-driver/mysql"
)

var client *elastic.Client
var db *sql.DB
var err error

func init() {
	// 创建Elasticsearch客户端
	client, err = elastic.NewClient(
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

	db, err = sql.Open("mysql", "root:~Cy710822@tcp(127.0.0.1:3306)/shopping?parseTime=True&loc=UTC")
	if err != nil {
		panic(err)
	}
}

func main() {
	IndexName := "goods"
	exists, err := client.IndexExists(IndexName).Do(context.Background())
	if err != nil {
		log.Fatalf("Error checking if index exists: %s", err)
		return
	}

	if !exists {
		_, err2 := client.CreateIndex(IndexName).Body(mappingTpl).Do(context.Background())
		if err2 != nil {
			return
		}
	}

	rows, err := db.Query("SELECT id, avatar, goods_name, shop_id, content, type, number, price, star FROM goods")
	if err != nil {
		log.Fatalf("Error querying database: %s", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("Error closing rows: %s", err)
		}
	}(rows)
	for rows.Next() {
		var goods model.Goods
		if err := rows.Scan(&goods.Id, &goods.Avatar, &goods.Goods_name, &goods.Shop_id,
			&goods.Content, &goods.Type, &goods.Number, &goods.Price, &goods.Star); err != nil {
			log.Fatalf("Error scanning row: %s", err)
		}

		// 将数据插入到 Elasticsearch
		response, err := client.Index().
			Index(IndexName).
			Id(goods.Id). // 使用数据库中唯一标识符作为ID
			BodyJson(goods).
			Do(context.Background())
		if err != nil {
			log.Printf("Error indexing document: %s", err)
		} else {
			log.Printf("Document inserted with ID: %s", response.Id)

			// 查询插入的数据
			getResponse, err := client.Get().
				Index(IndexName).
				Id(goods.Id). // 这里确保传递了有效的 ID
				Do(context.Background())
			if err != nil {
				log.Printf("Error getting document: %s", err)
			} else {
				if getResponse.Found {
					log.Printf("Document found: %s", getResponse.Source) // 可以对 Source 做 JSON 反序列化以便读取
				} else {
					log.Printf("Document not found")
				}
			}
		}
	}

}
