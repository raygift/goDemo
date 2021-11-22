package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var tr *http.Transport

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,
		// 下面的代码被干掉了
		//Dial: func(netw, addr string) (net.Conn, error) {
		//    conn, err := net.DialTimeout(netw, addr, time.Second*2) //设置建立连接超时
		//    if err != nil {
		//        return nil, err
		//    }
		//    err = conn.SetDeadline(time.Now().Add(time.Second * 3)) //设置发送接受数据超时
		//    if err != nil {
		//        return nil, err
		//    }
		//    return conn, nil
		//},
	}
}
func main() {

	fmt.Println("---------------------------------------------")

	/* 构建请求正文 */
	reqBody := strings.NewReader(`
			{
				"name": "helloworld"
			}
		`)

	/* 创建请求对象 */
	req, err := http.NewRequest("POST", "http://192.168.92.3:8080/", reqBody)
	if err != nil {
		panic(err)
	}

	/* 设置请求头域 */
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Transport: tr,
		Timeout:   3 * time.Second, // 超时加在这里，是每次调用的超时
	}
	ack, err := client.Do(req)
	/* 发请求收应答 */
	// ack, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	/* 读取应答正文 */
	ackBody, err := ioutil.ReadAll(ack.Body)
	/* 关闭应答正文，释放资源，无论是否异常 */
	ack.Body.Close()
	if err != nil {
		panic(err)
	}

	/* 输出应答状态 */
	fmt.Printf("HTTP Response StatusCode: %d\n", ack.StatusCode)
	fmt.Printf("HTTP Response Status: %s\n", ack.Status)

	/* 输出应答头域 */
	fmt.Printf("HTTP Response HEADER: %s\n", ack.Header.Get("my-http-head"))

	/* 输出应答正文 */
	fmt.Printf("HTTP Response BODY: %s\n", ackBody)
	time.Sleep(2 * time.Second)
}
