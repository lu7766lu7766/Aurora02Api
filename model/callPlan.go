package model

type CallPlan struct {
	UserID   string `gorm:"column:UserID"`
	CallOutID string `gorm:"column:CallOutID"`
	PlanName string `gorm:"column:PlanName"`
	StartCalledNumber string `gorm:"column:StartCalledNumber"`
	CalledCount string `gorm:"column:CalledCount"`
	CalloutCount string `gorm:"column:CalloutCount"`
	CallConCount string `gorm:"column:CallConCount"`
	CallSwitchCount string `gorm:"column:CallSwitchCount"`
	UseState string `gorm:"column:UseState"`
	NumberMode string `gorm:"column:NumberMode"`
	CalloutGroupID string `gorm:"column:CalloutGroupID"`
	ConcurrentCalls string `gorm:"column:ConcurrentCalls"`
}

func (CallPlan) TableName() string {
	return "CallPlan"
}