package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type DmToken struct {
	TokenWithTime string `json:"token_with_time,omitempty"`
	Token         string `json:"token,omitempty"`
	EncToken      string `json:"enc_token,omitempty"`
}

type DmRes struct {
	Api  *string         `json:"api,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
	Ret  []string        `json:"ret,omitempty"`
	V    *string         `json:"v,omitempty"`
}

type CommonParams struct {
	Jsv          string `json:"jsv"`
	AppKey       string `json:"appKey"`
	Type         string `json:"type"`
	DataType     string `json:"dataType"`
	H5Request    string `json:"H5Request"`
	AntiCreep    string `json:"AntiCreep"`
	AntiFlood    string `json:"AntiFlood"`
	T            string `json:"t"`
	RequestStart string `json:"requestStart"`
}

func Build() map[string]interface{} {
	t := time.Now()
	millis := t.UnixNano() / int64(time.Millisecond)

	commonParams := &CommonParams{
		Jsv:          "2.7.2",
		AppKey:       "12574478",
		Type:         "originaljson",
		DataType:     "json",
		H5Request:    "true",
		AntiCreep:    "true",
		AntiFlood:    "true",
		T:            fmt.Sprintf("%d", millis),
		RequestStart: fmt.Sprintf("%d", millis-1),
	}
	jsonBytes, _ := json.Marshal(commonParams)
	params := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &params)
	return params
}
