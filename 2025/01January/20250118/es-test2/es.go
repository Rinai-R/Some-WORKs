package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Goods struct {
	Id         string  `json:"id,omitempty"`
	Avatar     string  `json:"avatar,omitempty"`
	Goods_name string  `json:"goods_name,omitempty"`
	Shop_id    string  `json:"shop_id,omitempty"`
	Content    string  `json:"content,omitempty"`
	Type       string  `json:"Type,omitempty"`
	Number     int     `json:"number,omitempty"`
	Price      float64 `json:"price,omitempty"`
	Star       int     `json:"star,omitempty"`
	Score      float64
}

var es *elasticsearch.Client
var db *gorm.DB

func init() {
	var err error
	es, err = elasticsearch.NewClient(
		elasticsearch.Config{
			Addresses: []string{
				"http://192.168.195.129:9200",
			},
		},
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	_, err = es.Ping()
	if err != nil {
		log.Fatalf("Error pinging the elastic client: %s", err)
	}
	fmt.Println("Connected to elasticsearch")

	db, err = gorm.Open(mysql.Open("root:~Cy710822@tcp(127.0.0.1:3306)/shopping"))
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

func main() {
	IndexName := "goods"
	goods := Goods{
		Id:         "3",
		Goods_name: "部分数据修改",
	}
	tx := db.Begin()
	err := tx.Model(&Goods{}).Where("id = ?", goods.Id).Updates(goods).Error
	if err != nil {
		tx.Rollback()
		log.Println("Error updating the goods: ", err)
		return
	}
	data, err := json.Marshal(map[string]interface{}{
		"doc": goods,
	})
	if err != nil {
		tx.Rollback()
		fmt.Println("Error marshaling data: %w", err)
		return
	}
	req := esapi.UpdateRequest{
		Index:      IndexName,
		DocumentID: goods.Id,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error updating the goods: ", err)
		return
	}
	defer res.Body.Close()
	if res.IsError() {
		tx.Rollback()
		fmt.Println("Error updating the goods: ", res.Status())
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		fmt.Println("Error commiting the transaction: ", err)
		return
	}
	return
}
