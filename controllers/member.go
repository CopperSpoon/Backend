package controllers

import (
	"BBlog/models/model"
	"BBlog/models/repository"
	"BBlog/models/service"
	"encoding/json"
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

// MemberController operations for Member
type MemberController struct {
	beego.Controller
}

//Create -
func (m *MemberController) Create() {
	req := new(model.MemberReq)

	//解析並放進struct裡面
	if err := json.Unmarshal(m.Ctx.Input.RequestBody, &req); err != nil {
		fmt.Println("json:", err)

	}
	member := new(model.Member)

	member, err := service.CheckMember(req)

	memberRes := new(model.Response)
	memberRes.Message = repository.Create(member)

	//有錯誤的話就回傳沒有就傳錯誤訊息給前端
	if memberRes.Message == "" {
		memberRes.Code = "000"
		memberRes.Message = "Success"
	} else {
		memberRes.Code = "404"
		memberRes.Message = err.Error()
	}
	//將struct的內容以json的方式傳出去
	m.Data["json"] = &memberRes
	m.ServeJSON()
}
