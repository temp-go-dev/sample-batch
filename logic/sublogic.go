package logic

import (
	"fmt"

	"github.com/temp-go-dev/sample-batch/db"
	"github.com/temp-go-dev/sample-batch/model"
)

// Sublogicinterface interface
type Sublogicinterface interface {
	SelectData(string) error
	CreatePanic() error
}

// SublogicRegister st
type SublogicRegister struct {
}

// SelectData データを取得する
func (sub *SublogicRegister) SelectData(args string) ([]model.User, error) {
	db := db.GetDB()
	users := []model.User{}

	if args == "9999" {
		// エラーになるSQLを実行
		err := db.Raw("aaaaaaSELECT * FROM user").Scan(&users).Error
		if err != nil {
			fmt.Println("db errror")
			return nil, err
		}
	} else {
		err := db.Raw("SELECT * FROM user").Scan(&users).Error
		if err != nil {
			fmt.Println("db errror")
			return nil, err
		}
	}

	return users, nil
}

// OutputUser user情報をファイルに出力
func (sub *SublogicRegister) OutputUser(users []model.User) error {
	fmt.Println(users)
	return nil
}

// CreatePanic パニック生成
// func (sub *SublogicRegister) CreatePanic() error {

// 	fmt.Println("createPanic")
// 	panic(errorh.NewExitError(
// 		errorh.ExitCodeError,
// 		"E001",
// 		nil,
// 	))
// 	// panic("panic!!!!")
// }
