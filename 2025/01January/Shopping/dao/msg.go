package dao

import (
	"Golang/2025/01January/Shopping/model"
	"database/sql"
	"errors"
	"log"
)

func GetGoodsMsg(goods model.Goods) []model.Msg {
	var ans []model.Msg
	query := `select id from msg where goods_id = ? and parent_id IS NULL`

	row, err := db.Query(query, goods.Id, goods.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			return nil
		}
		log.Println(err)
		return nil
	}
	for row.Next() {
		var id string
		err = row.Scan(&id)
		query = `SELECT id, goods_id, user_id, content, praised_num, create_at, updated_at FROM msg WHERE id = ?`
		var MainMsg model.Msg
		err = db.QueryRow(id).Scan(&MainMsg.Id, &MainMsg.Goods_id, &MainMsg.User_id, &MainMsg.Content, &MainMsg.Praised_num, &MainMsg.Create_at, &MainMsg.Updated_at)
		if err != nil {
			log.Println(err)
			return nil
		}
		MainMsg.Response = append(MainMsg.Response, InOrder(ans, MainMsg.Id)...)
		ans = append(ans, MainMsg)
	}
	return ans
}

func InOrder(ans []model.Msg, parent_id int) []model.Msg {
	query := `SELECT id FROM msg WHERE parent_id = id`
	rows, err := db.Query(query, parent_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		log.Println(err)
		return nil
	}
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		var message model.Msg
		query = `SELECT id, parent_id, goods_id, user_id, content, praised_num, create_at, updated_at FROM msg WHERE id = ?`
		err = db.QueryRow(query, id).Scan(&message.Id, &message.Parent_id, &message.User_id, &message.Content, &message.Praised_num, &message.Create_at, &message.Updated_at)
		if err != nil {
			log.Println(err)
			return nil
		}
		message.Response = append(message.Response, InOrder(ans, message.Id)...)
		ans = append(ans, message)
	}
	return ans
}

// PubMsg 发布评论
func PubMsg(msg model.Msg) bool {
	query := `INSERT INTO msg (goods_id, user_id, content) values (?, ?, ?) `
	_, err := db.Exec(query, msg.Goods_id, msg.User_id, msg.Content)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Response(msg model.Msg) bool {
	query := `INSERT INTO msg (goods_id, user_id, content, parent_id) values (?, ?, ?, ?) `
	_, err := db.Exec(query, msg.Goods_id, msg.User_id, msg.Content, msg.Parent_id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Praise(msg model.Msg, user model.User) bool {
	query := `UPDATE msg SET praised_num = praised_num + 1 WHERE id = ?`
	_, err := db.Exec(query, msg.Id)
	if err != nil {
		log.Println(err)
		return false
	}
	query = `INSERT INTO praise (user_id, message_id) values (?, ?)`
	_, err = db.Exec(query, user.Id, msg.Id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func AlterMsg(msg model.Msg) bool {
	query := `UPDATE msg SET content = ? WHERE id = ?`
	_, err := db.Exec(query, msg.Content, msg.Id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func DelMsg(msg model.Msg) bool {
	query := `DELETE FROM msg WHERE id = ?`
	_, err := db.Exec(query, msg.Id)
	if err != nil {
		return false
	}
	return true
}

func PraiseMsg(praise model.Praise) bool {
	query := `INSERT INTO praise (user_id, message_id) values(?, ?)`
	_, err := db.Exec(query, praise.User_id, praise.Message_id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
