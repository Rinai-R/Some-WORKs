package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var chan1 = make(chan Buyer)
var ItemNum = 25
var mutex sync.Mutex

type Buyer struct {
	ID     int
	BuyNum int
}

func init() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机数生成器
}

func Consumer(u int) {
	RandomNum := rand.Intn(10) + 1
	PurChase := Buyer{u, RandomNum}
	chan1 <- PurChase
}

func ShopItem() {
	for ItemNum != 0 {
		mutex.Lock()
		var x Buyer

		x = <-chan1
		if x.BuyNum <= ItemNum {
			ItemNum = ItemNum - x.BuyNum
			fmt.Printf("用户%d成功购买了%d个商品，库存还有%d个商品\n", x.ID, x.BuyNum, ItemNum)
		} else {
			fmt.Printf("库存不足！目前只有%d个商品\n", ItemNum)
		}
		mutex.Unlock()
	}
	fmt.Println("今日商品已售罄，请下次再来~")
}

func main() {
	go ShopItem()

	go Consumer(1)
	go Consumer(2)
	go Consumer(3)
	go Consumer(4)
	go Consumer(5)
	go Consumer(6)
	go Consumer(7)
	go Consumer(8)
	go Consumer(9)

	time.Sleep(time.Second * 3)

}
