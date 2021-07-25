package gen

var routerTemplate = `package {SystemName}

import (
	"server/app/system/{SystemName}/api"

	"github.com/gogf/gf/net/ghttp"
	"github.com/iWinston/qk-library/frame/qcmd"
)

var Cmds = []qcmd.Domain{}

// README: 本系统不采用接口创建权限和角色
func InitNormalRouter(group *ghttp.RouterGroup) {
	group.Group("/{SystemName}", func(group *ghttp.RouterGroup) {
		group.Group("/{Name}", func(group *ghttp.RouterGroup) {
			group.REST("/", api.{CamelName})
			group.GET("/list", api.{CamelName}.List)
		})
	})
}
`
