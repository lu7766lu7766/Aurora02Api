package model

type RegisteredLogs struct {
	PingTime   string `gorm:"column:PingTime"`
	CustomerNO string `gorm:"column:CustomerNO"`
}

func (RegisteredLogs) TableName() string {
	return "RegisteredLogs"
}