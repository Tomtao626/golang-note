package base

import "github.com/astaxie/beego"

type Base struct {
	beego.Controller
	Uid int
	UserName string
	IsLogin bool
}

func (b *Base)BaseRender()  {

	b.Uid = 1
	b.UserName = "hallen"
	b.IsLogin = true
	b.Data["Uid"] = 1
	b.Data["UserName"] = "hallen"
	b.Data["IsLogin"] = true

}
