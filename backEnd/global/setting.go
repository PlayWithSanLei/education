package global

import (
	"github.com/impact-eintr/education/pkg/setting"
	"go.uber.org/zap"
)

var (
	Conf            *setting.Setting
	ServerSetting   *setting.ServerSettingS
	LoggerSetting   *setting.LoggerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	RBACSetting     *setting.RBACSettingS
)

func init() {
	var err error
	Conf, err = setting.NewSetting()
	if err != nil {
		panic(err)
	}
	// 初始化服务配置
	if err = Conf.ReadSection("server", &ServerSetting); err != nil {
		zap.L().Error("init failed, err: ", zap.Error(err))
		panic(err)
	}

}
