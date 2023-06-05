package models

import (
	"encoding/json"
	"github.com/fatih/structs"
)

type PerformParams struct{}

func BuildPerformParams() map[string]interface{} {
	commonParams := Build()
	params := structs.Map(&commonParams)
	params["api"] = "mtop.alibaba.detail.subpage.getdetail"
	params["method"] = "GET"
	params["v"] = "2.0"
	return params
}

type PerformForm struct{}

func BuildPerformForm(ticketId string, performId string) map[string]interface{} {
	exParams := map[string]interface{}{
		"dataType":       2,
		"dataId":         performId,
		"privilegeActId": "",
	}
	exParamsString, _ := json.Marshal(exParams)

	data := map[string]interface{}{
		"itemId":    ticketId,
		"bizCode":   "ali.china.damai",
		"scenario":  "itemsku",
		"exParams":  string(exParamsString),
		"dmChannel": "damai@damaih5_h5",
	}
	return data
}

type Sku struct {
	SkuID     string `json:"skuId"`
	ItemID    string `json:"itemId"`
	PriceName string `json:"priceName"`
	Price     string `json:"price"`
}

type Perform struct {
	PerformID   string `json:"performId"`
	PerformName string `json:"performName"`
	SkuList     []Sku  `json:"skuList"`
}

type PerformInfo struct {
	Perform Perform `json:"perform"`
}
