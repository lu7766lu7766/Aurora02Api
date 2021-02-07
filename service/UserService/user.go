package UserService

import (
	db "Aurora02Api/database"
	"Aurora02Api/model"
)

func GetSubEmp(userID string, subLevel int) map[int]model.SysUser {

	var a_res = map[int]model.SysUser{}
	var users []model.SysUser
	query := db.Eloquent.Where("ParentID = ?", userID).Find(&users)
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
		for k, v := range GetSubEmp(user.UserID, subLevel+1) {
			a_res[k] = v
		}
		i++
	}
	return a_res
	// log.Println(users.UserID)
}

func GetList() []model.SysUser {
	var users []model.SysUser
	var err error
	if err = db.Eloquent.Find(&users).Error; err != nil {
		panic(err)
	}
	return users
}
