package router

import (
	"blog/dao"
	"blog/utils"
)

// InitGlobalVariable 初始化全局变量
func InitGlobalVariable() {
	// 初始化 Viper
	utils.InitViper()
	// 初始化 Logger
	utils.InitLogger()
	// 初始化 DB
	dao.DB = utils.InitMySQLDB()
}
