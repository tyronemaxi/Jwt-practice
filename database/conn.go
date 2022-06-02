package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"website/util/flag"
)

const (
	NotDelete = 0
	IsDelete  = 1

	MysqlDriver    = "mysql"
	PostgresDriver = "postgres"
)

var (
	gormDB *gorm.DB
)

func Init() (err error) {
	if gormDB != nil {
		return
	}

	address := ""
	var dialector gorm.Dialector
	//if flag.DBDriver == MysqlDriver {
	address = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", flag.DBUser,
		flag.DBPassword, flag.DBHost, flag.DBPort, flag.DBDatabase)

	if flag.MysqlDriverParams != "" {
		address = fmt.Sprintf("%s?%s", address, flag.MysqlDriverParams)
	}

	dialector = mysql.New(mysql.Config{
		DSN:                       address, // DSN data source name
		DefaultStringSize:         256,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度， MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	})
	//} else if flag.DBDriver == PostgresDriver {
	//	address = fmt.Sprintf("user=%s, password=%s, dbname=%s host=%s port=%d sslmode=disable client_encoding=UTF8 TimeZone=Asia/Shanghai", flag.DBUser,
	//		flag.DBPassword, flag.DBDatabase, flag.DBHost, flag.DBPort)
	//
	//	dialector = postgres.New(postgres.Config{
	//		DSN: address,
	//		PreferSimpleProtocol: true,
	//	})
	//
	//}
	gormDB, err = gorm.Open(dialector)
	if err != nil {
		return err
	}
	return nil

}

type BaseDao struct {
	Engine *gorm.DB
}

func DB() *gorm.DB {
	if flag.DBDebugMode {
		return gormDB.Debug()
	}
	return gormDB
}

func (p *BaseDao) GetDB() *gorm.DB {
	return DB()
}
