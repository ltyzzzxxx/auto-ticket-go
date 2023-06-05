package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	SUCCESS_CODE      = 200
	SYSTEM_ERROR_CODE = 500
)

type TokenClient struct{}

func (tc *TokenClient) getValue(key string) string {
	api := "http://localhost:8123"
	query := api + "?key=" + key
	resp, err := http.Get(query)
	if err != nil {
		log.Printf("Fail to get %s.", key)
		return ""
	}
	var data map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Fail to get %s.", key)
		return ""
	}
	if data["code"] == "200" {
		log.Printf("Fail to get %s.", key)
		return ""
	}
	data = data["data"].(map[string]interface{})
	fmt.Printf("token客户端 %v\n", data["value"].(string))
	return data["value"].(string)
}

func (tc *TokenClient) GetBxToken() string {
	return tc.getValue("bx-umidtoken")
}

func (tc *TokenClient) GetBxUa() string {
	return tc.getValue("bx-ua")
}
