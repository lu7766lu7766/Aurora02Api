package service

import (
	DB "aurora02api/database"
	"aurora02api/model"
	"fmt"
	"github.com/nleeper/goment"
	"os"
	"aurora02api/tools"
	"path"
)

type DownloadService struct {}

func (this DownloadService) GetRecordFilePath(userID string, connectDate string , fileName string) string {
	return fmt.Sprintf("D:\\Recording\\%s\\%s\\%s", userID, connectDate, fileName)
}

func (this DownloadService) GetRecordFile(
		UserID string,
		ConnectDate string,
		FileName string) string {
	return this.GetRecordFilePath(UserID, ConnectDate, FileName)
	// "/Users/lu7766/go/src/aurora02api/main.go"
}

func (this DownloadService) GetRecordFilesToZip(
		UserID string,
		CallStartBillingDate string,
		CallStopBillingDate string,
		ExtensionNo string,
		OrgCalledID string,
		DurationCondition string,
		CallDuration string) string {

	db := DB.Connect()
	defer db.Close()

	emps := []string { UserID }
	sub_emp := UserService{}.GetSubEmp(UserID, 1)
	for _, v := range sub_emp {
		emps = append(emps, v.UserID)
	}

	var callOutCDRs []model.CallOutCDR
	db = db.Where("cast(CallStartBillingDate as datetime) between ? and ?", CallStartBillingDate, CallStopBillingDate).
		Where("RecordFile <> ''").
		Where("UserID in (?)", emps)
	// .Debug()
	if ExtensionNo != "" {
		db = db.Where("ExtensionNo = ?", ExtensionNo)
	}
	if OrgCalledID != "" {
		db = db.Where("OrgCalledID = ?", OrgCalledID)
	}
	if CallDuration != "" {
		if DurationCondition == "within" {
			db = db.Where("CallDuration <= ?", CallDuration)
		} else {
			db = db.Where("CallDuration > ?", CallDuration)
		}
	}

	query := db.Find(&callOutCDRs)

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
		path := this.GetRecordFilePath(data.UserID, moment.Format("YYYYMMDD"), data.RecordFile)
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

	fileName :=  UserID + "RecordFile.zip"
	filePath := path.Join(os.TempDir(), fileName)

	err := tools.Compress(files, filePath)
	if err != nil {
		panic(err)
	}
	
	return filePath
}