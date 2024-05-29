package domain

type Role struct {
	GlobalDBModel
	RoleName string `gorm:"type:varchar(20);not null;unique;comment:角色名"`
}

func (Role) TableName() string {
	return "sys_roles"
}
