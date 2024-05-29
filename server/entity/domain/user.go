package domain

type User struct {
	GlobalDBModel
	Username  string `gorm:"index;comment:用户名;type:varchar(10)"`
	Password  string `gorm:"comment:密码;type:varchar(20)" json:"-"`
	Email     string `gorm:"comment:邮箱;type:varchar(200)"`
	Enable    int8   `gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`
	HeaderImg string `gorm:"comment:用户头像;type:varchar(200)"`
	RoleId    uint   `gorm:"comment:用户角色ID"`
	Role      Role   `gorm:"foreignKey:RoleId;references:ID;comment:用户角色"`
}

func (User) TableName() string {
	return "sys_users"
}
