package initialize

import (
	"MSC2021/src/global"
	"MSC2021/src/initialize/internal"
	"MSC2021/src/models"
	"MSC2021/src/services"
	"github.com/nu7hatch/gouuid"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func InitTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
		models.Question{},
		models.Submission{},
		models.Group{},
		models.Interview{},
	)
	if err != nil {
		global.LOGGER.Error("register table failed", zap.Any("err", err))
		os.Exit(1)
	}
	global.LOGGER.Info("Successfully registered data models.")

	var superUser []models.User
	if err = db.Model(&models.User{}).Where("admin = ?", true).Find(&superUser).Error; err != nil || len(superUser) == 0 {
		global.LOGGER.Info("Admin not found, created it.")

		newSuperUser := models.User{
			Name:      "admin",
			StudentID: "00000000000",
			Admin:     true,
			Major:     "admin",
			Telephone: "15900000000",
			Email:     "admin@xdmsc.club",
			QQ:        "000000000",
			Level:     "2019",
			Intro:     "admin",
		}

		uid, err := uuid.NewV4()
		if uid == nil {
			global.LOGGER.Error("create super user failed", zap.Any("err", err))
			os.Exit(1)
		}
		global.LOGGER.Sugar().Warnf("Super User: %s %s", newSuperUser.Email, uid.String())
		if err != nil {
			global.LOGGER.Error("create super user failed", zap.Any("err", err))
			os.Exit(1)
		}
		newSuperUser.Password = uid.String()
		_, err = services.RegisterWithUser(newSuperUser)
		if err != nil {
			global.LOGGER.Error("create super user failed", zap.Any("err", err))
			os.Exit(1)
		}
	} else {
		global.LOGGER.Info("Database status recovered.")
	}
}

// GormMysql 初始化数据库 *gorm.DB
func GormMysql() *gorm.DB {
	m := global.CONFIG.Database
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.LOGGER.Info("Connected to database: " + m.Path)
		return db
	}
}

func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.CONFIG.Database.LogZap {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = internal.Default.LogMode(logger.Info)
			break
		}
		config.Logger = internal.Default.LogMode(logger.Silent)
	}
	return config
}
