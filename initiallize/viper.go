package initiallize

import (
	"github.com/spf13/viper"
	"kubeimooc.com/global"
)

// viper 是一种读取go语言的配置文件的工具包， 可以读取json, yaml, toml, hcl, 等格式的配置文件
func Viper() {
	v := viper.New() // 创建viper实例
	// 读取配置文件
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	// 反序列化, 将配置文件中的内容读取到全局变量中
	err = v.Unmarshal(&global.CONF)
	if err != nil {
		// 如果反序列化失败，则直接panic 就是直接停止程序运行
		panic(err.Error())
	}
}
