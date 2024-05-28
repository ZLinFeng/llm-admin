package domain

type User struct {
	GlobalDBModel
	Id        uint   `gorm:"not null;unique;primary_key;comment:用户ID;"`
	Username  string `gorm:"index;comment:用户名"`
	Password  string `gorm:"comment:密码" json:"-"`
	Email     string `gorm:"comment:邮箱"`
	Enable    int8   `gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
	HeaderImg string `gorm:"comment:用户头像"`
	RoleId    uint   `gorm:"comment:用户角色ID"`
	Role      Role   `gorm:"foreignKey:RoleId;references:RoleId;comment:用户角色"`
}

func (User) TableName() string {
	return "sys_users"
}
