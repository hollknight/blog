package model

import (
	"blog/global"
	"blog/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Model struct {
	ID         uint32    `gorm:"primary_key" json:"id"`
	CreatedBy  string    `json:"created_by"`
	ModifiedBy string    `json:"modified_by"`
	CreatedOn  time.Time `gorm:"default:'1000-01-01 00:00:00'" json:"created_on"`
	ModifiedOn time.Time `gorm:"default:'1000-01-01 00:00:00'" json:"modified_on"`
	DeletedOn  time.Time `gorm:"default:'1000-01-01 00:00:00'" json:"deleted_on"`
	IsDel      uint8     `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&collation=%s",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.Port,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
		databaseSetting.Collation,
	)

	config := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // 字符串字段的默认大小
		DisableDatetimePrecision:  true,  // 禁用日期时间精度，MySQL 5.6 之前不支持
		DontSupportRenameIndex:    true,  // 重命名索引时删除和创建，MySQL 5.7 之前不支持重命名索引，MariaDB
		DontSupportRenameColumn:   true,  // `change` 重命名列，MySQL 8 之前不支持重命名列，MariaDB
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}

	db, err := gorm.Open(mysql.New(config), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gormv2_",
			SingularTable: true, // 表前缀和表复数禁用
		},
	})
	if err != nil {
		return nil, err
	}

	// 将 Logger 设置为 debug 模式
	if global.ServerSetting.RunMode == "debug" {
		db.Debug() // db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// 注册回调函数
	db.Callback().Create().Before("gorm:create").Register("update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Before("gorm:update").Register("update_time_stamp", updateTimeStampForUpdateCallback)

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

// 创建前更新 CreatedOn 字段的回调函数
func updateTimeStampForCreateCallback(db *gorm.DB) {
	db.Statement.SetColumn("created_on", time.Now())
}

//func registerCallback(db *gorm.DB) {
//	// 自动添加uuid
//	err := db.Callback().Create().Before("gorm:create").Register("uuid", func (db *gorm.DB) {
//		db.Statement.SetColumn("id", NewUlid())
//	})
//	if err != nil {
//		log.Panicf("err: %+v", errx.WithStackOnce(err))
//	}
//}

// 更新前更新 ModifiedOn 字段的回调函数
func updateTimeStampForUpdateCallback(db *gorm.DB) {
	db.Statement.SetColumn("modified_on", time.Now())
}
