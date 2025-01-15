package dao

import (
	"Golang/2025/01January/20250101Shopping/model"
	"database/sql"
	"errors"
	"log"
)

// GetGoodsMsg 查询商品的评论
func GetGoodsMsg(goods model.Goods) []model.Msg {
	var ans []model.Msg
	query := `select id from msg where goods_id = ? and parent_id IS NULL`

	row, err := db.Query(query, goods.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(err)
			return nil
		}
		log.Println(err)
		return nil
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Println(err)
		}
	}(row)
	for row.Next() {
		var id string
		err = row.Scan(&id)
		query = `SELECT id, goods_id, user_id, content, praised_num, create_at, updated_at FROM msg WHERE id = ?`
		var MainMsg model.Msg
		err = db.QueryRow(query, id).Scan(&MainMsg.Id, &MainMsg.Goods_id, &MainMsg.User_id, &MainMsg.Content, &MainMsg.Praised_num, &MainMsg.Create_at, &MainMsg.Updated_at)
		if err != nil {
			log.Println(err)
			return nil
		}
		//采取中序遍历，根据父节点的id查找子节点
		//将儿子（回复该评论的内容）加入response成员切片中。
		MainMsg.Response = InOrder(MainMsg.Id)
		ans = append(ans, MainMsg)
	}
	return ans
}

func InOrder(parent_id string) []model.Msg {
	var ans []model.Msg
	query := `SELECT id FROM msg WHERE parent_id = ?`
	rows, err := db.Query(query, parent_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		log.Println(err)
		return nil
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		var message model.Msg
		query = `SELECT id, parent_id, goods_id, user_id, content, praised_num, create_at, updated_at FROM msg WHERE id = ?`
		err = db.QueryRow(query, id).Scan(&message.Id, &message.Parent_id, &message.Goods_id, &message.User_id, &message.Content, &message.Praised_num, &message.Create_at, &message.Updated_at)
		if err != nil {
			log.Println(err)
			return nil
		}
		//将儿子（回复该评论的内容）加入response成员切片中。
		message.Response = InOrder(message.Id)
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
	query := `SELECT goods_id FROM msg WHERE id = ?`
	err := db.QueryRow(query, msg.Parent_id).Scan(&msg.Goods_id)
	if err != nil {
		log.Println(err)
		return false
	}
	query = `INSERT INTO msg (goods_id, user_id, content, parent_id) values (?, ?, ?, ?) `
	_, err = db.Exec(query, msg.Goods_id, msg.User_id, msg.Content, msg.Parent_id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Praise 点赞评论，第二次点击为取消点赞
func Praise(praise model.Praise) bool {
	query := `SELECT 1 FROM praise WHERE message_id = ?  AND user_id = ?`
	var exist bool
	err := db.QueryRow(query, praise.Message_id, praise.User_id).Scan(&exist)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return false
	}
	tx, err := db.Begin()
	if !exist {
		query = `UPDATE msg SET praised_num = praised_num + 1 WHERE id = ?`
		_, err = tx.Exec(query, praise.Message_id)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				return false
			}
			log.Println(err)
			return false
		}
		query = `INSERT INTO praise (user_id, message_id) values (?, ?)`
		_, err = tx.Exec(query, praise.User_id, praise.Message_id)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				return false
			}
			log.Println(err)
			return false
		}
	} else {
		query = `UPDATE msg SET praised_num = praised_num - 1 WHERE id = ?`
		_, err = tx.Exec(query, praise.Message_id)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				return false
			}
			log.Println(err)
			return false
		}
		query = `DELETE FROM praise  WHERE message_id = ? AND user_id = ?`
		_, err = tx.Exec(query, praise.Message_id, praise.User_id)
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				return false
			}
			log.Println(err)
			return false
		}
	}
	err = tx.Commit()
	if err != nil {
		return false
	}
	return true
}

// AlterMsg 更新评论的内容
func AlterMsg(msg model.Msg) bool {
	query := `UPDATE msg SET content = ? WHERE id = ? AND user_id = ?`
	_, err := db.Exec(query, msg.Content, msg.Id, msg.User_id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// DelMsg 删
func DelMsg(msg model.Msg) bool {
	query := `DELETE FROM msg WHERE id = ? AND user_id = ?`
	_, err := db.Exec(query, msg.Id, msg.User_id)
	if err != nil {
		return false
	}
	return true
}
