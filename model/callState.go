package model

type CallState struct {
	CalledId   string `gorm:"column:CalledId"`
	CallDuration string `gorm:"column:CallDuration"`
	Seat string `gorm:"column:Seat"`
	NormalCall bool `gorm:"column:NormalCall"`
	ExtensionNo string `gorm:"column:ExtensionNo"`
	CalloutGroupID string `gorm:"column:CalloutGroupID"`
	OnMonitor bool `gorm:"column:OnMonitor"`
	PingTime string `gorm:"column:PingTime"`
	// RegisteredLogs []RegisteredLogs `gorm:"foreignkey:CustomerNO;association_foreignkey:ExtensionNo"`
}

func (CallState) TableName() string {
	return "CallState"
}