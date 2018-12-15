// 路由配置

package main

import (
	"github.com/go-apibox/api"
)

var apiRoutes = []*api.Route{
	api.NewRoute("Test.Ok", TestOkAction),
	api.NewRoute("Test.Error", TestErrorAction),
}
