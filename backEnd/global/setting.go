package global

import (
	"fmt"
	"log"

	"github.com/impact-eintr/education/internal/dao/db"
	"github.com/impact-eintr/education/pkg/logger"
	"github.com/impact-eintr/education/pkg/rbac"
	"github.com/impact-eintr/education/pkg/setting"
	sf "github.com/impact-eintr/education/pkg/snowflake"
	"github.com/impact-eintr/education/pkg/trans"
	"go.uber.org/zap"
)

var (
	Conf            *setting.Setting
	ServerSetting   *setting.ServerSettingS
	LoggerSetting   *setting.LoggerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RBACSetting     *setting.RBACSettingS
)

func Init() {
	var err error

	Conf, err = setting.NewSetting()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	if err = Conf.ReadSection("log", &LoggerSetting); err != nil {
		fmt.Println("init logger failed, err: ", err)
		panic(err)
	}

	_ = logger.L
	zap.L().Debug("logger init success...")

	// 初始化服务配置
	if err = Conf.ReadSection("server", &ServerSetting); err != nil {
		zap.L().Error("init failed, err: ", zap.Error(err))
		panic(err)
	}

	// 初始化ID生成器
	_ = sf.S
	zap.L().Debug("ID init success...")

	// 初始化翻译器设置
	_ = trans.T
	zap.L().Debug("Trans init success...")

	// 初始化sql设置
	if err = Conf.ReadSection("database", &DatabaseSetting); err != nil {
		zap.L().Error("init database failed, err:", zap.Error(err))
		panic(err)
	}

	_ = db.DB
	zap.L().Debug("DB init success...")

	// 初始化RBAC

	err = Conf.ReadSection("rbac", &RBACSetting)
	log.Println(err)
	if err != nil {
		zap.L().Error("init RBAC failed, err:", zap.Error(err))
		panic(err)
	}

	_ = rbac.R
	zap.L().Debug("RBAC init success...")

}
