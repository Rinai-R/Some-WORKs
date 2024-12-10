package utils

import (
	"Golang/12December/20241202/Lottery/dao"
	"math/rand"
)

func Draw(lottery_id string) (string, string) {
	PrizeNum := dao.SearchPrizeNum(lottery_id)
	remain, MAP_num_id, MAP_id_rem := dao.SearchRemain(lottery_id, PrizeNum)

	if remain <= 0 {
		return "None", "None"
	}
	random_num := rand.Intn(PrizeNum * 2)
	if MAP_id_rem[MAP_num_id[random_num]] <= 0 {
		return "False", "None"
	}
	PrizeName := dao.SearchPrizeName(MAP_num_id[random_num])
	dao.UpdateRemain(MAP_num_id[random_num], MAP_id_rem[MAP_num_id[random_num]])
	return "True", PrizeName
}
