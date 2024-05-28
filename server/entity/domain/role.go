package domain

type Role struct {
	GlobalDBModel
	RoleId   uint   `gorm:"not null;unique;primary_key;comment:角色ID;"`
	RoleName string `gorm:"not null;unique;comment:角色名"`
}

func (Role) TableName() string {
	return "sys_roles"
}
