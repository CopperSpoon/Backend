package routers

import (
	"BBlog/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/BBlog/member", &controllers.MemberController{}, "post:Create")

}
