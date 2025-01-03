package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:~Cy710822@tcp(127.0.0.1:3306)/lottery")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func CreateLottery(event_name string, end_time time.Time, start_time time.Time, host string) int {
	query := "INSERT INTO lottery_events(event_name, end_time, start_time, host) VALUES(?, ?, ?, ?)"
	idGet, err := db.Exec(query, event_name, end_time, start_time, host)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	id, err := idGet.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return int(id)
}

func PrizeAdd(name string, num int, lottery_id string) {
	query := "INSERT INTO prizes(name, num, remain_num, lottery_id) VALUES(?, ?, ?, ?)"
	_, err := db.Exec(query, name, num, num, lottery_id)
	if err != nil {
		log.Fatal(err)
	}
}

func LotteryQuery(user_id string, lottery_id string) {
	query := "INSERT INTO lottery_query(lottery_id, user_id) VALUES(?, ?)"
	_, err := db.Exec(query, lottery_id, user_id)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func SearchTime(lottery_id string) time.Time {
	query := "SELECT end_time FROM lottery_events WHERE id = ?"
	var et []byte
	err := db.QueryRow(query, lottery_id).Scan(&et)

	if err != nil {
		log.Fatal(err)
		return time.Now()
	}
	ET, err2 := time.Parse("2006-01-02 15:04:05", string(et))
	if err2 != nil {
		log.Fatal(err2)
		return time.Now()
	}

	return ET
}

func SearchRemain(lottery_id string, PrizeNum int) (int, map[int]int, map[int]int) {
	var remain int
	remain = 0
	var MAP_num_id map[int]int
	MAP_num_id = make(map[int]int)
	var MAP_id_num map[int]int
	MAP_id_num = make(map[int]int)
	for i := 0; i < PrizeNum; i++ {
		var id int
		query := "SELECT id AS prize_id FROM prizes WHERE lottery_id = ? LIMIT 1 OFFSET ?"
		err := db.QueryRow(query, lottery_id, i).Scan(&id)
		if err != nil {
			log.Fatal(err)
			return 0, MAP_num_id, MAP_id_num
		}
		query2 := "SELECT remain_num FROM prizes WHERE id = ?"
		var remain_num int
		err2 := db.QueryRow(query2, id).Scan(&remain_num)
		if err2 != nil {
			log.Fatal(err2)
			return 0, MAP_num_id, MAP_id_num
		}
		MAP_num_id[i] = id
		MAP_id_num[id] = remain_num
		remain += remain_num
	}
	return remain, MAP_num_id, MAP_id_num
}

func SearchPrizeNum(lottery_id string) int {
	query := "SELECT COUNT(*) AS num_of_related_prizes FROM prizes WHERE  lottery_id = ?"
	var num int
	err := db.QueryRow(query, lottery_id).Scan(&num)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return num
}

func UpdateRemain(prize_id int, remain_num int) {
	query := "UPDATE prizes SET remain_num = ? WHERE id = ?"
	_, err := db.Exec(query, remain_num-1, prize_id)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func SearchPrizeName(id int) string {
	query := "SELECT name FROM prizes WHERE id = ?"
	var name string
	err := db.QueryRow(query, id).Scan(&name)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return name
}

func LotteryDel(lottery_id string) bool {
	query := "DELETE FROM prizes WHERE lottery_id = ?"
	_, err := db.Exec(query, lottery_id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	query1 := "DELETE FROM lottery_query WHERE lottery_id = ?"
	_, err = db.Exec(query1, lottery_id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	query3 := "DELETE FROM lottery_events WHERE id = ?"
	_, err = db.Exec(query3, lottery_id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
