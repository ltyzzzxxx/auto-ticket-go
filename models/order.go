package models

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
)

type OrderParams struct{}

func BuildOrderParams() map[string]interface{} {
	commonParams := Build()
	params := structs.Map(&commonParams)
	params["v"] = "4.0"
	params["api"] = "mtop.trade.order.build.h5"
	params["method"] = "POST"
	params["ttid"] = "#t#ip##_h5_2014"
	params["globalCode"] = "ali.china.damai"
	params["tb_eagleeyex_scm_project"] = "20190509-aone2-join-test"
	params["AntiFlood"] = true
	return params
}

type OrderForm struct{}

func BuildOrderForm(itemID *string, skuID *string, byNum int) map[string]interface{} {
	extParams := map[string]interface{}{
		"channel":        "damai_app",
		"damai":          "1",
		"umpChannel":     "100031004",
		"subChannel":     "damai@damaih5_h5",
		"atomSplit":      "1",
		"serviceVersion": "2.0.0",
		"customerType":   "default",
	}

	extParamsJSON, _ := json.Marshal(extParams)

	data := map[string]interface{}{
		"buyNow":    true,
		"exParams":  string(extParamsJSON),
		"buyParam":  fmt.Sprintf("%s_%d_%s", *itemID, byNum, *skuID),
		"dmChannel": "damai@damaih5_h5",
	}

	return data
}

type OrderInfoContainer struct{}

type OrderInfoData struct{}

type OrderInfoEndpoint struct{}

type OrderInfoGlobal struct {
	SecretKey   string `json:"secretKey"`
	SecretValue string `json:"secretValue"`
}

type OrderInfoHierarchy struct {
	Component []string        `json:"component"`
	Root      string          `json:"root"`
	BaseType  []string        `json:"baseType"`
	Structure json.RawMessage `json:"structure"`
}

type OrderInfoLinkageCommon struct {
	QueryParams    string `json:"queryParams"`
	Compress       bool   `json:"compress"`
	ValidateParams string `json:"validateParams"`
	Structures     string `json:"structures"`
	SubmitParams   string `json:"submitParams"`
}

type OrderInfoLinkage struct {
	Input     []string               `json:"input"`
	Request   []string               `json:"request"`
	Signature string                 `json:"signature"`
	Common    OrderInfoLinkageCommon `json:"common"`
}

type OrderInfo struct {
	Data      json.RawMessage    `json:"data"`
	Endpoint  OrderInfoEndpoint  `json:"endpoint"`
	Global    OrderInfoGlobal    `json:"global"`
	Hierarchy OrderInfoHierarchy `json:"hierarchy"`
	Linkage   OrderInfoLinkage   `json:"linkage"`
}

type SubmitOrderParams struct{}

func BuildSubmitOrderParams(submitref string) map[string]interface{} {
	commonParams := Build()
	params := structs.Map(&commonParams)
	params["api"] = "mtop.trade.order.create.h5"
	params["v"] = "4.0"
	params["submitref"] = submitref
	params["timeout"] = "15000"
	params["isSec"] = "1"
	params["ecode"] = "1"
	params["post"] = "1"
	params["ttid"] = "#t#ip##_h5_2014"
	params["globalCode"] = "ali.china.damai"
	params["tb_eagleeyex_scm_project"] = "20190509-aone2-join-test"
	return params
}
