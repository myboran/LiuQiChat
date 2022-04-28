package global

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	MysqlConn *gorm.DB
)

type SendAllLog struct {
	Id   uint64 `gorm:"primaryKey"`
	Uuid string
	Time time.Time
	Msg  string
}

func InitMysql() {
	var err error
	MysqlConn, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		MysqlInfo.User,
		MysqlInfo.PassWord,
		MysqlInfo.Host,
		MysqlInfo.Port,
		MysqlInfo.DBName,
	))
	if err != nil {
		fmt.Println("connect db err: ", err)
	}

	if MysqlConn.HasTable(&SendAllLog{}) { //判断表是否存在
		//存在就自动适配表，也就说原先没字段的就增加字段
	} else {
		MysqlConn.CreateTable(&SendAllLog{}) //不存在就创建新表
	}
}

func SaveSendAllLog(uuid, msg string) {

	data := SendAllLog{
		Uuid: uuid,
		Time: time.Now(),
		Msg:  msg,
	}
	MysqlConn.Create(&data)
}
