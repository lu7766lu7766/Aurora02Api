package DownloadService

import (
	db "Aurora02Api/database"
	"Aurora02Api/model"
	"Aurora02Api/service/UserService"
	"Aurora02Api/tools"
	"fmt"
	"os"
	"path"

	"github.com/nleeper/goment"
)

func GetRecordFilePath(userID string, connectDate string, fileName string) string {
	return fmt.Sprintf("D:\\Recording\\%s\\%s\\%s", userID, connectDate, fileName)
}

func GetRecordFile(
	UserID string,
	ConnectDate string,
	FileName string) string {
	return GetRecordFilePath(UserID, ConnectDate, FileName)
	// "/Users/lu7766/go/src/aurora02api/main.go"
}

func GetRecordFilesToZip(
	UserID string,
	CallStartBillingDate string,
	CallStopBillingDate string,
	ExtensionNo string,
	OrgCalledID string,
	DurationCondition string,
	CallDuration string) string {

	emps := []string{UserID}
	sub_emp := UserService.GetSubEmp(UserID, 1)
	for _, v := range sub_emp {
		emps = append(emps, v.UserID)
	}

	var callOutCDRs []model.CallOutCDR
	db.Eloquent = db.Eloquent.Where("cast(CallStartBillingDate as datetime) between ? and ?", CallStartBillingDate, CallStopBillingDate).
		Where("RecordFile <> ''").
		Where("UserID in (?)", emps)
	// .Debug()
	if ExtensionNo != "" {
		db.Eloquent = db.Eloquent.Where("ExtensionNo = ?", ExtensionNo)
	}
	if OrgCalledID != "" {
		db.Eloquent = db.Eloquent.Where("OrgCalledID = ?", OrgCalledID)
	}
	if CallDuration != "" {
		if DurationCondition == "within" {
			db.Eloquent = db.Eloquent.Where("CallDuration <= ?", CallDuration)
		} else {
			db.Eloquent = db.Eloquent.Where("CallDuration > ?", CallDuration)
		}
	}

	query := db.Eloquent.Find(&callOutCDRs)

	if query.Error != nil {
		panic(query.Error)
	}

	// fmt.Println(callOutCDRs)

	if len(callOutCDRs) == 0 {
		return ""
	}

	var files []*os.File
	for _, data := range callOutCDRs {
		moment, _ := goment.New(data.CallStartBillingDate, "YYYY/MM/DD")
		path := GetRecordFilePath(data.UserID, moment.Format("YYYYMMDD"), data.RecordFile)
		if _, err := os.Stat(path); err == nil {
			// path/to/whatever exists
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			files = append(files, file)
		}
	}

	fileName := UserID + "RecordFile.zip"
	filePath := path.Join(os.TempDir(), fileName)

	err := tools.Compress(files, filePath)
	if err != nil {
		panic(err)
	}

	return filePath
}
