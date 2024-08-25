package main

import (
	"fmt"
	"github.com/hewo233/cxsj_homework1/common"
	"github.com/hewo233/cxsj_homework1/model"
)

func main() {
	db, _ := common.InitDB()
	defer common.CloseDB(db)

	for {
		var choose string
		var info model.CxsjUser

		_, _ = fmt.Scanln(&choose)
		if choose == "add" {
			fmt.Println("Please input email, name, password, gender:")
			_, _ = fmt.Scanln(&info.Email, &info.Name, &info.Password, &info.Gender)
			_ = common.AddRecord(&info, db)
		} else if choose == "delete" {
			fmt.Println("Please input email:")
			_, _ = fmt.Scanln(&info.Email)
			_ = common.DeleteRecord(info.Email, db)
		} else if choose == "update" {
			fmt.Println("Please input email, name, password, gender:")
			_, _ = fmt.Scanln(&info.Email, &info.Name, &info.Password, &info.Gender)
			_ = common.ModifyRecord(&info, db)
		} else if choose == "query" {
			fmt.Println("Please input email:")
			_, _ = fmt.Scanln(&info.Email)
			userInfo, _ := common.QueryRecordByEmail(info.Email, db)
			fmt.Println(userInfo)
		} else if choose == "exit" {
			break
		} else {
			fmt.Println("Invalid command")
		}
	}
}
