package service

import (
	DB "aurora02api/database"
	"aurora02api/model"
	// "fmt"
)

type SysLookoutService struct {}

func (this SysLookoutService) GetCallStatusContent(UserID string) map[string]interface{} {

	// fmt.Println(UserID)

	db := DB.Connect()
	defer db.Close()

	var res map[string]interface{}
	res = make(map[string]interface{})


	var data1 []model.CallState
	db.
		Select("CalledId", "CallDuration", "Seat", "NormalCall").
		Where("CallDuration = 0").
		Where("ExtensionNo = '' or ExtensionNo is null").
		Where("UserID = ?", UserID).Find(&data1)

	res["data1"] = data1

	// fmt.Println(res, data1)
	//
	//
	var data2 []model.CallState
	db.
		Joins("left join RegisteredLogs on ExtensionNo = CustomerNO").
		Where("CallDuration > 0").
		Where("ExtensionNo <> '' or ExtensionNo is not null").
		Where("UserID = ?", UserID).Find(&data2)
		// .Preload("RegisteredLogs")
		// .Joins("left join RegisteredLogs on ExtensionNo = CustomerNO")
		// .Rows()

	res["data2"] = data2
	//
	var data3 []model.CallPlan
	db.Where("UserID = ?", UserID).Find(&data3)

	res["data3"] = data3
	//
	var waitExtensionNoCount int
	db.
		Table("CallState").
		Where("CallDuration > 1").
		Where("ExtensionNo = 'system'").
		Where("UserID = ?", UserID).
		Count(&waitExtensionNoCount)
	res["waitExtensionNoCount"] = waitExtensionNoCount

	var extensionNoCount int
	db.	
		Table("CallState").
		Where("CallDuration > 0").
		Where("ExtensionNo is not null or ExtensionNo <> ''").
		Where("UserID = ?", UserID).
		Count(&extensionNoCount)
	res["extensionNoCount"] = extensionNoCount

	var user model.SysUser
	db.Where("UserID = ?", UserID).Find(&user)
	res["balance"] = user.Balance
	res["suspend"] = (user.Suspend != "1")

	// fmt.Println(res)

	return res
}