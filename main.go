package main

import (
	"auto-ticket-go/config"
	"auto-ticket-go/service"
	"fmt"
)

func main() {
	//tc := client.TokenClient{}
	//value := tc.GetBxToken()
	//fmt.Println(value)

	//client.GetToken("_bl_uid=LOlm6ih5h8edwO7y29Xa0p6n1UFd; _m_h5_tk=f8974e26b93f5e3bf647f859527267b9_1685887906783; _m_h5_tk_enc=147bedfc7de278382702ee0c8a89a6dc; cna=gupzGlh4QlUCAXEaZnOgOyjx; _samesite_flag_=true; cookie2=126645d18589b84ee2c64dd03f4de540; t=f3a98a7ba93ea17e7885f1c16020e533; _tb_token_=e8e765f8a775f; xlly_s=1; ab.storage.deviceId.a9882122-ac6c-486a-bc3b-fab39ef624c5=%7B%22g%22%3A%223b956199-bd4b-80b4-2e2b-ed511446070e%22%2C%22c%22%3A1685879643125%2C%22l%22%3A1685879643125%7D; _hvn_login=18; munb=2215564693912; csg=177d427c; l=fBMuGAV4NcIuDtsdBO5aourza779uIRfhsPzaNbMiIEGa6CNtFwHcNC_JbxHSdtjgTCEWHxygJ-c1d3J7N4NixDDBeAHjt4KnxvO0MP9K; tfstk=cPAfBV4W_hdP4fGgiigyOV5ov7CNZpVfzqsklDOFtyzUXd8fisNFAWPeSuCNWa1..; usercode=224984435; dm_nickname=%E9%BA%A6%E5%AD%904bACC; havanaId=2215564693912; isg=BIKCeJNwUiETrE79h0Bi6rbX047kU4Zth8uq98ybs_WgHyOZtOBrfbFdyxtjT_4F")

	conf := config.LoadGlobalConfig()

	dt := service.NewDmTicket(&conf.Accounts[0])
	fmt.Printf("dt: %v \n", dt)
	fmt.Println("******************")
	dt.GetTicketInfo("710947802955")
}
