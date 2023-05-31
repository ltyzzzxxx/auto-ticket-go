package models

import (
	"encoding/json"
	"github.com/fatih/structs"
)

type PerformParams struct {}

func (pp *PerformParams) build() (map[string]interface{}, error) {
	commonParams := Build()
	params := structs.Map(&commonParams)
	params["api"] = "mtop.alibaba.detail.subpage.getdetail"
	params["method"] = "GET"
	params["v"] = "2.0"
	return params, nil
}

type PerformForm struct {}

func (pf *PerformForm) build(ticketId string, performId string) (map[string]interface{}, error)  {
	exParams := map[string]interface{}{
		"dataType":      2,
		"dataId":        performId,
		"privilegeActId": "",
	}
	exParamsString, err := json.Marshal(exParams)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"itemId":     ticketId,
		"bizCode":    "ali.china.damai",
		"scenario":   "itemsku",
		"exParams":   string(exParamsString),
		"dmChannel":  "damai@damaih5_h5",
	}
	return data, nil
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

