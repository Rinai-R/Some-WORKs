package dao

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func Publish(message string, username string) string {
	userid := GetUserId(username)
	query := "INSERT INTO messages(user_id, host_id, context) VALUES (?, ?, ?)"
	idGet, err := db.Exec(query, userid, userid, message)
	if err != nil {
		log.Println(err)
		return ""
	}
	id, err2 := idGet.LastInsertId()
	if err2 != nil {
		log.Fatal(err)
		return ""
	}
	ID := strconv.FormatInt(id, 10)
	return ID
}

func Reply(parent_id string, username string, message string) string {
	user_id := GetUserId(username)
	query := "SELECT host_id, is_closed FROM messages WHERE id = ?"
	var host_id string
	var closed int
	err := db.QueryRow(query, parent_id).Scan(&host_id, &closed)
	if err != nil {
		log.Println(err)
		return ""
	}
	if closed == 1 {
		return "closed"
	}
	query2 := "INSERT INTO messages(user_id, host_id, context, parent_id) VALUES (?, ?, ?, ?)"
	idGet, err2 := db.Exec(query2, user_id, host_id, message, parent_id)
	if err2 != nil {
		log.Fatal(err2)
		return ""
	}
	id, err3 := idGet.LastInsertId()
	if err3 != nil {
		log.Fatal(err3)
	}
	ID := strconv.FormatInt(id, 10)
	return ID
}

func CloseMsg(id string, username string) bool {
	user_id := GetUserId(username)
	query := "SELECT host_id FROM messages WHERE id = ?"
	var host_id string
	err := db.QueryRow(query, id).Scan(&host_id)
	if err != nil {
		log.Println(err)
		return false
	}
	query2 := "SELECT user_id FROM messages WHERE id = ?"
	var True_UserId string
	err2 := db.QueryRow(query2, id).Scan(&True_UserId)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	if True_UserId != user_id && host_id != user_id {
		return false
	}
	query3 := "UPDATE messages SET is_closed = ? WHERE id = ?"
	_, err3 := db.Exec(query3, 1, id)
	if err3 != nil {
		log.Println(err3)
		return false
	}
	return true
}

func InOrder(id string, messages *[]string, depth int) bool {
	query := "SELECT context, updated_at, user_id, parent_id, id FROM messages WHERE parent_id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return true
		}
		log.Println(err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var Ctext string
		var updated_at string
		var userId string
		var parent_id string
		var Id string
		err2 := rows.Scan(&Ctext, &updated_at, &userId, &parent_id, &Id)
		if err2 != nil {
			log.Println(err2)
			return false
		}
		var tmp string
		for i := 0; i < depth; i++ {
			tmp += fmt.Sprintf("      ")
		}
		tmp += fmt.Sprintf("消息 %s：用户 %s 于 %s 对 %s 消息给予了回复: %s   ", Id, userId, updated_at, parent_id, Ctext)
		*messages = append(*messages, tmp)
		InOrder(Id, messages, depth+1)
	}
	return true
}

func GetAllMsg(username string) []string {
	user_id := GetUserId(username)
	var messages []string
	query := "SELECT id FROM messages WHERE user_id = ? AND parent_id is NULL"
	rows, err := db.Query(query, user_id)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return nil
	}
	for rows.Next() {
		var id string
		err2 := rows.Scan(&id)
		if err2 != nil {
			log.Println(err2)
			return nil
		}
		tmp := fmt.Sprintf("\n您的消息 %s 的留言：\n    ", id)
		messages = append(messages, tmp)
		depth := 1
		InOrder(id, &messages, depth)
	}
	return messages
}

func DelMsg(id string, username string) bool {
	user_id := GetUserId(username)
	query := "SELECT host_id FROM messages WHERE id = ?"
	var host_id string
	err := db.QueryRow(query, id).Scan(&host_id)
	if err != nil {
		log.Println(err)
		return false
	}
	query2 := "SELECT user_id FROM messages WHERE id = ?"
	var True_UserId string
	err2 := db.QueryRow(query2, id).Scan(&True_UserId)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	if True_UserId != user_id && host_id != user_id {
		return false
	}
	query3 := "DELETE FROM messages WHERE id = ?"
	_, err3 := db.Exec(query3, id)
	if err3 != nil {
		log.Println(err3)
		return false
	}
	return true
}

func OpenMsg(id string, username string) bool {
	user_id := GetUserId(username)
	query := "SELECT host_id FROM messages WHERE id = ?"
	var host_id string
	err := db.QueryRow(query, id).Scan(&host_id)
	if err != nil {
		log.Println(err)
		return false
	}
	query2 := "SELECT user_id FROM messages WHERE id = ?"
	var True_UserId string
	err2 := db.QueryRow(query2, id).Scan(&True_UserId)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	if True_UserId != user_id && host_id != user_id {
		return false
	}
	query3 := "UPDATE messages SET is_closed = ? WHERE id = ?"
	_, err3 := db.Exec(query3, 0, id)
	if err3 != nil {
		log.Println(err3)
		return false
	}
	return true
}

func MsgExistsByIdAndUser(id string, username string) bool {
	user_id := GetUserId(username)
	var Exist bool
	query := "SELECT 1 FROM messages WHERE id = ? AND user_id = ?"
	db.QueryRow(query, id, user_id).Scan(&Exist)
	if Exist {
		return true
	}
	return false
}
func MsgExistsById(id string) bool {
	var Exist bool
	query := "SELECT 1 FROM messages WHERE id = ? AND user_id = ?"
	db.QueryRow(query, id).Scan(&Exist)
	if Exist {
		return true
	}
	return false
}

func ChangeMsg(id string, message string, username string) bool {
	user_id := GetUserId(username)
	query := "UPDATE messages SET context = ? WHERE id = ? AND user_id = ?"
	_, err := db.Exec(query, message, id, user_id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
