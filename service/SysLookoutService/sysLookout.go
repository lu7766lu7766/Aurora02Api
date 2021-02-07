package SysLookoutService

import (
	db "Aurora02Api/database"
	"Aurora02Api/model"
	"math"
)

func GetCallStatusContent(UserID string) map[string]interface{} {

	var res map[string]interface{}
	res = make(map[string]interface{})

	var data1 []model.CallState
	db.Eloquent.
		// Select("CalledId, CallDuration, Seat, NormalCall").
		// Where("CallDuration = 0").
		// Where("ExtensionNo = '' or ExtensionNo is null").
		Where("UserID = ?", UserID).Find(&data1)

	res["data1"] = data1

	// var data2 []model.CallState
	// db.
	// 	Select("ExtensionNo, CalledId, CalloutGroupID, CallDuration, PingTime, Seat, NormalCall, OnMonitor").
	// 	Joins("left join RegisteredLogs on ExtensionNo = CustomerNO").
	// 	Where("CallDuration > 0 or ExtensionNo = 'system").
	// 	// Or("ExtensionNo = 'system'").
	// 	// Where("ExtensionNo <> '' or ExtensionNo is not null").
	// 	Where("UserID = ?", UserID).
	// 	Find(&data2)

	// res["data2"] = data2
	//
	var data3 []model.CallPlan
	db.Eloquent.Where("UserID = ?", UserID).Find(&data3)

	res["data3"] = data3
	//
	var waitExtensionNoCount int
	db.Eloquent.
		Table("CallState").
		Where("CallDuration > 1").
		Where("ExtensionNo = 'system'").
		Where("UserID = ?", UserID).
		Count(&waitExtensionNoCount)
	res["waitExtensionNoCount"] = waitExtensionNoCount

	var extensionNoCount int
	db.Eloquent.
		Table("CallState").
		Where("CallDuration > 0").
		Where("ExtensionNo is not null or ExtensionNo <> ''").
		Where("UserID = ?", UserID).
		Count(&extensionNoCount)
	res["extensionNoCount"] = extensionNoCount

	var user model.SysUser
	db.Eloquent.Where("UserID = ?", UserID).Find(&user)
	res["balance"] = math.Round(user.Balance*100) / 100
	res["suspend"] = !user.Suspend

	return res
}
