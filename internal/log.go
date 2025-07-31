package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"time"
)

func init() {
	db.Register(new(UserLog))
}

type UserLog struct {
	Id      string    `json:"id"`
	Name    string    `json:"name,omitempty"`
	Action  string    `json:"action,omitempty"`
	Client  string    `json:"client,omitempty"`
	Ip      string    `json:"ip,omitempty"`
	Created time.Time `json:"created,omitempty" xorm:"created"`
}

func init() {

	api.Register("POST", "saas/log/count", curd.ApiCount[UserLog]())

	api.Register("POST", "saas/log/search", curd.ApiSearch[UserLog]())

	api.Register("GET", "saas/log/list", curd.ApiList[UserLog]())

	api.Register("POST", "saas/log/create", curd.ApiCreate[UserLog]())

	api.Register("GET", "saas/log/:id", curd.ApiGet[UserLog]())

	api.Register("GET", "saas/log/:id/delete", curd.ApiDelete[UserLog]())

}
