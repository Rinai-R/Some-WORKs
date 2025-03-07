package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {

	// 创建一个新的 Resty 客户端
	client := resty.New()

	// 假设在其他地方已经获取到的 session、courseId 和 rollCallId
	session := "V2-1-65547dc9-835d-4daa-8e3e-6652178824da.MjUwMzkw.1741320621302.doqHJ3Dgvsw7RhrZmsdVJ3QHbR0"
	courseId := "64102"    // 课程 ID
	rollCallId := "331981" // 签到 ID

	// 1. 获取当前正在签到的课程
	headers1 := map[string]string{
		"Accept-Language":  "zh-Hans",
		"Host":             "lms.tc.cqupt.edu.cn",
		"Origin":           "http://mobile.tc.cqupt.edu.cn",
		"Referer":          "http://mobile.tc.cqupt.edu.cn/",
		"X-Forwarded-User": "P338kFwtHL4GEPN3",
		"X-Requested-With": "XMLHttpRequest",
		"X-SESSION-ID":     session,
	}

	resp1, err := client.R().
		SetHeaders(headers1).
		Get("http://lms.tc.cqupt.edu.cn/api/radar/rollcalls?api_version=1.10")

	if err != nil {
		fmt.Println("获取正在签到的课程失败:", err)
		return
	}
	// 处理获取到的课程数据
	fmt.Println("正在签到的课程响应状态:", resp1.Status())
	fmt.Println("课程信息:", string(resp1.Body()))

	// 2. 获取二维码数据
	url2 := fmt.Sprintf("http://lms.tc.cqupt.edu.cn/api/course/%s/rollcall/%s/qr_code", courseId, rollCallId)

	headers2 := map[string]string{
		"Accept-Language":  "zh-CN,zh;q=0.9,en;q=0.8",
		"Host":             "lms.tc.cqupt.edu.cn",
		"Origin":           "http://mobile.tc.cqupt.edu.cn",
		"Referer":          "http://mobile.tc.cqupt.edu.cn/",
		"X-Forwarded-User": "P338kFwtHL4GEPN3",
		"X-SESSION-ID":     session,
	}

	resp2, err := client.R().
		SetHeaders(headers2).
		Get(url2)

	if err != nil {
		fmt.Println("获取二维码数据失败:", err)
		return
	}

	// 处理获取到的二维码数据
	fmt.Println("二维码数据响应状态:", resp2.Status())
	fmt.Println("二维码信息:", string(resp2.Body()))

	// 3. 发送签到请求
	url3 := fmt.Sprintf("http://lms.tc.cqupt.edu.cn/api/rollcall/%s/answer_qr_rollcall", rollCallId)

	headers3 := map[string]string{
		"Accept-Language":  "zh-Hans",
		"Host":             "lms.tc.cqupt.edu.cn",
		"Origin":           "http://mobile.tc.cqupt.edu.cn",
		"Referer":          "http://mobile.tc.cqupt.edu.cn/",
		"X-Forwarded-User": "P338kFwtHL4GEPN3",
		"X-Requested-With": "XMLHttpRequest",
		"Content-Type":     "application/json",
		"X-SESSION-ID":     session,
	}

	// 假设 getData(courseId, rollCallId) 返回签到所需的数据
	bodyData := getData(courseId, rollCallId) // 根据您的逻辑生成请求体

	resp3, err := client.R().
		SetHeaders(headers3).
		SetBody(bodyData).
		Put(url3)

	if err != nil {
		fmt.Println("发送签到请求失败:", err)
		return
	}

	// 处理签到请求的响应
	fmt.Println("签到响应状态:", resp3.Status())
	fmt.Println("签到结果:", string(resp3.Body()))
}

// 示例函数，返回签到请求体的生成逻辑
func getData(courseId string, rollCallId string) interface{} {
	// 根据具体需求构建请求体
	return map[string]interface{}{
		"courseId":   courseId,
		"rollCallId": rollCallId,
		// 其他需要的参数
		//
	}
}
