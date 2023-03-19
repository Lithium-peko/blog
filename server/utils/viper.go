package utils

import (
	"blog/config"
	"flag"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// InitViper 优先级: 命令行 > 默认值
func InitViper() {
	// 根据命令行读取配置文件路径
	var configPath string
	flag.StringVar(&configPath, "c", "", "chose config file.")
	flag.Parse()
	if configPath != "" { // 命令行读取到参数
		log.Printf("命令行参数为空, 配置文件路径为: %s\n", configPath)
	} else { // 命令行未读取到参数
		log.Println("命令行参数为空, 默认加载 :config/config.toml")
		configPath = "config/config.toml"
	}

	// 目前读取固定路径的配置文件
	v := viper.New()
	v.SetConfigFile(configPath)
	v.AutomaticEnv()                                   // 允许使用环境变量
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // SERVER_APPMODE => SERVER.APPMODE

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Panic("配置文件读取失败: ", err)
	}

	// 加载配置文件内容到结构体对象
	if err := v.Unmarshal(&config.Cfg); err != nil {
		log.Panic("配置文件内容加载失败: ", err)
	}

	log.Println("配置文件内容加载成功")
}
