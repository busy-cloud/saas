package internal

import "time"

// User 用户
type User struct {
	Id        string    `json:"id" xorm:"pk"`
	TenantId  string    `json:"tenant_id,omitempty" xorm:"index"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Cellphone string    `json:"cellphone,omitempty"`
	Admin     bool      `json:"admin,omitempty"`
	Roles     []string  `json:"roles,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
}

type Role struct {
	Id         string    `json:"id" xorm:"pk"`
	TenantId   string    `json:"tenant_id,omitempty" xorm:"index"`
	Name       string    `json:"name,omitempty"`
	Privileges string    `json:"privileges,omitempty"`
	Disabled   bool      `json:"disabled,omitempty"`
	Created    time.Time `json:"created,omitempty" xorm:"created"`
}

// Password 密码
type Password struct {
	Id       string `json:"id" xorm:"pk"`
	Password string `json:"password"`
}

type UserLog struct {
	Id      string    `json:"id"`
	Name    string    `json:"name,omitempty"`
	Action  string    `json:"action,omitempty"`
	Client  string    `json:"client,omitempty"`
	Ip      string    `json:"ip,omitempty"`
	Created time.Time `json:"created,omitempty" xorm:"created"`
}

type Tenant struct {
	Id       string    `json:"id" xorm:"pk"`
	ParentId string    `json:"parent_id,omitempty" xorm:"index"`
	Name     string    `json:"name,omitempty"`
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created,omitempty" xorm:"created"`
}
