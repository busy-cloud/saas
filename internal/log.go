package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
)

func init() {

	api.Register("POST", "saas/log/count", curd.ApiCount[UserLog]())

	api.Register("POST", "saas/log/search", curd.ApiSearch[UserLog]())

	api.Register("GET", "saas/log/list", curd.ApiList[UserLog]())

	api.Register("POST", "saas/log/create", curd.ApiCreate[UserLog]())

	api.Register("GET", "saas/log/:id", curd.ApiGet[UserLog]())

	api.Register("GET", "saas/log/:id/delete", curd.ApiDelete[UserLog]())

}
