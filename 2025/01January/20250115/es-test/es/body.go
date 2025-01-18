package main

var mappingTpl string

func init() {
	mappingTpl = `{
  "mappings": {
    "properties": {
      "all": {
        "type": "text",
        "analyzer": "ik_max_word"
      },
      "id": {
        "type": "keyword"
      },
      "goods_name": {
        "type": "text",
        "analyzer": "ik_max_word",
        "copy_to": "all"
      },
      "avatar": {
        "type": "keyword",
        "index": false
      },
      "shop_id": {
        "type": "keyword"
      },
      "content": {
        "type": "text",
        "analyzer": "ik_max_word",
        "copy_to": "all"
      },
      "star": {
        "type": "integer"
      },
      "price": {
        "type": "float"
      },
      "number": {
        "type": "integer"
      },
      "Type": {
        "type": "keyword",
        "copy_to": "all"
      }
    }
  }
}`

}
