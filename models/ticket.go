package models

import (
	"encoding/json"
	"github.com/fatih/structs"
)

type TicketInfoForm struct {
	ItemID    string `json:"itemId"`
	DMChannel string `json:"dmChannel"`
}

func BuildTicketInfoForm(ticketID string) map[string]interface{} {
	form := TicketInfoForm{
		ItemID:    ticketID,
		DMChannel: "damai@damaih5_h5",
	}
	return structs.Map(form)
}

type TicketInfoParams struct{}

func BuildTicketInfoParams() map[string]interface{} {
	commonParams := Build()
	params := structs.Map(&commonParams)
	params["api"] = "mtop.alibaba.damai.detail.getdetail"
	params["v"] = "1.2"
	return params
}

type SimpleSku struct {
	SkuID   string `json:"skuId"`
	SkuName string `json:"skuName"`
}

type SimplePerform struct {
	PerformID   string `json:"performId"`
	ItemID      string `json:"itemId"`
	PerformName string `json:"performName"`
}

type PerformBase struct {
	Name               string    `json:"name"`
	TimeSpan           string    `json:"timeSpan"`
	PerformBaseTagDesc string    `json:"performBaseTagDesc"`
	Performs           []Perform `json:"performs"`
}

type TicketDetail struct {
	SellStartTimestamp string        `json:"sellStartTime"`
	SellStartTimeStr   string        `json:"sellStartTimeStr"`
	PerformBases       []PerformBase `json:"performBases"`
}

type StaticDataItemBase struct {
	ItemID   string `json:"itemId"`
	ItemName string `json:"itemName"`
}

type StaticData struct {
	ItemBase StaticDataItemBase `json:"itemBase"`
}

type DetailViewComponentItem struct {
	StaticData     StaticData      `json:"staticData"`
	DynamicExtData json.RawMessage `json:"dynamicExtData"`
	Item           TicketDetail    `json:"item"`
}

type DetailViewComponentMap struct {
	Atmosphere json.RawMessage         `json:"atmosphere"`
	Item       DetailViewComponentItem `json:"item"`
}

type TicketInfo struct {
	DetailViewComponentMap DetailViewComponentMap `json:"detailViewComponentMap"`
}
