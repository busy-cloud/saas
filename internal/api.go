package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/db"
	"github.com/gin-gonic/gin"
)

func init() {

	api.RegisterUnAuthorized("POST", "saas/user/login", login)
	//api.RegisterUnAuthorized("POST", "login", login)

	api.RegisterUnAuthorized("POST", "saas/user/auth", auth)
	//api.RegisterUnAuthorized("GET", "auth", auth)

	api.Register("GET", "saas/user/logout", logout)
	//api.Register("GET", "logout", logout)

	api.Register("POST", "saas/user/password", password)
	//api.Register("POST", "password", password)

	api.Register("GET", "saas/user/me", userMe)
	//api.Register("GET", "me", userMe)

}

func userMe(ctx *gin.Context) {
	id := ctx.GetString("user")
	var user User
	has, err := db.Engine().ID(id).Get(&user)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if !has {
		api.Fail(ctx, "用户不存在")
		return
	}
	api.OK(ctx, &user)
}
