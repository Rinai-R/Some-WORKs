package api

import (
	"Golang/12December/20241202/Lottery/dao"
	"Golang/12December/20241202/Lottery/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Host(c *gin.Context) {
	host := c.PostForm("host")
	end_time := c.PostForm("end_time")
	start_time := c.PostForm("start_time")
	event_name := c.PostForm("event_name")

	if host == "" || end_time == "" || start_time == "" || event_name == "" {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "你的申请里好像少了点什么~",
		})
		return
	}
	et, err := time.Parse("2006-01-02 15:04:05", end_time)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}
	st, err1 := time.Parse("2006-01-02 15:04:05", start_time)
	if err1 != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": err1.Error(),
		})
	}
	id := dao.CreateLottery(event_name, et, st, host)

	c.JSON(200, gin.H{
		"code":         200,
		"host":         host,
		"lottery_name": event_name,
		"end_time":     et,
		"start_time":   st,
		"event_id":     id,
	})
	return
}

func AddPrize(c *gin.Context) {
	name := c.PostForm("name")
	lottery_id := c.PostForm("lottery_id")
	num := c.PostForm("num")
	N, err := strconv.Atoi(num)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
	}
	dao.PrizeAdd(name, N, lottery_id)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
	return
}

func QueryPrize(c *gin.Context) {
	lottery_id := c.PostForm("lottery_id")
	user_id := c.PostForm("user_id")
	ET := dao.SearchTime(lottery_id)
	if time.Now().Add(8 * time.Hour).After(ET) {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "你来晚了",
		})
		return
	}
	Key, PrizeName := utils.Draw(lottery_id)
	if Key == "None" {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "宝宝没货了",
		})
		return
	} else if Key == "False" {
		c.JSON(200, gin.H{
			"code":    400,
			"message": "没中奖~",
		})
		return
	}
	mes := fmt.Sprintf("你！中！奖！了！奖品为：%s", PrizeName)
	dao.LotteryQuery(user_id, lottery_id)
	c.JSON(200, gin.H{
		"code":    200,
		"message": mes,
		"ETime":   ET,
		"TNow":    time.Now(),
	})
	return
}

func DelLottery(c *gin.Context) {
	lottery_id := c.PostForm("lottery_id")
	flag := dao.LotteryDel(lottery_id)
	if !flag {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "删除失败了，也许是你已经删了？",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "删除成功啦~",
	})
	return
}
