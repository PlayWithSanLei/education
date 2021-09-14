package setting

import "github.com/spf13/viper"

type Setting struct {
	VP *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")  // 设置文件名称（无后缀）
	vp.AddConfigPath("./confs") // 设置文件所在路径

	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误
			return nil, err
		} else {
			// 配置文件找到后发生了其他错误
			return nil, err
		}
	}
	return &Setting{vp}, nil
}
