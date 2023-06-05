package service

import (
	"auto-ticket-go/client"
	"auto-ticket-go/config"
	"auto-ticket-go/models"
	"encoding/json"
	"fmt"
	"strings"
)

type DmTicket struct {
	cookie  string
	client  *client.DmClient
	account *config.Account
}

func NewDmTicket(account *config.Account) *DmTicket {
	cookie := strings.ReplaceAll(strings.TrimSpace(account.Cookie), " ", "")
	parts := strings.Split(cookie, ";")
	var cleanedParts []string
	for _, part := range parts {
		if !strings.HasPrefix(part, "_m_h5_tk") {
			cleanedParts = append(cleanedParts, part)
		}
	}
	cleanedCookie := strings.Join(cleanedParts, ";")
	client := client.DmClient{}
	return &DmTicket{cookie: cleanedCookie, client: &client, account: account}
}

func (dt *DmTicket) GetTicketInfo(ticketID string) models.TicketInfo {
	url := "https://mtop.damai.cn/h5/mtop.alibaba.damai.detail.getdetail/1.2"
	params := models.BuildTicketInfoParams()
	p := make(map[string]string)
	for k, v := range params {
		p[k] = v.(string)
	}
	fmt.Printf("p: %v \n", p)
	fmt.Println("******************")
	data := models.BuildTicketInfoForm(ticketID)
	fmt.Printf("data: %v \n", data)
	fmt.Println("******************")
	res := dt.client.Request(dt.cookie, url, p, data)
	//fmt.Printf("res: %v \n", res)
	fmt.Println("******************")
	//fmt.Println(string(res.Data))
	var resultMap map[string]interface{}
	json.Unmarshal(res.Data, &resultMap)
	fmt.Println(resultMap)
	return models.TicketInfo{}
}
