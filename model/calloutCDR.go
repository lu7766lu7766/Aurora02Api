package model

type CallOutCDR struct {
	UserID   string `gorm:"column:UserID"`
	CallStartBillingDate string `gorm:"column:CallStartBillingDate"`
	RecordFile string `gorm:"column:RecordFile"`
	ExtensionNo string `gorm:"column:ExtensionNo"`
	OrgCalledID string `gorm:"column:OrgCalledID"`
	CallDuration string `gorm:"column:CallDuration"`
}

func (CallOutCDR) TableName() string {
	return "CallOutCDR"
}