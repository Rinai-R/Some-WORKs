package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	for i := 0; i <= 100; i++ {
		go func() {
			for {
				rsp, err := client.R().SetHeaders(map[string]string{}).Post("http://121.40.101.83/LanShanLibrary/login?username=TWind&password=123456")
				if err != nil {
					fmt.Println(err)
					return
				}
				if rsp.StatusCode() != 200 {
					continue
				}
				fmt.Println(rsp)
			}
		}()
	}
	select {}
}
