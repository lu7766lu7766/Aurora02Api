package model

type CallPlan struct {
	UserID   string `gorm:"column:UserID"`
	CallOutID string `gorm:"column:CallOutID"`
	PlanName string `gorm:"column:PlanName"`
	StartCalledNumber string `gorm:"column:StartCalledNumber"`
	EndCalledNumber string `gorm:"column:EndCalledNumber"`
	CalledCount int `gorm:"column:CalledCount"`
	CalloutCount int `gorm:"column:CalloutCount"`
	CallConCount int `gorm:"column:CallConCount"`
	CallSwitchCount int `gorm:"column:CallSwitchCount"`
	UseState bool `gorm:"column:UseState"`
	NumberMode int `gorm:"column:NumberMode"`
	CalloutGroupID string `gorm:"column:CalloutGroupID"`
	ConcurrentCalls string `gorm:"column:ConcurrentCalls"`
}

func (CallPlan) TableName() string {
	return "CallPlan"
}