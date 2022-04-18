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
	"github.com/ymboom0042/mp-tbk/pkg/initialize"
	"github.com/ymboom0042/mp-tbk/router"
)

func main() {
	// gin engine
	r := gin.Default()

	// 全局配置
	initialize.Initialize()

	// 路由
	router.NewApiRouter(&r.RouterGroup).Route()

	// 驱动
	err := r.Run(viper.GetString("app.port"))
	if err != nil {
		return
	}
}
