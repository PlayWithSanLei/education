package db

import (
	"database/sql"

	"github.com/impact-eintr/education/global"
	"github.com/impact-eintr/eorm"
)

var DB = struct{}{}

func init() {
	setting := eorm.Settings{
		DriverName: "mysql",
		User:       global.DatabaseSetting.User,
		Password:   global.DatabaseSetting.Password,
		Database:   global.DatabaseSetting.DBname,
		Host:       global.DatabaseSetting.Host,
		Options:    map[string]string{"charset": "utf8mb4"},
	}

	var err error
	global.DB, err = sql.Open(setting.DriverName, setting.DataSourceName())
	if err != nil {
		panic(err)
	}

	global.DBClients = make(chan *eorm.Client, 50)
	for i := 0; i < cap(global.DBClients); i++ {
		c, err := eorm.NewClientWithDBconn(global.DB)
		if err != nil {
			panic(err)
		}
		global.DBClients <- c
	}
}
