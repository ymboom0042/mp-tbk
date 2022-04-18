/**
 * @Author: YMBoom
 * @Description:
 * @File:  api
 * @Version: 1.0.0
 * @Date: 2022/01/13 11:18
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ymboom0042/mp-tbk/controller"
)

type apiRouter struct {
	group *gin.RouterGroup
}

func NewApiRouter(g *gin.RouterGroup) *apiRouter {
	return &apiRouter{
		group: g.Group("/api"),
	}
}

func (ar *apiRouter) Route() {
	ar.auth()
	ar.message()
	ar.test()
}

// 微信认证
func (ar *apiRouter) auth() {
	var a controller.AuthController
	ar.group.GET("/wx", a.Auth)
}

func (ar *apiRouter) message() {
	var m controller.MessageController
	ar.group.POST("/wx", m.Receive)
}

func (ar *apiRouter) test() {
	var t controller.TestController
	ar.group.GET("/test", t.Test)
}
