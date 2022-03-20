package data

import (
	"github.com/google/wire"
	"github.com/mars-projects/mars/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderGormSet is data providers.
var ProviderGormSet = wire.NewSet(NewGormClient)

func NewGormClient(conf *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                conf.Database.Driver,
		DSN:                       conf.Database.Source,
		SkipInitializeWithVersion: false, // 根s据当前 MySQL 版本自动配置
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	if err != nil {
		panic(err)
	}
	return db
}

//func migrate(db *gorm.DB) error {
//	err := db.Debug().AutoMigrate(&models.Migration{})
//	if err != nil {
//		return err
//	}
//	migration.Migrate.SetDb(db.Debug())
//	migration.Migrate.Migrate()
//	return nil
//}
