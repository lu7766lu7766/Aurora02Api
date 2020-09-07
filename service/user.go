package service

import (
	DB "Aurora02Api/database"
	"Aurora02Api/model"
)

type UserService struct{}

func (this UserService) GetSubEmp(userID string, subLevel int) map[int]model.SysUser {
	db := DB.Connect()
	defer db.Close()

	var a_res = map[int]model.SysUser{}
	var users []model.SysUser
	query := db.Where("ParentID = ?", userID).Find(&users)
	if query.Error != nil {
		panic(query.Error)
	}
	// .Debug()
	// fmt.Println(users)

	var (
		i   = 0
		len = len(users)
	)
	for i < len {
		var user = users[i]
		user.Class = subLevel
		a_res[i] = user
		for k, v := range this.GetSubEmp(user.UserID, subLevel+1) {
			a_res[k] = v
		}
		i++
	}
	return a_res
	// log.Println(users.UserID)
}

func (this UserService) GetList() []model.SysUser {
	db := DB.Connect()
	defer db.Close()
	var users []model.SysUser
	query := db.Find(&users)
	if query.Error != nil {
		panic(query.Error)
	}
	return users
}
