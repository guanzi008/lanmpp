// Action测试

package main

import (
	"github.com/go-apibox/api"
)

func TestOkAction(c *api.Context) interface{} {
	return "ok"
}

func TestErrorAction(c *api.Context) interface{} {
	return c.Error.New(api.ErrorInvalidParam, "Password", "TooShort")
}
