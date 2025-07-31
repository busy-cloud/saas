package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"github.com/gin-gonic/gin"
	"time"
)

func init() {
	db.Register(new(User), new(Password))
}

// User 用户
type User struct {
	Id        string    `json:"id" xorm:"pk"`
	TenantId  string    `json:"tenant_id" xorm:"index"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Cellphone string    `json:"cellphone,omitempty"`
	Roles     []string  `json:"roles,omitempty"`
	Admin     bool      `json:"admin,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
}

// Password 密码
type Password struct {
	Id       string `json:"id" xorm:"pk"`
	Password string `json:"password"`
}

func init() {

	api.Register("POST", "saas/user/count", curd.ApiCount[User]())

	api.Register("POST", "saas/user/search", curd.ApiSearch[User]())

	api.Register("GET", "saas/user/list", curd.ApiList[User]())

	api.Register("POST", "saas/user/create", curd.ApiCreateHook[User](curd.GenerateID[User](), nil))

	api.Register("GET", "saas/user/:id", curd.ApiGet[User]())

	api.Register("POST", "saas/user/:id", curd.ApiUpdate[User]("id", "tenant_id", "name", "email", "cellphone", "admin", "disabled"))

	api.Register("GET", "saas/user/:id/delete", curd.ApiDeleteHook[User](nil, nil))

	api.Register("POST", "saas/user/:id/password", userPassword)

	api.Register("GET", "saas/user/:id/enable", curd.ApiDisableHook[User](false, nil, nil))

	api.Register("GET", "saas/user/:id/disable", curd.ApiDisableHook[User](true, nil, nil))
}

type userPasswordObj struct {
	Password string `json:"password"`
}

func userPassword(ctx *gin.Context) {

	var obj userPasswordObj
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		api.Error(ctx, err)
		return
	}

	var p Password
	p.Id = ctx.Param("id")
	p.Password = md5hash(obj.Password)

	_, _ = db.Engine().ID(p.Id).Delete(new(Password)) //不管有没有都删掉
	_, err := db.Engine().InsertOne(&p)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	api.OK(ctx, nil)
}
