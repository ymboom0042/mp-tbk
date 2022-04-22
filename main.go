/**
 * @Author: YMBoom
 * @Description:
 * @File:  mian
 * @Version: 1.0.0
 * @Date: 2022/01/13 11:11
 */
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mp-17208-top/pkg/initialize"
	"mp-17208-top/router"
)

func main() {
	// gin engine
	r := gin.Default()

	// 全局配置
	initialize.Initialize()

	// 路由
	router.NewApiRouter(&r.RouterGroup).Route()

	// 驱动
	r.Run(viper.GetString("app.port"))
}
