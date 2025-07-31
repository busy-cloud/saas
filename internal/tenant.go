package internal

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"time"
)

func init() {
	db.Register(new(Tenant))
}

// Tenant 租户
type Tenant struct {
	Id       string    `json:"id" xorm:"pk"`
	Name     string    `json:"name,omitempty"`
	ParentId string    `json:"parent_id,omitempty" xorm:"index"`
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created,omitempty" xorm:"created"`
}

func init() {
	api.Register("POST", "saas/tenant/count", curd.ApiCount[Tenant]())
	api.Register("POST", "saas/tenant/search", curd.ApiSearch[Tenant]())
	api.Register("GET", "saas/tenant/list", curd.ApiList[Tenant]())
	api.Register("POST", "saas/tenant/create", curd.ApiCreateHook[Tenant](curd.GenerateID[Tenant](), nil))
	api.Register("GET", "saas/tenant/:id", curd.ApiGet[Tenant]())
	api.Register("POST", "saas/tenant/:id", curd.ApiUpdate[Tenant]("id", "name", "parent_id", "disabled"))
	api.Register("GET", "saas/tenant/:id/delete", curd.ApiDeleteHook[Tenant](nil, nil))
	api.Register("GET", "saas/tenant/:id/enable", curd.ApiDisableHook[Tenant](false, nil, nil))
	api.Register("GET", "saas/tenant/:id/disable", curd.ApiDisableHook[Tenant](true, nil, nil))
}
