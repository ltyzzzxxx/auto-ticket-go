package client

import (
	"auto-ticket-go/models"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type DmClient struct {
	client      *http.Client
	tokenClient TokenClient
	dmToken     models.DmToken
	bxToken     string
}

// 构造函数
//func NewDmClient(cookie string) *DmClient {
//	tokenClient := TokenClient{}
//	bxToken := tokenClient.GetBxToken()
//	dmToken := getToken(cookie)
//
//	//headers := http.Header{}
//	//baseURL := "https://mtop.damai.cn/"
//	//headers.Set("Origin", baseURL)
//	//headers.Set("Referer", baseURL)
//	//cookieValue := strings.Join([]string{
//	//	cookie,
//	//	"_m_h5_tk_enc=" + dmToken.EncToken,
//	//	"_m_h5_tk=" + dmToken.TokenWithTime,
//	//}, ";")
//	//headers.Set("Cookie", cookieValue)
//
//	//client := http.DefaultClient
//	return &DmClient{
//		//client: client,
//		dmToken: dmToken,
//		tokenClient: tokenClient,
//		bxToken: bxToken,
//	}
//}

// 封装统一请求
func (dm *DmClient) Request(cookie string, api string, params map[string]string, data map[string]interface{}) models.DmRes {

	fmt.Printf("初始params %v\n", params)

	// 设置header
	tokenClient := TokenClient{}
	bxToken := tokenClient.GetBxToken()
	dmToken := getToken(cookie)
	headers := http.Header{}
	baseURL := "https://mtop.damai.cn/"
	headers.Set("Origin", baseURL)
	headers.Set("Referer", baseURL)
	headers.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3")
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	cookieValue := strings.Join([]string{
		cookie,
		"_m_h5_tk_enc=" + dmToken.EncToken,
		"_m_h5_tk=" + dmToken.TokenWithTime,
	}, ";")
	fmt.Printf("cookieValue %v\n", cookieValue)
	fmt.Println()
	headers.Set("Cookie", cookieValue)
	// 构造签名sign
	dataJson, _ := json.Marshal(data)
	s := fmt.Sprintf("%s&%s&%s&%s",
		dmToken.Token,
		params["t"],
		params["appKey"],
		string(dataJson))
	sign := fmt.Sprintf("%x", md5.Sum([]byte(s)))
	// 构造请求体
	form := url.Values{}
	form.Add("data", string(dataJson))
	body := bytes.NewBufferString(form.Encode())
	fmt.Printf("dataJson %v\n", string(dataJson))
	fmt.Println()
	// 构造请求参数
	p := url.Values{}
	for k, v := range params {
		p.Add(k, v)
	}
	p.Add("sign", sign)
	p.Add("bx-umidtoken", bxToken)
	p.Add("bx-ua", tokenClient.GetBxUa())
	fmt.Printf("sign %v\n", sign)
	fmt.Println()
	fmt.Printf("bx-umidtoken %v\n", p["bx-umidtoken"])
	fmt.Println()
	fmt.Printf("bx-ua %v\n", p["bx-ua"])
	fmt.Println()
	// 构造请求
	req, _ := http.NewRequest("POST", api+"?"+p.Encode(), body)
	req.Header = headers
	fmt.Printf("request url %v\n", req.URL)
	fmt.Printf("request header %v\n", req.Header)
	fmt.Printf("request body %v\n", req.Body)
	fmt.Println()
	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(resp)
	// 响应处理，返回结果
	if err != nil {
		log.Printf("request failed")
		return models.DmRes{}
	}
	defer resp.Body.Close()
	respData := models.DmRes{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return models.DmRes{}
	}
	return respData
}

// 获取用户token
func getToken(cookie string) models.DmToken {
	headers := http.Header{}
	headers.Add("origin", "https://mtop.damai.cn/")
	headers.Add("referer", "https://mtop.damai.cn/")
	headers.Add("cookie", cookie)
	params := url.Values{}
	for k, v := range models.BuildTicketInfoParams() {
		params.Set(k, v.(string))
	}
	url := "https://mtop.damai.cn/h5/mtop.damai.wireless.search.broadcast.list/1.0/?"
	req, _ := http.NewRequest("GET", url+params.Encode(), nil)
	req.Header = headers
	client := &http.Client{}
	resp, _ := client.Do(req)
	token := models.DmToken{}
	for _, c := range resp.Cookies() {
		if c.Name == "_m_h5_tk" {
			token.TokenWithTime = c.Value
			token.Token = strings.Split(token.TokenWithTime, "_")[0]
		}
		if c.Name == "_m_h5_tk_enc" {
			token.EncToken = c.Value
		}
	}
	return token
}
