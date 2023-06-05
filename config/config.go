package config

import (
	"github.com/go-yaml/yaml"
	"os"
)

type Ticket struct {
	ID       string `json:"id"`
	Num      uint   `json:"num"`
	Sessions uint   `json:"sessions"`
	Grade    uint   `json:"grade"`
}

type Account struct {
	Cookie             string `json:"cookie"`
	Remark             string `json:"remark"`
	Ticket             Ticket `json:"ticket"`
	EarliestSubmitTime uint   `json:"earliest_submit_time"`
	Interval           uint   `json:"interval"`
}

type Config struct {
	Accounts []Account `json:"accounts"`
}

func LoadConfig(path string) *Config {
	data, _ := os.ReadFile(path)
	var config Config
	yaml.Unmarshal(data, &config)
	//var m map[string]interface{}
	//yaml.Unmarshal(data, &m)
	//var config Config
	//accounts := m["accounts"].([]interface{})
	//for _, a := range accounts {
	//	account := a.(map[interface{}]interface{})
	//	ticketMap := account["ticket"].(map[interface{}]interface{})
	//	ticket := Ticket{
	//		ID:       ticketMap["id"].(string),
	//		Num:      uint(ticketMap["num"].(int)),
	//		Sessions: uint(ticketMap["sessions"].(int)),
	//		Grade:    uint(ticketMap["grade"].(int)),
	//	}
	//	var interval *uint64
	//	val := uint64(account["interval"].(int))
	//	interval = &val
	//	var earliestSubmitTime *int64
	//	v := int64(account["earliest_submit_time"].(int))
	//	earliestSubmitTime = &v
	//	config.Accounts = append(config.Accounts, Account{
	//		Cookie:             account["cookie"].(string),
	//		Remark:             account["remark"].(string),
	//		Ticket:             ticket,
	//		Interval:           interval,
	//		EarliestSubmitTime: earliestSubmitTime,
	//	})
	//}
	return &config
}

func LoadGlobalConfig() *Config {
	return LoadConfig("config/config.yaml")
}
