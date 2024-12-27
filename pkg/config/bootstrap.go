package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// **嵌入文件只能在写embed指令的Go文件的同级目录或者子目录中
//
/*//go:embed *.yaml
var configs embed.FS*/

func InitConfig(path string) {
	//env := os.Getenv("ENV")
	// 根据环境变量 ENV 决定要读取的应用启动配置
	//configFileStream, err := configs.ReadFile("application." + env + ".yaml")
	//if err != nil {
	//	panic(err)
	//}
	vp := viper.New()
	vp.SetConfigFile(path)
	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}

	vp.UnmarshalKey("app", &App)
	vp.UnmarshalKey("database", &Database)
	vp.UnmarshalKey("redis", &Redis)

	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		_ = vp.ReadInConfig()
		vp.UnmarshalKey("app", &App)
		vp.UnmarshalKey("database", &Database)
		vp.UnmarshalKey("redis", &Redis)
	})
}
