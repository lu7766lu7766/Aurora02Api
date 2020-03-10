package model

type SysUser struct {
	UserID   string `gorm:"column:UserID"`
	UserName string `gorm:"column:UserName"`
	UserGroup string `gorm:"column:UserGroup"`
	MenuList string `gorm:"column:MenuList"`
	ParentID string `gorm:"column:ParentID"`
	PermissionControl string `gorm:"column:PermissionControl"`
	Balance float64 `gorm:"column:Balance"`
	Suspend string `gorm:"column:Suspend"`
	Class int
}

func (SysUser) TableName() string {
	return "SysUser"
}