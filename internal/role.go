package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"time"
)

func init() {
	db.Register(new(Role))
}

// Role 角色
type Role struct {
	Id         string    `json:"id" xorm:"pk"`
	TenantId   string    `json:"tenant_id" xorm:"index"`
	Name       string    `json:"name,omitempty"`
	Privileges []string  `json:"privileges,omitempty" xorm:"LONGTEXT"`
	Disabled   bool      `json:"disabled,omitempty"`
	Created    time.Time `json:"created,omitempty" xorm:"created"`
}

func init() {
	api.Register("POST", "saas/role/count", curd.ApiCount[Role]())
	api.Register("POST", "saas/role/search", curd.ApiSearch[Role]())
	api.Register("GET", "saas/role/list", curd.ApiList[Role]())
	api.Register("POST", "saas/role/create", curd.ApiCreateHook[Role](curd.GenerateID[Role](), nil))
	api.Register("GET", "saas/role/:id", curd.ApiGet[Role]())
	api.Register("POST", "saas/role/:id", curd.ApiUpdate[Role]("id", "tenant_id", "name", "privileges", "disabled"))
	api.Register("GET", "saas/role/:id/delete", curd.ApiDeleteHook[Role](nil, nil))
	api.Register("GET", "saas/role/:id/enable", curd.ApiDisableHook[Role](false, nil, nil))
	api.Register("GET", "saas/role/:id/disable", curd.ApiDisableHook[Role](true, nil, nil))
}
